package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mr-tron/base58/base58"
	"github.com/pkg/errors"
	"github.com/schollz/find3/server/main/src/logging"
	"github.com/schollz/find3/server/main/src/models"
	"github.com/schollz/stringsizer"
	flock "github.com/theckman/go-flock"
)

// MakeTables creates two tables, a `keystore` table:
//
// 	KEY (TEXT)	VALUE (TEXT)
//
// and also a `sensors` table for the sensor data:
//
// 	TIMESTAMP (INTEGER)	DEVICE(TEXT) LOCATION(TEXT)
//
// the sensor table will dynamically create more columns as new types
// of sensor data are inserted. The LOCATION column is optional and
// only used for learning/classification.
func (d *Database) MakeTables() (err error) {
	sqlStmt := `create table keystore (key text not null primary key, value text);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		d.logger.Log.Error(err)
		return
	}
	sqlStmt = `create index keystore_idx on keystore(key);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		d.logger.Log.Error(err)
		return
	}
	sqlStmt = `create table sensors (timestamp integer not null primary key, deviceid text, locationid text, unique(timestamp));`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		d.logger.Log.Error(err)
		return
	}
	sqlStmt = `CREATE TABLE location_predictions (timestamp integer NOT NULL PRIMARY KEY, prediction TEXT, UNIQUE(timestamp));`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		d.logger.Log.Error(err)
		return
	}
	sqlStmt = `CREATE TABLE devices (id TEXT PRIMARY KEY, name TEXT);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		d.logger.Log.Error(err)
		return
	}
	sqlStmt = `CREATE TABLE locations (id TEXT PRIMARY KEY, name TEXT);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		d.logger.Log.Error(err)
		return
	}

	sqlStmt = `CREATE TABLE gps (mac TEXT PRIMARY KEY, lat REAL, lon REAL, alt REAL, timestamp INTEGER);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		d.logger.Log.Error(err)
		return
	}

	sensorDataSS, _ := stringsizer.New()
	err = d.Set("sensorDataStringSizer", sensorDataSS.Save())
	if err != nil {
		return
	}
	return
}

// Columns will list the columns
func (d *Database) Columns() (columns []string, err error) {
	rows, err := d.db.Query("SELECT * FROM sensors LIMIT 1")
	if err != nil {
		err = errors.Wrap(err, "Columns")
		return
	}
	columns, err = rows.Columns()
	rows.Close()
	if err != nil {
		err = errors.Wrap(err, "Columns")
		return
	}
	return
}

// Get will retrieve the value associated with a key.
func (d *Database) Get(key string, v interface{}) (err error) {
	stmt, err := d.db.Prepare("select value from keystore where key = ?")
	if err != nil {
		return errors.Wrap(err, "problem preparing SQL")
	}
	defer stmt.Close()
	var result string
	err = stmt.QueryRow(key).Scan(&result)
	if err != nil {
		return errors.Wrap(err, "problem getting key")
	}

	err = json.Unmarshal([]byte(result), &v)
	if err != nil {
		return
	}
	// d.logger.Log.Debugf("got %s from '%s'", string(result), key)
	return
}

// Set will set a value in the database, when using it like a keystore.
func (d *Database) Set(key string, value interface{}) (err error) {
	var b []byte
	b, err = json.Marshal(value)
	if err != nil {
		return err
	}
	tx, err := d.db.Begin()
	if err != nil {
		return errors.Wrap(err, "Set")
	}
	stmt, err := tx.Prepare("insert or replace into keystore(key,value) values (?, ?)")
	if err != nil {
		return errors.Wrap(err, "Set")
	}
	defer stmt.Close()

	_, err = stmt.Exec(key, string(b))
	if err != nil {
		return errors.Wrap(err, "Set")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "Set")
	}

	// d.logger.Log.Debugf("set '%s' to '%s'", key, string(b))
	return
}

// Set will set a value in the database, when using it like a keystore.
func (d *Database) Dump() (err error) {
	command := fmt.Sprintf("sqlite3 d.name .dump")
	d.logger.Log.Debug(command)
	out, err := exec.Command(command).Output()
	fmt.Println(out)
	return
}

// AddPrediction will insert or update a prediction in the database
func (d *Database) AddPrediction(timestamp int64, aidata []models.LocationPrediction) (err error) {
	// truncate to two digits
	for i := range aidata {
		aidata[i].Probability = float64(int64(float64(aidata[i].Probability)*100)) / 100
	}

	var b []byte
	b, err = json.Marshal(aidata)
	if err != nil {
		return err
	}
	tx, err := d.db.Begin()
	if err != nil {
		return errors.Wrap(err, "AddPrediction")
	}
	stmt, err := tx.Prepare("insert or replace into location_predictions (timestamp,prediction) values (?, ?)")
	if err != nil {
		return errors.Wrap(err, "AddPrediction")
	}
	defer stmt.Close()

	_, err = stmt.Exec(timestamp, string(b))
	if err != nil {
		return errors.Wrap(err, "AddPrediction")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "AddPrediction")
	}
	return
}

// GetPrediction will retrieve models.LocationAnalysis associated with that timestamp
func (d *Database) GetPrediction(timestamp int64) (aidata []models.LocationPrediction, err error) {
	stmt, err := d.db.Prepare("SELECT prediction FROM location_predictions WHERE timestamp = ?")
	if err != nil {
		err = errors.Wrap(err, "problem preparing SQL")
		return
	}
	defer stmt.Close()
	var result string
	err = stmt.QueryRow(timestamp).Scan(&result)
	if err != nil {
		err = errors.Wrap(err, "problem getting key")
		return
	}

	err = json.Unmarshal([]byte(result), &aidata)
	if err != nil {
		return
	}
	// d.logger.Log.Debugf("got %s from '%s'", string(result), key)
	return
}

// AddSensor will insert a sensor data into the database
// TODO: AddSensor should be special case of AddSensors
func (d *Database) AddSensor(s models.SensorData) (err error) {
	// determine the current table coluss
	oldColumns := make(map[string]struct{})
	columnList, err := d.Columns()
	if err != nil {
		return
	}
	for _, column := range columnList {
		oldColumns[column] = struct{}{}
	}

	// get string sizer
	var sensorDataStringSizerString string
	err = d.Get("sensorDataStringSizer", &sensorDataStringSizerString)
	if err != nil {
		return
	}
	sensorDataSS, err := stringsizer.New(sensorDataStringSizerString)
	if err != nil {
		return
	}

	// setup the database
	tx, err := d.db.Begin()
	if err != nil {
		return errors.Wrap(err, "AddSensor")
	}

	// first add new columns in the sensor data
	deviceID, err := d.AddName("devices", s.Device)
	if err != nil {
		return errors.Wrap(err, "problem getting device ID")
	}
	locationID := ""
	if len(s.Location) > 0 {
		locationID, err = d.AddName("locations", s.Location)
		if err != nil {
			return errors.Wrap(err, "problem getting location ID")
		}
	}
	args := make([]interface{}, 3)
	args[0] = s.Timestamp
	args[1] = deviceID
	args[2] = locationID
	argsQ := []string{"?", "?", "?"}
	for sensor := range s.Sensors {
		if _, ok := oldColumns[sensor]; !ok {
			stmt, err := tx.Prepare("alter table sensors add column " + sensor + " text")
			if err != nil {
				return errors.Wrap(err, "AddSensor, adding column")
			}
			_, err = stmt.Exec()
			if err != nil {
				return errors.Wrap(err, "AddSensor, adding column")
			}
			d.logger.Log.Debugf("adding column %s", sensor)
			columnList = append(columnList, sensor)
			stmt.Close()
		}
	}

	// organize arguments in the correct order
	for _, sensor := range columnList {
		if _, ok := s.Sensors[sensor]; !ok {
			continue
		}
		argsQ = append(argsQ, "?")
		args = append(args, sensorDataSS.ShrinkMapToString(s.Sensors[sensor]))
	}

	// only use the columns that are in the payload
	newColumnList := make([]string, len(columnList))
	j := 0
	for i, c := range columnList {
		if i >= 3 {
			if _, ok := s.Sensors[c]; !ok {
				continue
			}
		}
		newColumnList[j] = c
		j++
	}
	newColumnList = newColumnList[:j]

	sqlStatement := "insert or replace into sensors(" + strings.Join(newColumnList, ",") + ") values (" + strings.Join(argsQ, ",") + ")"
	stmt, err := tx.Prepare(sqlStatement)
	d.logger.Log.Debug("columns", columnList)
	d.logger.Log.Debug("args", args)
	if err != nil {
		return errors.Wrap(err, "AddSensor, prepare "+sqlStatement)
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return errors.Wrap(err, "AddSensor, execute")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "AddSensor")
	}

	// update the map key slimmer
	err = d.Set("sensorDataStringSizer", sensorDataSS.Save())
	if err != nil {
		return
	}

	d.logger.Log.Debug("inserted sensor data")
	return

}

// GetSensorFromTime will return a sensor data for a given timestamp
func (d *Database) GetSensorFromTime(timestamp interface{}) (s models.SensorData, err error) {
	sensors, err := d.GetAllFromPreparedQuery("SELECT * FROM sensors WHERE timestamp = ?", timestamp)
	if err != nil {
		err = errors.Wrap(err, "GetSensorFromTime")
	} else {
		s = sensors[0]
	}
	return
}

// Get will retrieve the value associated with a key.
func (d *Database) GetLastSensorTimestamp() (timestamp int64, err error) {
	stmt, err := d.db.Prepare("SELECT timestamp FROM sensors ORDER BY timestamp DESC LIMIT 1")
	if err != nil {
		err = errors.Wrap(err, "problem preparing SQL")
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&timestamp)
	if err != nil {
		err = errors.Wrap(err, "problem getting key")
	}
	return
}

// GetSensorFromGreaterTime will return a sensor data for a given timeframe
func (d *Database) GetSensorFromGreaterTime(timeBlockInMilliseconds int64) (sensors []models.SensorData, err error) {
	latestTime, err := d.GetLastSensorTimestamp()
	if err != nil {
		return
	}
	minimumTimestamp := latestTime - timeBlockInMilliseconds
	sensors, err = d.GetAllFromPreparedQuery("SELECT * FROM (SELECT * FROM sensors WHERE timestamp > ? GROUP BY deviceid ORDER BY timestamp DESC)", minimumTimestamp)
	return
}

// GetAnalysisFromGreaterTime will return the analysis for a given timeframe
// func (d *Database) GetAnalysisFromGreaterTime(timestamp interface{}) {
// 	select sensors.timestamp, devices.name, location_predictions.prediction from sensors inner join location_predictions on location_predictions.timestamp=sensors.timestamp inner join devices on sensors.deviceid=devices.id WHERE sensors.timestamp > 0 GROUP BY devices.name ORDER BY sensors.timestamp DESC;
// }

// GetAllForClassification will return a sensor data for classifying
func (d *Database) GetAllForClassification() (s []models.SensorData, err error) {
	return d.GetAllFromQuery("SELECT * FROM sensors WHERE sensors.locationid !=''")
}

// GetLatest will return a sensor data for classifying
func (d *Database) GetLatest(device string) (s models.SensorData, err error) {
	deviceID, err := d.GetID("devices", device)
	if err != nil {
		return
	}
	var sensors []models.SensorData
	sensors, err = d.GetAllFromPreparedQuery("SELECT * FROM sensors WHERE deviceID=? ORDER BY timestamp DESC LIMIT 1", deviceID)
	if err != nil {
		return
	}
	if len(sensors) > 0 {
		s = sensors[0]
	} else {
		err = errors.New("no rows found")
	}
	return
}

func (d *Database) GetKeys(keylike string) (keys []string, err error) {
	query := "SELECT key FROM keystore WHERE key LIKE ?"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		err = errors.Wrap(err, query)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(keylike)
	if err != nil {
		err = errors.Wrap(err, query)
		return
	}
	defer rows.Close()

	keys = []string{}
	for rows.Next() {
		var key string
		err = rows.Scan(&key)
		if err != nil {
			err = errors.Wrap(err, "scanning")
			return
		}
		keys = append(keys, key)
	}
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "rows")
	}
	return
}

func (d *Database) GetDevices() (devices []string, err error) {
	query := "SELECT devicename FROM (SELECT devices.name as devicename,COUNT(devices.name) as counts FROM sensors INNER JOIN devices ON sensors.deviceid = devices.id GROUP by devices.name) ORDER BY counts DESC"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		err = errors.Wrap(err, query)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		err = errors.Wrap(err, query)
		return
	}
	defer rows.Close()

	devices = []string{}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			err = errors.Wrap(err, "scanning")
			return
		}
		devices = append(devices, name)
	}
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "rows")
	}
	return
}

func (d *Database) GetIDToName(table string) (idToName map[string]string, err error) {
	idToName = make(map[string]string)
	query := "SELECT id,name FROM " + table
	stmt, err := d.db.Prepare(query)
	if err != nil {
		err = errors.Wrap(err, query)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		err = errors.Wrap(err, query)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var name, id string
		err = rows.Scan(&id, &name)
		if err != nil {
			err = errors.Wrap(err, "scanning")
			return
		}
		idToName[id] = name
	}
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "rows")
	}
	return
}

func GetFamilies() (families []string) {
	files, err := ioutil.ReadDir(DataFolder)
	if err != nil {
		log.Fatal(err)
	}

	families = make([]string, len(files))
	i := 0
	for _, f := range files {
		if !strings.Contains(f.Name(), ".sqlite3.db") {
			continue
		}
		b, err := base58.Decode(strings.TrimSuffix(f.Name(), ".sqlite3.db"))
		if err != nil {
			continue
		}
		families[i] = string(b)
		i++
	}
	if i > 0 {
		families = families[:i]
	} else {
		families = []string{}
	}
	return
}

// GetID will get the ID of an element in a table (devices/locations) and return an error if it doesn't exist
func (d *Database) GetID(table string, name string) (id string, err error) {
	// first check to see if it has already been added
	stmt, err := d.db.Prepare("SELECT id FROM " + table + " WHERE name = ?")
	defer stmt.Close()
	if err != nil {
		err = errors.Wrap(err, "problem preparing SQL")
		return
	}
	err = stmt.QueryRow(name).Scan(&id)
	return
}

// AddName will add a name to a table (devices/locations) and return the ID. If the device already exists it will just return it.
func (d *Database) AddName(table string, name string) (deviceID string, err error) {
	// first check to see if it has already been added
	deviceID, err = d.GetID(table, name)
	if err == nil {
		return
	}
	// logger.Log.Debugf("creating new name for %s in %s", name, table)

	// get the current count
	stmt, err := d.db.Prepare("SELECT COUNT(id) FROM " + table)
	if err != nil {
		err = errors.Wrap(err, "problem preparing SQL")
		stmt.Close()
		return
	}
	var currentCount int
	err = stmt.QueryRow().Scan(&currentCount)
	stmt.Close()
	if err != nil {
		err = errors.Wrap(err, "problem getting device count")
		return
	}

	// transform the device name into an ID with the current count
	currentCount++
	deviceID = stringsizer.Transform(currentCount)
	// logger.Log.Debugf("transformed (%d) %s -> %s", currentCount, name, deviceID)

	// add the device name and ID
	tx, err := d.db.Begin()
	if err != nil {
		err = errors.Wrap(err, "AddName")
		return
	}
	query := "insert into " + table + "(id,name) values (?, ?)"
	// logger.Log.Debugf("running query: '%s'", query)
	stmt, err = tx.Prepare(query)
	if err != nil {
		err = errors.Wrap(err, "AddName")
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(deviceID, name)
	if err != nil {
		err = errors.Wrap(err, "AddName")
	}
	err = tx.Commit()
	if err != nil {
		err = errors.Wrap(err, "AddName")
		return
	}
	return
}

func Exists(name string) (err error) {
	name = strings.TrimSpace(name)
	name = path.Join(DataFolder, base58.FastBase58Encoding([]byte(name))+".sqlite3.db")
	if _, err = os.Stat(name); err != nil {
		err = errors.New("database '" + name + "' does not exist")
	}
	return
}

// Open will open the database for transactions by first aquiring a filelock.
func Open(family string, readOnly ...bool) (d *Database, err error) {
	d = new(Database)
	d.family = strings.TrimSpace(family)

	// convert the name to base64 for file writing
	// override the name
	if len(readOnly) > 1 && readOnly[1] {
		d.name = path.Join(DataFolder, d.family)
	} else {
		d.name = path.Join(DataFolder, base58.FastBase58Encoding([]byte(d.family))+".sqlite3.db")
	}
	d.logger, err = logging.New()
	if err != nil {
		return
	}
	d.Debug(DebugMode)

	// if read-only, make sure the database exists
	if _, err = os.Stat(d.name); err != nil && len(readOnly) > 0 && readOnly[0] {
		err = errors.New(fmt.Sprintf("group '%s' does not exist", d.family))
		return
	}

	// obtain a lock on the database
	// d.logger.Log.Debugf("getting filelock on %s", d.name+".lock")
	d.fileLock = flock.NewFlock(d.name + ".lock")
	for {
		locked, err := d.fileLock.TryLock()
		if err == nil && locked {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	// check if it is a new database
	newDatabase := false
	if _, err := os.Stat(d.name); os.IsNotExist(err) {
		newDatabase = true
	}

	// open sqlite3 database
	d.db, err = sql.Open("sqlite3", d.name)
	if err != nil {
		return
	}
	// d.logger.Log.Debug("opened sqlite3 database")

	// create new database tables if needed
	if newDatabase {
		err = d.MakeTables()
		if err != nil {
			return
		}
		d.logger.Log.Debug("made tables")
	}

	return
}

func (d *Database) Debug(debugMode bool) {
	if debugMode {
		d.logger.SetLevel("debug")
	} else {
		d.logger.SetLevel("info")
	}
}

// Close will close the database connection and remove the filelock.
func (d *Database) Close() (err error) {
	// close filelock
	err = d.fileLock.Unlock()
	if err != nil {
		d.logger.Log.Error(err)
	} else {
		os.Remove(d.name + ".lock")
		// d.logger.Log.Debug("removed filelock")
	}

	// close database
	err2 := d.db.Close()
	if err2 != nil {
		err = err2
		d.logger.Log.Error(err)
	} else {
		// d.logger.Log.Debug("closed database")
	}
	return
}

func (d *Database) GetAllFromQuery(query string) (s []models.SensorData, err error) {
	// d.logger.Log.Debug(query)
	rows, err := d.db.Query(query)
	if err != nil {
		err = errors.Wrap(err, "GetAllFromQuery")
		return
	}
	defer rows.Close()

	// parse rows
	s, err = d.getRows(rows)
	if err != nil {
		err = errors.Wrap(err, query)
	}
	return
}

// GetAllFromPreparedQuery
func (d *Database) GetAllFromPreparedQuery(query string, args ...interface{}) (s []models.SensorData, err error) {
	// prepare statement
	stmt, err := d.db.Prepare(query)
	if err != nil {
		err = errors.Wrap(err, query)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		err = errors.Wrap(err, query)
		return
	}
	defer rows.Close()
	s, err = d.getRows(rows)
	if err != nil {
		err = errors.Wrap(err, query)
	}
	return
}

func (d *Database) getRows(rows *sql.Rows) (s []models.SensorData, err error) {
	// first get the columns
	columnList, err := d.Columns()
	if err != nil {
		return
	}

	// get the string sizer for the sensor data
	var sensorDataStringSizerString string
	err = d.Get("sensorDataStringSizer", &sensorDataStringSizerString)
	if err != nil {
		return
	}
	sensorDataSS, err := stringsizer.New(sensorDataStringSizerString)
	if err != nil {
		return
	}

	deviceIDToName, err := d.GetIDToName("devices")
	if err != nil {
		return
	}

	locationIDToName, err := d.GetIDToName("locations")
	if err != nil {
		return
	}

	s = []models.SensorData{}
	// loop through rows
	for rows.Next() {
		var arr []interface{}
		for i := 0; i < len(columnList); i++ {
			arr = append(arr, new(interface{}))
		}
		err = rows.Scan(arr...)
		if err != nil {
			err = errors.Wrap(err, "getRows")
			return
		}
		s0 := models.SensorData{
			// the underlying value of the interface pointer and cast it to a pointer interface to cast to a byte to cast to a string
			Timestamp: int64((*arr[0].(*interface{})).(int64)),
			Family:    d.family,
			Device:    deviceIDToName[string((*arr[1].(*interface{})).([]uint8))],
			Location:  locationIDToName[string((*arr[2].(*interface{})).([]uint8))],
			Sensors:   make(map[string]map[string]interface{}),
		}
		// add in the sensor data
		for i, colName := range columnList {
			if i < 3 {
				continue
			}
			if *arr[i].(*interface{}) == nil {
				continue
			}
			shortenedJSON := string((*arr[i].(*interface{})).([]uint8))
			s0.Sensors[colName], err = sensorDataSS.ExpandMapFromString(shortenedJSON)
			if err != nil {
				return
			}
		}
		s = append(s, s0)
	}
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "getRows")
	}
	return
}

// SetGPS will set a GPS value in the GPS database
func (d *Database) SetGPS(gps models.GPS) (err error) {
	tx, err := d.db.Begin()
	if err != nil {
		return errors.Wrap(err, "SetGPS")
	}
	stmt, err := tx.Prepare("insert or replace into gps(mac,lat,lon,alt,timestamp) values (?, ?, ?, ?, ?)")
	if err != nil {
		return errors.Wrap(err, "SetGPS")
	}
	defer stmt.Close()

	_, err = stmt.Exec(gps.Mac, gps.Latitude, gps.Longitude, gps.Altitude, gps.Timestamp)
	if err != nil {
		return errors.Wrap(err, "SetGPS")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "SetGPS")
	}

	return
}

// GetGPS will return a GPS for a given mac, if it exists
// if it doesn't exist it will return an error
func (d *Database) GetGPS(mac string) (gps models.GPS, err error) {
	query := "SELECT mac,lat,lon,alt,timestamp FROM gps WHERE mac == ?"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		err = errors.Wrap(err, query)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(mac)
	if err != nil {
		err = errors.Wrap(err, query)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&gps.Mac, &gps.Latitude, &gps.Longitude, &gps.Altitude, &gps.Timestamp)
		if err != nil {
			err = errors.Wrap(err, "scanning")
			return
		}
	}
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "rows")
	}
	if gps.Mac == "" {
		err = errors.New(mac + " does not exist in gps table")
	}
	return
}

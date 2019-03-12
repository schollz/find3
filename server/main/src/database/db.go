package database

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mr-tron/base58/base58"
	"github.com/pkg/errors"
	"github.com/schollz/find3/server/main/src/models"
	"github.com/schollz/sqlite3dump"
	"github.com/schollz/stringsizer"
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
		logger.Log.Error(err)
		return
	}
	sqlStmt = `create index keystore_idx on keystore(key);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		logger.Log.Error(err)
		return
	}
	sqlStmt = `create table sensors (timestamp integer not null primary key, deviceid text, locationid text, unique(timestamp));`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		logger.Log.Error(err)
		return
	}
	sqlStmt = `CREATE TABLE location_predictions (timestamp integer NOT NULL PRIMARY KEY, prediction TEXT, UNIQUE(timestamp));`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		logger.Log.Error(err)
		return
	}
	sqlStmt = `CREATE TABLE devices (id TEXT PRIMARY KEY, name TEXT);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		logger.Log.Error(err)
		return
	}
	sqlStmt = `CREATE TABLE locations (id TEXT PRIMARY KEY, name TEXT);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		logger.Log.Error(err)
		return
	}

	sqlStmt = `CREATE TABLE gps (id INTEGER PRIMARY KEY, timestamp INTEGER, mac TEXT, loc TEXT, lat REAL, lon REAL, alt REAL);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		logger.Log.Error(err)
		return
	}

	sqlStmt = `create index devices_name on devices (name);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		logger.Log.Error(err)
		return
	}

	sqlStmt = `CREATE INDEX sensors_devices ON sensors (deviceid);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		logger.Log.Error(err)
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
	// logger.Log.Debugf("got %s from '%s'", string(result), key)
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

	// logger.Log.Debugf("set '%s' to '%s'", key, string(b))
	return
}

// Dump will output the string version of the database
func (d *Database) Dump() (dumped string, err error) {
	var b bytes.Buffer
	out := bufio.NewWriter(&b)
	err = sqlite3dump.Dump(d.name, out)
	if err != nil {
		return
	}
	out.Flush()
	dumped = string(b.Bytes())
	return
}

// GetAllFingerprints returns all the fingerprints
func (d *Database) GetAllFingerprints() (s []models.SensorData, err error) {
	return d.GetAllFromQuery("SELECT * FROM sensors ORDER BY timestamp")
}

// AddPrediction will insert or update a prediction in the database
func (d *Database) AddPrediction(timestamp int64, aidata []models.LocationPrediction) (err error) {
	// make sure we have a prediction
	if len(aidata) == 0 {
		err = errors.New("no predictions to add")
		return
	}

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
		return errors.Wrap(err, "begin AddPrediction")
	}
	stmt, err := tx.Prepare("insert or replace into location_predictions (timestamp,prediction) values (?, ?)")
	if err != nil {
		return errors.Wrap(err, "stmt AddPrediction")
	}
	defer stmt.Close()

	_, err = stmt.Exec(timestamp, string(b))
	if err != nil {
		return errors.Wrap(err, "exec AddPrediction")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "commit AddPrediction")
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
	// logger.Log.Debugf("got %s from '%s'", string(result), key)
	return
}

// AddSensor will insert a sensor data into the database
// TODO: AddSensor should be special case of AddSensors
func (d *Database) AddSensor(s models.SensorData) (err error) {
	startTime := time.Now()
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
	previousCurrent := sensorDataSS.Current

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
			logger.Log.Debugf("adding column %s", sensor)
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
	// logger.Log.Debug("columns", columnList)
	// logger.Log.Debug("args", args)
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
	if previousCurrent != sensorDataSS.Current {
		err = d.Set("sensorDataStringSizer", sensorDataSS.Save())
		if err != nil {
			return
		}
	}

	logger.Log.Debugf("[%s] inserted sensor data, %s", s.Family, time.Since(startTime))
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

// GetLastSensorTimestamp gets will retrieve the value associated with a key.
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

// TotalLearnedCount gets will retrieve the value associated with a key.
func (d *Database) TotalLearnedCount() (count int64, err error) {
	stmt, err := d.db.Prepare("SELECT count(timestamp) FROM sensors WHERE locationid != ''")
	if err != nil {
		err = errors.Wrap(err, "problem preparing SQL")
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&count)
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
	logger.Log.Debugf("using minimum timestamp of %d", minimumTimestamp)
	sensors, err = d.GetAllFromPreparedQuery("SELECT * FROM sensors WHERE timestamp > ? GROUP BY deviceid ORDER BY timestamp DESC", minimumTimestamp)
	return
}

func (d *Database) NumDevices() (num int, err error) {
	stmt, err := d.db.Prepare("select count(id) from devices")
	if err != nil {
		err = errors.Wrap(err, "problem preparing SQL")
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&num)
	if err != nil {
		err = errors.Wrap(err, "problem getting key")
	}
	return
}

func (d *Database) GetDeviceFirstTimeFromDevices(devices []string) (firstTime map[string]time.Time, err error) {
	firstTime = make(map[string]time.Time)
	query := fmt.Sprintf("select n,t from (select devices.name as n,sensors.timestamp as t from sensors inner join devices on sensors.deviceid=devices.id WHERE devices.name IN ('%s') order by timestamp desc) group by n", strings.Join(devices, "','"))

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
		var name string
		var ts int64
		err = rows.Scan(&name, &ts)
		if err != nil {
			err = errors.Wrap(err, "scanning")
			return
		}
		// if _, ok := firstTime[name]; !ok {
		firstTime[name] = time.Unix(0, ts*1000000).UTC()
		// }
	}
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "rows")
	}
	return
}

func (d *Database) GetDeviceFirstTime() (firstTime map[string]time.Time, err error) {

	firstTime = make(map[string]time.Time)
	query := "select n,t from (select devices.name as n,sensors.timestamp as t from sensors inner join devices on sensors.deviceid=devices.id order by timestamp desc) group by n"
	// query := "select devices.name,sensors.timestamp from sensors inner join devices on sensors.deviceid=devices.id order by timestamp desc"
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
		var name string
		var ts int64
		err = rows.Scan(&name, &ts)
		if err != nil {
			err = errors.Wrap(err, "scanning")
			return
		}
		// if _, ok := firstTime[name]; !ok {
		firstTime[name] = time.Unix(0, ts*1000000).UTC()
		// }
	}
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "rows")
	}
	return
}

func (d *Database) GetDeviceCountsFromDevices(devices []string) (counts map[string]int, err error) {
	counts = make(map[string]int)
	query := fmt.Sprintf("select devices.name,count(sensors.timestamp) as num from sensors inner join devices on sensors.deviceid=devices.id WHERE devices.name in ('%s') group by sensors.deviceid", strings.Join(devices, "','"))
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
		var name string
		var count int
		err = rows.Scan(&name, &count)
		if err != nil {
			err = errors.Wrap(err, "scanning")
			return
		}
		counts[name] = count
	}
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "rows")
	}
	return
}

func (d *Database) GetDeviceCounts() (counts map[string]int, err error) {
	counts = make(map[string]int)
	query := "select devices.name,count(sensors.timestamp) as num from sensors inner join devices on sensors.deviceid=devices.id group by sensors.deviceid"
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
		var name string
		var count int
		err = rows.Scan(&name, &count)
		if err != nil {
			err = errors.Wrap(err, "scanning")
			return
		}
		counts[name] = count
	}
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "rows")
	}
	return
}

func (d *Database) GetLocationCounts() (counts map[string]int, err error) {
	counts = make(map[string]int)
	query := "SELECT locations.name,count(sensors.timestamp) as num from sensors inner join locations on sensors.locationid=locations.id group by sensors.locationid"
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
		var name string
		var count int
		err = rows.Scan(&name, &count)
		if err != nil {
			err = errors.Wrap(err, "scanning")
			return
		}
		counts[name] = count
	}
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "rows")
	}
	return
}

// GetAnalysisFromGreaterTime will return the analysis for a given timeframe
// func (d *Database) GetAnalysisFromGreaterTime(timestamp interface{}) {
// 	select sensors.timestamp, devices.name, location_predictions.prediction from sensors inner join location_predictions on location_predictions.timestamp=sensors.timestamp inner join devices on sensors.deviceid=devices.id WHERE sensors.timestamp > 0 GROUP BY devices.name ORDER BY sensors.timestamp DESC;
// }

// GetAllForClassification will return a sensor data for classifying
func (d *Database) GetAllForClassification() (s []models.SensorData, err error) {
	return d.GetAllFromQuery("SELECT * FROM sensors WHERE sensors.locationid !='' ORDER BY timestamp")
}

// GetAllNotForClassification will return a sensor data for classifying
func (d *Database) GetAllNotForClassification() (s []models.SensorData, err error) {
	return d.GetAllFromQuery("SELECT * FROM sensors WHERE sensors.locationid =='' ORDER BY timestamp")
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
		err = errors.Wrap(err, fmt.Sprintf("problem scanning rows, only got %d devices", len(devices)))
	}
	return
}

func (d *Database) GetLocations() (locations []string, err error) {
	// get only the active locations
	query := "SELECT locations.name FROM sensors INNER JOIN locations ON sensors.locationid=locations.id GROUP BY locations.name"
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

	locations = []string{}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			err = errors.Wrap(err, "scanning")
			return
		}
		locations = append(locations, name)
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

func (d *Database) DeleteLocation(locationName string) (err error) {
	id, err := d.GetID("locations", locationName)
	if err != nil {
		return
	}
	stmt, err := d.db.Prepare("DELETE FROM sensors WHERE locationid = ?")
	if err != nil {
		err = errors.Wrap(err, "problem preparing SQL")
		return

	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
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

// GetName will get the name of an element in a table (devices/locations) and return an error if it doesn't exist
func (d *Database) GetName(table string, id string) (name string, err error) {
	// first check to see if it has already been added
	stmt, err := d.db.Prepare("SELECT name FROM " + table + " WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		err = errors.Wrap(err, "problem preparing SQL")
		return
	}
	err = stmt.QueryRow(id).Scan(&name)
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

func (d *Database) Delete() (err error) {
	logger.Log.Debugf("deleting %s", d.family)
	return os.Remove(d.name)
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

	// if read-only, make sure the database exists
	if _, err = os.Stat(d.name); err != nil && len(readOnly) > 0 && readOnly[0] {
		err = errors.New(fmt.Sprintf("group '%s' does not exist", d.family))
		return
	}

	// obtain a lock on the database
	// logger.Log.Debugf("getting filelock on %s", d.name+".lock")
	for {
		var ok bool
		databaseLock.Lock()
		if _, ok = databaseLock.Locked[d.name]; !ok {
			databaseLock.Locked[d.name] = true
		}
		databaseLock.Unlock()
		if !ok {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	// logger.Log.Debugf("got filelock")

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
	// logger.Log.Debug("opened sqlite3 database")

	// create new database tables if needed
	if newDatabase {
		err = d.MakeTables()
		if err != nil {
			return
		}
		logger.Log.Debug("made tables")
	}

	return
}

func (d *Database) Debug(debugMode bool) {
	if debugMode {
		logger.SetLevel("debug")
	} else {
		logger.SetLevel("info")
	}
}

// Close will close the database connection and remove the filelock.
func (d *Database) Close() (err error) {
	if d.isClosed {
		return
	}
	// close database
	err2 := d.db.Close()
	if err2 != nil {
		err = err2
		logger.Log.Error(err)
	}

	// close filelock
	// logger.Log.Debug("closing lock")
	databaseLock.Lock()
	delete(databaseLock.Locked, d.name)
	databaseLock.Unlock()
	d.isClosed = true
	return
}

func (d *Database) GetAllFromQuery(query string) (s []models.SensorData, err error) {
	// logger.Log.Debug(query)
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
	// startQuery := time.Now()
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
	// logger.Log.Debugf("%s: %s", query, time.Since(startQuery))
	// startQuery = time.Now()
	defer rows.Close()
	s, err = d.getRows(rows)
	if err != nil {
		err = errors.Wrap(err, query)
	}
	// logger.Log.Debugf("getRows %s: %s", query, time.Since(startQuery))
	return
}

func (d *Database) getRows(rows *sql.Rows) (s []models.SensorData, err error) {
	// first get the columns
	logger.Log.Debug("getting columns")
	columnList, err := d.Columns()
	if err != nil {
		return
	}

	// get the string sizer for the sensor data
	logger.Log.Debug("getting sensorstringsizer")
	var sensorDataStringSizerString string
	err = d.Get("sensorDataStringSizer", &sensorDataStringSizerString)
	if err != nil {
		return
	}
	sensorDataSS, err := stringsizer.New(sensorDataStringSizerString)
	if err != nil {
		return
	}

	// logger.Log.Debug("getting devices")
	// deviceIDToName, err := d.GetIDToName("devices")
	// if err != nil {
	// 	return
	// }

	logger.Log.Debug("getting locations")
	locationIDToName, err := d.GetIDToName("locations")
	if err != nil {
		return
	}
	logger.Log.Debug("got locations")

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
		deviceID := string((*arr[1].(*interface{})).([]uint8))
		s0 := models.SensorData{
			// the underlying value of the interface pointer and cast it to a pointer interface to cast to a byte to cast to a string
			Timestamp: int64((*arr[0].(*interface{})).(int64)),
			Family:    d.family,
			Device:    deviceID,
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

	for i := range s {
		deviceName, errFind := d.GetName("devices", s[i].Device)
		if errFind != nil {
			err = errors.Wrap(errFind, "can't get name of "+s[i].Device)
			logger.Log.Error(err)
			continue
		}
		s[i].Device = deviceName
	}
	return
}

// SetGPS will set a GPS value in the GPS database
func (d *Database) SetGPS(p models.SensorData) (err error) {
	tx, err := d.db.Begin()
	if err != nil {
		return errors.Wrap(err, "SetGPS")
	}
	stmt, err := tx.Prepare("insert or replace into gps(timestamp ,mac, loc, lat, lon, alt) values (?, ?, ?, ?, ?,?)")
	if err != nil {
		return errors.Wrap(err, "SetGPS")
	}
	defer stmt.Close()

	for sensorType := range p.Sensors {
		for mac := range p.Sensors[sensorType] {
			_, err = stmt.Exec(p.Timestamp, sensorType+"-"+mac, p.Location, p.GPS.Latitude, p.GPS.Longitude, p.GPS.Altitude)
			if err != nil {
				return errors.Wrap(err, "SetGPS")
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "SetGPS")
	}
	return
}

// GetAverageGPS will return an average wardrived GPS
// for a given location
func (d *Database) GetAverageGPS(location string) (lat float64, lon float64, err error) {
	query := "SELECT avg(lat),avg(lon) FROM gps WHERE loc == ?"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		err = errors.Wrap(err, query)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(location)
	if err != nil {
		err = errors.Wrap(err, query)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&lat, &lon)
		if err != nil {
			err = errors.Wrap(err, "scanning")
			return
		}
	}
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "rows")
	}
	return
}

// // GetGPS will return a GPS for a given mac, if it exists
// // if it doesn't exist it will return an error
// func (d *Database) GetGPS(mac string) (gps models.GPS, err error) {
// 	query := "SELECT mac,lat,lon,alt,timestamp FROM gps WHERE mac == ?"
// 	stmt, err := d.db.Prepare(query)
// 	if err != nil {
// 		err = errors.Wrap(err, query)
// 		return
// 	}
// 	defer stmt.Close()
// 	rows, err := stmt.Query(mac)
// 	if err != nil {
// 		err = errors.Wrap(err, query)
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		err = rows.Scan(&gps.Mac, &gps.Latitude, &gps.Longitude, &gps.Altitude, &gps.Timestamp)
// 		if err != nil {
// 			err = errors.Wrap(err, "scanning")
// 			return
// 		}
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		err = errors.Wrap(err, "rows")
// 	}
// 	if gps.Mac == "" {
// 		err = errors.New(mac + " does not exist in gps table")
// 	}
// 	return
// }

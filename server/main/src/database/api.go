package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mr-tron/base58/base58"
	"github.com/pkg/errors"
	"github.com/schollz/find3/server/main/src/models"
	"github.com/schollz/stringsizer"
)

// MakeTables creates two tables, a `keystore` table:
//
// 	KEY (TEXT)	VALUE (TEXT)
//
// and also a `sensors` table for the sensor data:
//
// 	TIMESTAMP (INTEGER)	FAMILY(TEXT)	DEVICE(TEXT)	LOCATOIN(TEXT)
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
	sqlStmt = `create table sensors (timestamp integer not null primary key, family text, device text, location text, unique(timestamp));`
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

	// save empty string sizers
	ss, _ := stringsizer.New()
	err = d.Set("sensorDataStringSizer", ss.Save())
	if err != nil {
		return
	}
	err = d.Set("deviceNameStringSizer", ss.Save())
	if err != nil {
		return
	}
	d.logger.Log.Debug("initiate map key shrinker")
	return
}

// Columns will list the columns
func (d *Database) Columns() (columns []string, err error) {
	rows, err := d.db.Query("select * from sensors limit 1")
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

// AddPrediction will insert or update a prediction in the database
func (d *Database) AddPrediction(timestamp int64, aidata models.LocationAnalysis) (err error) {
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
func (d *Database) GetPrediction(timestamp int64) (aidata models.LocationAnalysis, err error) {
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

	// get string sizers
	var sensorDataStringSizerString string
	err = d.Get("sensorDataStringSizer", &sensorDataStringSizerString)
	if err != nil {
		return
	}
	sensorDataSS, err := stringsizer.New(sensorDataStringSizerString)
	if err != nil {
		return
	}
	var deviceNameStringSizerString string
	err = d.Get("deviceNameStringSizer", &deviceNameStringSizerString)
	if err != nil {
		return
	}
	deviceNamesSS, err := stringsizer.New(deviceNameStringSizerString)
	if err != nil {
		return
	}

	// setup the database
	tx, err := d.db.Begin()
	if err != nil {
		return errors.Wrap(err, "AddSensor")
	}

	// first add new columns in the sensor data
	args := make([]interface{}, 4)
	args[0] = s.Timestamp
	args[1] = s.Family
	args[2] = deviceNamesSS.ShrinkString(s.Device)
	args[3] = s.Location
	argsQ := []string{"?", "?", "?", "?"}
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
		if i >= 4 {
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
		return errors.Wrap(err, "AddSensor, prepare")
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
	err = d.Set("deviceNameStringSizer", deviceNamesSS.Save())
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

// GetAllForClassification will return a sensor data for classifying
func (d *Database) GetAllForClassification() (s []models.SensorData, err error) {
	return d.GetAllFromQuery("SELECT * FROM sensors WHERE location !=''")
}

// GetLatest will return a sensor data for classifying
func (d *Database) GetLatest(device interface{}) (s models.SensorData, err error) {
	var sensors []models.SensorData
	sensors, err = d.GetAllFromPreparedQuery("SELECT * FROM sensors WHERE device = ? ORDER BY timestamp DESC LIMIT 1", device)
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
	query := "SELECT device FROM (SELECT device,COUNT(device) AS counts FROM sensors GROUP BY device) WHERE counts > 2 ORDER BY counts DESC;"
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
		var key string
		err = rows.Scan(&key)
		if err != nil {
			err = errors.Wrap(err, "scanning")
			return
		}
		devices = append(devices, key)
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

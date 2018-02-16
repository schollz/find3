package database

import (
	"database/sql"
	"fmt"
	"os"
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

func Exists(name string) (err error) {
	name = strings.TrimSpace(name)
	name = path.Join(DataFolder, base58.FastBase58Encoding([]byte(name))+".sqlite3.db")
	if _, err = os.Stat(name); err != nil {
		err = errors.New("database '" + name + "' does not exist")
	}
	return
}

// Open will open the database for transactions by first aquiring a filelock.
func Open(name string, readOnly ...bool) (d *Database, err error) {
	name = strings.TrimSpace(name)
	d = new(Database)

	// convert the name to base64 for file writing
	// override the name
	if len(readOnly) > 1 && readOnly[1] {
		d.name = path.Join(DataFolder, name)
	} else {
		d.name = path.Join(DataFolder, base58.FastBase58Encoding([]byte(name))+".sqlite3.db")
	}
	d.logger, err = logging.New()
	if err != nil {
		return
	}
	d.Debug(DebugMode)

	// if read-only, make sure the database exists
	if _, err = os.Stat(d.name); err != nil && len(readOnly) > 0 && readOnly[0] {
		err = errors.New(fmt.Sprintf("group '%s' does not exist", name))
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

		// save empty string sizers
		ss, _ := stringsizer.New()
		err = d.Set("sensorDataStringSizer", ss.Save())
		if err != nil {
			return
		}
		err = d.Set("deviceNamesStringSizer", ss.Save())
		if err != nil {
			return
		}
		d.logger.Log.Debug("initiate map key shrinker")
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
	d.logger.Log.Debug(query)
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
	ss, err := stringsizer.New(sensorDataStringSizerString)
	if err != nil {
		return
	}

	s = make([]models.SensorData, 100000)
	sI := 0
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
			Family:    string((*arr[1].(*interface{})).([]uint8)),
			Device:    string((*arr[2].(*interface{})).([]uint8)),
			Location:  string((*arr[3].(*interface{})).([]uint8)),
			Sensors:   make(map[string]map[string]interface{}),
		}
		// add in the sensor data
		for i, colName := range columnList {
			if i < 4 {
				continue
			}
			if *arr[i].(*interface{}) == nil {
				continue
			}
			shortenedJSON := string((*arr[i].(*interface{})).([]uint8))
			s0.Sensors[colName], err = ss.ExpandMapFromString(shortenedJSON)
			if err != nil {
				return
			}
		}
		s[sI] = s0
		sI++
	}
	s = s[:sI]
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "getRows")
	}
	return
}

package database

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"path"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/schollz/mapslimmer"
	log "github.com/sirupsen/logrus"
	flock "github.com/theckman/go-flock"
)

// Open will open the database for transactions by first aquiring a filelock.
func Open(name string, readOnly ...bool) (d *Database, err error) {
	d = new(Database)

	// convert the name to base64 for file writing
	d.name = path.Join(DataFolder, base64.URLEncoding.EncodeToString([]byte(name))+".sqlite3.db")
	d.logger = log.WithFields(log.Fields{
		"name": name + "(" + base64.URLEncoding.EncodeToString([]byte(name)) + ")",
	})

	// if read-only, make sure the database exists
	if _, err = os.Stat(d.name); err != nil && len(readOnly) > 0 && readOnly[0] {
		err = errors.New(fmt.Sprintf("group '%s' does not exist", name))
		return
	}

	// obtain a lock on the database
	d.fileLock = flock.NewFlock(d.name + ".lock")
	for {
		locked, err := d.fileLock.TryLock()
		if err == nil && locked {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	d.logger.Info("got filelock")

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
	d.logger.Info("opened sqlite3 database")

	// create new database tables if needed
	if newDatabase {
		err = d.MakeTables()
		if err != nil {
			return
		}
		d.logger.Info("made tables")

		ms, err2 := mapslimmer.Init()
		if err2 != nil {
			err = err2
			return
		}
		err = d.Set("slimmer", ms.Slimmer())
		d.logger.Info("initiate map key shrinker")
	}

	return
}

// Close will close the database connection and remove the filelock.
func (d *Database) Close() (err error) {
	// close filelock
	err = d.fileLock.Unlock()
	if err != nil {
		d.logger.Error(err)
	} else {
		os.Remove(d.name + ".lock")
		d.logger.Info("removed filelock")
	}

	// close database
	err2 := d.db.Close()
	if err2 != nil {
		err = err2
		d.logger.Error(err)
	} else {
		d.logger.Info("closed database")
	}
	return
}

func (d *Database) GetAllFromQuery(query string) (s []SensorData, err error) {
	d.logger.Debug(query)
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
func (d *Database) GetAllFromPreparedQuery(query string, args ...interface{}) (s []SensorData, err error) {
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

func (d *Database) getRows(rows *sql.Rows) (s []SensorData, err error) {
	// first get the columns
	columnList, err := d.Columns()
	if err != nil {
		return
	}

	// get the slimmer
	var slimmer string
	err = d.Get("slimmer", &slimmer)
	if err != nil {
		return
	}
	ms, err := mapslimmer.Init(slimmer)
	if err != nil {
		return
	}

	s = make([]SensorData, 100000)
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
		s0 := SensorData{
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
			unslimmed := string((*arr[i].(*interface{})).([]uint8))
			s0.Sensors[colName], err = ms.Loads(unslimmed)
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

package database

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"os"
	"path"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/schollz/mapslimmer"
	log "github.com/sirupsen/logrus"
	flock "github.com/theckman/go-flock"
)

// Open will open the database for transactions by first aquiring a filelock.
func Open(name string, needToAuthenticate ...bool) (d *Database, err error) {
	d = new(Database)

	// convert the name to base64 for file writing
	d.name = path.Join(DataFolder, base64.URLEncoding.EncodeToString([]byte(name))+".sqlite3.db")
	d.logger = log.WithFields(log.Fields{
		"name": name + "(" + base64.URLEncoding.EncodeToString([]byte(name)) + ")",
	})

	// TODO: Authenticate
	if len(needToAuthenticate) > 0 {
		if needToAuthenticate[0] {
			err = errors.New("authentication failed (doesn't exist)")
			if err != nil {
				return
			}
		}
	}
	d.logger.Info("authenticated")

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

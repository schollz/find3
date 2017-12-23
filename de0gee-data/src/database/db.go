package database

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"path"
	"time"

	_ "github.com/mattn/go-sqlite3"
	flock "github.com/theckman/go-flock"
)

// Init will initialize a database by authenticating (if needed).
func Init(name string, needToAuthenticate ...bool) (d *Database, err error) {
	d = new(Database)
	d.name = path.Join(DataFolder, base64.URLEncoding.EncodeToString([]byte(name))+".sqlite3.db")
	if len(needToAuthenticate) > 0 {
		if needToAuthenticate[0] {
			// TODO: Authenticate
			err = errors.New("authentication failed (doesn't exist)")
			if err != nil {
				return
			}
		}
	}
	return
}

// Open will open the database for transactions by first aquiring a filelock.
func (d *Database) Open() (err error) {
	// Try to obtain a lock on the database
	d.fileLock = flock.NewFlock(d.name + ".lock")
	for {
		locked, err := d.fileLock.TryLock()
		if err == nil && !locked {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	d.db, err = sql.Open("sqlite3", d.name)
	if err == nil {
		err = d.makeTables()
	}
	return
}

// Close will close the database connection and remove the filelock.
func (d *Database) Close() (err error) {
	d.fileLock.Unlock()
	return d.db.Close()
}

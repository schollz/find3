package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	flock "github.com/theckman/go-flock"
)

// DataFolder is set to where you want each Sqlite3 database to be stored
var DataFolder = "."

// Database is the main structure for holding the information
// pertaining to the name of the database.
type Database struct {
	name     string
	db       *sql.DB
	fileLock *flock.Flock
	logger   *log.Entry
}

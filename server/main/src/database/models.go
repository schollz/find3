package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/schollz/find3/server/main/src/logging"
	flock "github.com/theckman/go-flock"
)

// DataFolder is set to where you want each Sqlite3 database to be stored
var DataFolder = "."

// Database is the main structure for holding the information
// pertaining to the name of the database.
type Database struct {
	name     string
	family   string
	db       *sql.DB
	fileLock *flock.Flock
	logger   *logging.SeelogWrapper
}

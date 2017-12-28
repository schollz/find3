package database

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	flock "github.com/theckman/go-flock"
)

// DataFolder is set to where you want each Sqlite3 database to be stored
var DataFolder = "."
var AIPort = "8002"

// Database is the main structure for holding the information
// pertaining to the name of the database.
type Database struct {
	name     string
	db       *sql.DB
	fileLock *flock.Flock
	logger   *log.Entry
}

// SensorData is the typical data structure for storing sensor data.
type SensorData struct {
	Timestamp float64                           `json:"t"` // Timestamp is the unique identifier
	Family    string                            `json:"f"` // Family is a group of devices
	Device    string                            `json:"d"` // Devices are unique within a family
	Location  string                            `json:"l"` // Location is optional, used for classification
	Sensors   map[string]map[string]interface{} `json:"s"` // Sensors contains a map of map of sensor data
}

// Save will inserts the fingerprint into a database
func (d SensorData) Save() (err error) {
	if d.Family == "" {
		err = errors.New("family cannot be empty")
	} else if d.Device == "" {
		err = errors.New("device cannot be empty")
	} else if d.Timestamp <= 0 {
		err = errors.New("timestamp is not valid")
	}
	if err != nil {
		return
	}
	db, _ := Open(d.Family)
	defer db.Close()
	err = db.AddSensor(d)
	return
}

// FINDFingerprint is the prototypical information from the fingerprinting device
type FINDFingerprint struct {
	Group           string   `json:"group"`
	Username        string   `json:"username"`
	Location        string   `json:"location"`
	Timestamp       int64    `json:"timestamp"`
	WifiFingerprint []Router `json:"wifi-fingerprint"`
}

// Router is the router information for each invdividual mac address
type Router struct {
	Mac  string `json:"mac"`
	Rssi int    `json:"rssi"`
}

// Convert will convert a FINDFingerprint into the new type of data,
// for backwards compatibility with FIND.
func (f FINDFingerprint) Convert() (d SensorData) {
	d = SensorData{
		Timestamp: float64(f.Timestamp),
		Family:    f.Group,
		Device:    f.Username,
		Location:  f.Location,
		Sensors:   make(map[string]map[string]interface{}),
	}
	if len(f.WifiFingerprint) > 0 {
		d.Sensors["wifi"] = make(map[string]interface{})
		for _, fingerprint := range f.WifiFingerprint {
			d.Sensors["wifi"][fingerprint.Mac] = float64(fingerprint.Rssi)
		}
	}
	return
}

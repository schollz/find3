package models

import (
	"errors"
)

// SensorData is the typical data structure for storing sensor data.
type SensorData struct {
	// Timestamp is the unique identifier, the time in milliseconds
	Timestamp int64 `json:"t"`
	// Family is a group of devices
	Family string `json:"f"`
	// Device are unique within a family
	Device string `json:"d"`
	// Location is optional, used for classification
	Location string `json:"l"`
	// Sensors contains a map of map of sensor data
	Sensors map[string]map[string]interface{} `json:"s"`
}

// Save will inserts the fingerprint into a database
func (d SensorData) Validate() (err error) {
	if d.Family == "" {
		err = errors.New("family cannot be empty")
	} else if d.Device == "" {
		err = errors.New("device cannot be empty")
	} else if d.Timestamp <= 0 {
		err = errors.New("timestamp is not valid")
	}
	if d.Timestamp == 1 {
		return
	}
	numFingerprints := 0
	for sensorType := range d.Sensors {
		numFingerprints += len(d.Sensors[sensorType])
	}
	if numFingerprints == 0 {
		err = errors.New("sensor data cannot be empty")
	}
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
		Timestamp: int64(f.Timestamp),
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

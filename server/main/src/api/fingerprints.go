package api

import (
	"sync"
	"time"

	"github.com/schollz/find3/server/main/src/database"
	"github.com/schollz/find3/server/main/src/models"
)

type UpdateCounterMap struct {
	// Data maps family -> counts of locations
	Count map[string]int
	sync.RWMutex
}

var globalUpdateCounter UpdateCounterMap

func init() {
	globalUpdateCounter.Lock()
	defer globalUpdateCounter.Unlock()
	globalUpdateCounter.Count = make(map[string]int)
}

// SaveSensorData will add sensor data to the database
func SaveSensorData(p models.SensorData) (err error) {
	err = p.Validate()
	if err != nil {
		return
	}
	db, err := database.Open(p.Family)
	if err != nil {
		return
	}
	err = db.AddSensor(p)
	db.Close()
	if err != nil {
		return
	}
	if p.Location != "" {
		go updateCounter(p.Family)
	}
	return
}

// SavePrediction will add sensor data to the database
func SavePrediction(s models.SensorData, p models.LocationAnalysis) (err error) {
	db, err := database.Open(s.Family)
	if err != nil {
		return
	}
	err = db.AddPrediction(s.Timestamp, p.Guesses)
	db.Close()
	return
}

// HasGPS returns true if any of the specified mac addresses has
// a GPS coordinate in the database
func HasGPS(fingerprint models.SensorData) (yes bool, err error) {
	db, err := database.Open(fingerprint.Family)
	if err != nil {
		return
	}
	defer db.Close()

	for sensorType := range fingerprint.Sensors {
		for mac := range fingerprint.Sensors[sensorType] {
			_, errGet := db.GetGPS(sensorType + "-" + mac)
			if errGet == nil {
				yes = true
				return
			}
		}
	}
	return
}

// AddGPSData will add GPS data to the database
func AddGPSData(data models.FingerprintWithGPS) (err error) {
	logger.Log.Debugf("adding GPS %+v", data)
	db, err := database.Open(data.Fingerprint.Family)
	if err != nil {
		return
	}
	defer db.Close()

	for sensorType := range data.Fingerprint.Sensors {
		for mac := range data.Fingerprint.Sensors[sensorType] {
			logger.Log.Debugf("adding gps data for %s", mac)
			err = db.SetGPS(models.GPS{
				Longitude: data.GPS.Longitude,
				Latitude:  data.GPS.Latitude,
				Altitude:  data.GPS.Altitude,
				Timestamp: data.Fingerprint.Timestamp,
				Mac:       sensorType + "-" + mac,
			})
			if err != nil {
				return
			}
		}
	}
	return
}

func updateCounter(family string) {
	globalUpdateCounter.Lock()
	if _, ok := globalUpdateCounter.Count[family]; !ok {
		globalUpdateCounter.Count[family] = 0
	}
	globalUpdateCounter.Count[family]++
	count := globalUpdateCounter.Count[family]
	globalUpdateCounter.Unlock()

	logger.Log.Debugf("'%s' has %d new fingerprints", family, count)
	if count < 5 {
		return
	}
	db, err := database.Open(family)
	if err != nil {
		return
	}
	var lastCalibrationTime time.Time
	err = db.Get("LastCalibrationTime", &lastCalibrationTime)
	defer db.Close()
	if err != nil {
		return
	}
	if time.Since(lastCalibrationTime) < 5*time.Minute && count < 20 {
		return
	}
	logger.Log.Infof("have %d new fingerprints for '%s', re-calibrating since last calibration was %s", count, family, time.Since(lastCalibrationTime))
	globalUpdateCounter.Lock()
	globalUpdateCounter.Count[family] = 0
	globalUpdateCounter.Unlock()

	// debounce the calibration time
	err = db.Set("LastCalibrationTime", time.Now().UTC())
	if err != nil {
		logger.Log.Error(err)
	}

	go Calibrate(family, true)
}

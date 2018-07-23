package api

import (
	"encoding/json"
	"net/http"
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

	// add GPS if it exists
	if p.Location != "" {
		if p.GPS.Longitude != 0 && p.GPS.Latitude != 0 {
			db.SetGPS(p)
		}
	}
	db.Close()

	// add GPS for each mac
	go AddGPSForMacs(p)

	if err != nil {
		return
	}

	if p.Location != "" {
		go updateCounter(p.Family)
	}
	return
}

// AddGPSForMacs will add the GPS for the given macs by attempting
// to find the GPS coordinates of the mac address.
func AddGPSForMacs(s models.SensorData) {
	if s.Location == "" {
		return
	}
	if _, ok := s.Sensors["wifi"]; !ok {
		return
	}
	for device := range s.Sensors["wifi"] {
		lat, lon := func() (lat, lon float64) {
			db, err := database.Open(s.Family)
			if err != nil {
				return
			}
			_, errGet := db.GetGPS("wifi-" + device)
			db.Close()
			if errGet == nil {
				return
			} else {
				logger.Log.Debug(errGet.Error())
			}
			type MacData struct {
				Ready      bool    `json:"ready"`
				MacAddress string  `json:"mac"`
				Exists     bool    `json:"exists"`
				Latitude   float64 `json:"lat,omitempty"`
				Longitude  float64 `json:"lon,omitempty"`
				Error      string  `json:"err,omitempty"`
			}
			var md MacData
			resp, err := http.Get("https://mac2gps.schollz.com/" + device)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			err = json.NewDecoder(resp.Body).Decode(&md)
			if err != nil {
				return
			}
			if md.Ready && md.Exists {
				lat = md.Latitude
				lon = md.Longitude
				logger.Log.Debugf("found GPS: %+v", md)
			}
			return
		}()
		if lat != 0 && lon != 0 {
			db, err := database.Open(s.Family)
			if err != nil {
				return
			}
			logger.Log.Debugf("[%s] setting GPS for %s@%s: %2.5f,%2.5f", s.Family, device, s.Location, lat, lon)
			err = db.SetGPSForMac(s.Location, device, lat, lon)
			db.Close()
			if err == nil {
				logger.Log.Error(err)
			}
		}
	}
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
	if err == nil {
		if time.Since(lastCalibrationTime) < 5*time.Minute {
			return
		}
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

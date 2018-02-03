package api

import (
	"github.com/de0gee/de0gee-data/src/database"
	"github.com/de0gee/de0gee-data/src/models"
)

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
	defer db.Close()
	err = db.AddSensor(p)
	return
}

package api

import (
	"github.com/pkg/errors"
	"github.com/schollz/find3/server/main/src/database"
	"github.com/schollz/find3/server/main/src/models"
)

// GetGPSData returns the latest GPS data
func GetGPSData(family string) (gpsData map[string]models.SensorData, err error) {
	gpsData = make(map[string]models.SensorData)

	d, err := database.Open(family, true)
	if err != nil {
		err = errors.Wrap(err, "You need to add learning data first")
		return
	}
	defer d.Close()

	locations, err := d.GetLocations()
	if err != nil {
		err = errors.Wrap(err, "problem getting locations")
		return
	}

	// initialize
	for _, location := range locations {
		gpsData[location] = models.SensorData{
			GPS: models.GPS{
				Latitude:  -1,
				Longitude: -1,
			},
		}
	}

	// get auto GPS data
	var autoGPS map[string]models.SensorData
	errGet := d.Get("autoGPS", &autoGPS)
	if errGet == nil {
		for location := range autoGPS {
			gpsData[location] = models.SensorData{
				GPS: models.GPS{
					Latitude:  autoGPS[location].GPS.Latitude,
					Longitude: autoGPS[location].GPS.Longitude,
				},
			}
		}
	}

	// get custom GPS data and override gpsdata
	var customGPS map[string]models.SensorData
	errGet = d.Get("customGPS", &customGPS)
	if errGet == nil {
		for location := range customGPS {
			gpsData[location] = models.SensorData{
				GPS: models.GPS{
					Latitude:  customGPS[location].GPS.Latitude,
					Longitude: customGPS[location].GPS.Longitude,
				},
			}
		}
	}

	return
}

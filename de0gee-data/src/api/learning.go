package api

import (
	"fmt"
	"os"
	"strings"

	"github.com/de0gee/de0gee-data/src/models"
)

func dumpSensorsToCSV(datas []models.SensorData, csvFile string) (err error) {
	logger.Log.Infof("dumping %d fingerprints to %s", len(datas), csvFile)
	// open CSV file for writing
	f, err := os.Create(csvFile)
	if err != nil {
		return
	}
	defer f.Close()

	// determine all possible columns
	sensorColumns := make(map[string]int)
	columnCount := 1
	for _, data := range datas {
		for sensorType := range data.Sensors {
			for sensorName := range data.Sensors[sensorType] {
				name := fmt.Sprintf("%s-%s", sensorType, sensorName)
				if _, ok := sensorColumns[name]; !ok {
					sensorColumns[name] = columnCount
					columnCount++
				}
			}
		}
	}

	// get column names
	columns := make([]string, columnCount)
	columns[0] = "location"
	for column := range sensorColumns {
		columns[sensorColumns[column]] = column
	}
	f.WriteString(strings.Join(columns, ",") + "\n")

	for _, data := range datas {
		columns = make([]string, columnCount)
		columns[0] = data.Location
		for sensorType := range data.Sensors {
			for sensorName := range data.Sensors[sensorType] {
				columns[sensorColumns[fmt.Sprintf("%s-%s", sensorType, sensorName)]] = fmt.Sprintf("%3.9f", data.Sensors[sensorType][sensorName])
			}
		}
		f.WriteString(strings.Join(columns, ",") + "\n")
	}

	return
}

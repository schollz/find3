package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/de0gee/de0gee-data/src/database"
	"github.com/de0gee/de0gee-data/src/models"
	"github.com/de0gee/de0gee-data/src/utils"
)

// Calibrate will send the sensor data for a specific family to the machine learning algorithms
func Calibrate(family string) (err error) {

	// inquire the AI
	type Payload struct {
		Family     string `json:"family"`
		CSVFile    string `json:"csv_file"`
		DataFolder string `json:"data_folder"`
	}
	var p Payload
	p.CSVFile = utils.RandomString(8) + ".csv"
	p.Family = family
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	p.DataFolder = dir

	// gather the data
	db, err := database.Open(family)
	if err != nil {
		return
	}
	datas, err := db.GetAllForClassification()
	if err != nil {
		return
	}
	db.Close()

	logger.Log.Debugf("writing %s data to %s", family, path.Join(p.DataFolder, p.CSVFile))

	err = dumpSensorsToCSV(datas, path.Join(p.DataFolder, p.CSVFile))
	if err != nil {
		return
	}
	defer os.Remove(path.Join(p.DataFolder, p.CSVFile))

	url := "http://localhost:" + AIPort + "/learn"
	bPayload, err := json.Marshal(p)
	if err != nil {
		return
	}
	logger.Log.Debugf("sending payload: %s", bPayload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bPayload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var target AnalysisResponse
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		return
	}

	if target.Success {
		logger.Log.Debugf("success: %s", target.Message)
	} else {
		logger.Log.Debugf("failure: %s", target.Message)
		err = errors.New("failed in AI server: " + target.Message)
	}
	return
}

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

package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/de0gee/de0gee-data/src/database"
	"github.com/de0gee/de0gee-data/src/models"
	"github.com/de0gee/de0gee-data/src/utils"
)

// Calibrate will send the sensor data for a specific family to the machine learning algorithms
func Calibrate(family string, crossValidation ...bool) (err error) {

	// inquire the AI
	type Payload struct {
		Family     string `json:"family"`
		CSVFile    string `json:"csv_file"`
		DataFolder string `json:"data_folder"`
	}
	var p Payload
	p.CSVFile = utils.RandomString(8) + ".csv"
	p.Family = family
	p.DataFolder = DataFolder

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

	// for cross validation only
	var datasTest []models.SensorData
	if len(crossValidation) > 0 && crossValidation[0] {
		// randomize data order
		for i := range datas {
			j := rand.Intn(i + 1)
			datas[i], datas[j] = datas[j], datas[i]
		}

		// split the data to use 70% to learn, 30% to test
		splitI := int(0.7 * float64(len(datas)))
		datasTest = datas[splitI:]
		datas = datas[:splitI]
		logger.Log.Debugf("splitting data for cross validation (%d -> %d)", len(datas), splitI)
	}

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
	if err != nil {
		return
	}

	if len(crossValidation) > 0 && crossValidation[0] {
		go FindBestAlgorithm(datasTest)
	}
	return
}

func FindBestAlgorithm(datas []models.SensorData) (err error) {
	predictionAnalysis := make(map[string]map[string]map[string]int)
	logger.Log.Debugf("finding best algorithm for %d data", len(datas))
	aidatas := make([]models.LocationAnalysis, len(datas))
	for i := range datas {
		t := time.Now()
		aidatas[i], err = AnalyzeSensorData(datas[i])
		if err != nil {
			return
		}
		logger.Log.Debugf("got analysis for %d/%d in %s", i, len(datas), time.Since(t))
	}

	for i, aidata := range aidatas {
		for _, prediction := range aidata.Predictions {
			if _, ok := predictionAnalysis[prediction.Name]; !ok {
				predictionAnalysis[prediction.Name] = make(map[string]map[string]int)
				for trueLoc := range aidata.LocationNames {
					predictionAnalysis[prediction.Name][aidata.LocationNames[trueLoc]] = make(map[string]int)
					for guessLoc := range aidata.LocationNames {
						predictionAnalysis[prediction.Name][aidata.LocationNames[trueLoc]][aidata.LocationNames[guessLoc]] = 0
					}
				}
			}
			correctLocation := datas[i].Location
			guessedLocation := aidata.LocationNames[prediction.Locations[0]]
			predictionAnalysis[prediction.Name][correctLocation][guessedLocation]++
		}
	}

	// normalize prediction analysis
	// initialize location totals
	locationTotals := make(map[string]int)
	for _, data := range datas {
		if _, ok := locationTotals[data.Location]; !ok {
			locationTotals[data.Location] = 0
		}
		locationTotals[data.Location]++
	}
	algorithmEfficacy := make(map[string]map[string]BinaryStats)
	for alg := range predictionAnalysis {
		if _, ok := algorithmEfficacy[alg]; !ok {
			algorithmEfficacy[alg] = make(map[string]BinaryStats)
		}
		// calculate true/false positives/negatives
		tp := 0
		fp := 0
		tn := 0
		fn := 0
		for correctLocation := range predictionAnalysis[alg] {
			for l1 := range predictionAnalysis[alg] {
				for l2 := range predictionAnalysis[alg] {
					if correctLocation == l1 && correctLocation == l2 {
						tp = predictionAnalysis[alg][l1][l2]
					} else if correctLocation == l1 && correctLocation != l2 {
						fp += predictionAnalysis[alg][l1][l2]
					} else if correctLocation != l1 && correctLocation == l2 {
						fn += predictionAnalysis[alg][l1][l2]
					} else if correctLocation != l1 && correctLocation != l2 {
						tn += predictionAnalysis[alg][l1][l2]
					}
				}
			}
			algorithmEfficacy[alg][correctLocation] = NewBinaryStats(tp, fp, tn, fn)
		}
	}

	logger.Log.Debugf("determining best guess for %d datas", len(aidatas))
	correct := 0
	for i := range aidatas {
		bestGuess := determineBestGuess(aidatas[i], algorithmEfficacy)
		if bestGuess.Location == datas[i].Location {
			correct++
		}
	}
	logger.Log.Infof("correct: %d/%d", correct, len(aidatas))

	logger.Log.Infof("algorithmEfficacy: %+v", algorithmEfficacy)
	// gather the data
	db, err := database.Open(datas[0].Family)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	defer db.Close()
	err = db.Set("PercentCorrect", float64(correct)/float64(len(datas)))
	if err != nil {
		logger.Log.Error(err)
	}
	err = db.Set("PredictionAnalysis", predictionAnalysis)
	if err != nil {
		logger.Log.Error(err)
	}
	err = db.Set("AlgorithmEfficacy", algorithmEfficacy)
	if err != nil {
		logger.Log.Error(err)
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

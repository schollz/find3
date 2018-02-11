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
	var aidata models.LocationAnalysis
	predictionScores := make(map[string]int)
	predictionAnalysis := make(map[string]map[string]map[string]int)
	logger.Log.Debugf("finding best algorithm for %d data", len(datas))

	for _, j := range rand.Perm(len(datas)) {
		data := datas[j]
		// if i > 20 {
		// 	break
		// }
		aidata, err = AnalyzeSensorData(data)
		if err != nil {
			return
		}
		for _, prediction := range aidata.Predictions {
			if _, ok := predictionScores[prediction.Name]; !ok {
				predictionScores[prediction.Name] = 0
				predictionAnalysis[prediction.Name] = make(map[string]map[string]int)
				for trueLoc := range aidata.LocationNames {
					predictionAnalysis[prediction.Name][aidata.LocationNames[trueLoc]] = make(map[string]int)
					for guessLoc := range aidata.LocationNames {
						predictionAnalysis[prediction.Name][aidata.LocationNames[trueLoc]][aidata.LocationNames[guessLoc]] = 0
					}
				}
			}
			correctLocation := data.Location
			guessedLocation := aidata.LocationNames[prediction.Locations[0]]
			predictionAnalysis[prediction.Name][correctLocation][guessedLocation]++
			logger.Log.Debugf("%s|%s|%s", prediction.Name, guessedLocation, correctLocation)
			if guessedLocation == correctLocation {
				predictionScores[prediction.Name]++
			} else {
				if prediction.Name == "Extended Naive Bayes" {

				}
			}
		}
	}
	logger.Log.Debugf("prediction scores: %+v", predictionScores)
	bestScoreMethod := ""
	bestScore := 0
	for prediction := range predictionScores {
		if predictionScores[prediction] > bestScore {
			bestScore = predictionScores[prediction]
			bestScoreMethod = prediction
		}
	}
	predictionScores["total"] = len(datas)

	// normalize prediction analysis
	// initialize location totals
	locationTotals := make(map[string]int)
	for _, data := range datas {
		if _, ok := locationTotals[data.Location]; !ok {
			locationTotals[data.Location] = 0
		}
		locationTotals[data.Location]++
	}
	algorithmEfficacy := make(map[string]map[string]float64)
	for alg := range predictionAnalysis {
		if _, ok := algorithmEfficacy[alg]; !ok {
			algorithmEfficacy[alg] = make(map[string]float64)
		}
		for loc := range predictionAnalysis[alg] {
			percentageRight := float64(predictionAnalysis[alg][loc][loc]) / float64(locationTotals[loc])
			// true positive rate
			tpr := percentageRight
			// false positive rate
			fpr := float64(0)
			total := 0
			for notLoc := range predictionAnalysis[alg][loc] {
				if notLoc == loc {
					continue
				}
				total += locationTotals[notLoc]
				fpr += float64(predictionAnalysis[alg][notLoc][loc])
			}
			fpr = fpr / float64(total)
			algorithmEfficacy[alg][loc] = tpr - fpr
		}
	}

	// gather the data
	db, err := database.Open(datas[0].Family)
	if err != nil {
		return
	}
	defer db.Close()
	err = db.Set("BestAlgorithm", bestScoreMethod)
	if err != nil {
		return
	}
	err = db.Set("BestAlgorithmData", predictionScores)
	if err != nil {
		return
	}
	err = db.Set("PredictionAnalysis", predictionAnalysis)
	if err != nil {
		return
	}
	err = db.Set("AlgorithmEfficacy", algorithmEfficacy)
	if err != nil {
		return
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

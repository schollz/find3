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

	"github.com/schollz/find3/server/main/src/database"
	"github.com/schollz/find3/server/main/src/models"
	"github.com/schollz/find3/server/main/src/utils"
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
	db, err := database.Open(family, true)
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

		// triage into different locations
		dataLocations := make(map[string][]int)
		for i := range datas {
			if _, ok := dataLocations[datas[i].Location]; !ok {
				dataLocations[datas[i].Location] = []int{}
			}
			dataLocations[datas[i].Location] = append(dataLocations[datas[i].Location], i)
		}

		// for each location, make test set and learn set
		datasTest = make([]models.SensorData, len(datas))
		datasTestI := 0
		datasLearn := make([]models.SensorData, len(datas))
		datasLearnI := 0
		for loc := range dataLocations {
			splitI := 1
			numDataPoints := len(dataLocations[loc])
			if numDataPoints < 2 {
				logger.Log.Warnf("[%s] not enough data to split %s", family, loc)
			} else if numDataPoints < 10 {
				splitI = numDataPoints / 2 // 50% split
			} else {
				splitI = numDataPoints * 7 / 10 // 70:30 split
			}
			for i, s := range dataLocations[loc] {
				if i < splitI {
					// used for learning
					datasLearn[datasLearnI] = datas[s]
					datasLearnI++
				} else {
					datasTest[datasTestI] = datas[s]
					datasTestI++
				}
			}
			logger.Log.Debugf("[%s] splitting %s data for cross validation (%d -> %d)", family, loc, numDataPoints, splitI)
		}

		datas = datasLearn[:datasLearnI]
		datasTest = datasTest[:datasTestI]
		logger.Log.Debugf("[%s] learning: %d, testing: %d", family, len(datas), len(datasTest))
	}

	logger.Log.Debugf("[%s] writing data to %s", family, path.Join(p.DataFolder, p.CSVFile))
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
	if len(datas) == 0 {
		err = errors.New("no data specified")
		return
	}
	predictionAnalysis := make(map[string]map[string]map[string]int)
	logger.Log.Debugf("[%s] finding best algorithm for %d data", datas[0].Family, len(datas))

	t := time.Now()
	type Job struct {
		data models.SensorData
		i    int
	}
	type Result struct {
		data models.LocationAnalysis
		i    int
	}
	jobs := make(chan Job, len(datas))
	results := make(chan Result, len(datas))
	workers := 9
	for w := 0; w < workers; w++ {
		go func(id int, jobs <-chan Job, results chan<- Result) {
			for job := range jobs {
				aidata, err := AnalyzeSensorData(job.data)
				if err != nil {
					logger.Log.Warnf("%s: %+v", err.Error(), job.data)
				}
				results <- Result{data: aidata, i: job.i}
			}
		}(w, jobs, results)
	}
	for i, data := range datas {
		jobs <- Job{data: data, i: i}
	}
	close(jobs)
	aidatas := make([]models.LocationAnalysis, len(datas))
	for i := 0; i < len(datas); i++ {
		result := <-results
		aidatas[result.i] = result.data
	}
	logger.Log.Infof("[%s] analyzed %d data in %s", datas[0].Family, len(datas), time.Since(t))

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
			if len(prediction.Locations) == 0 {
				logger.Log.Warn("prediction.Locations is empty!")
				continue
			}
			if len(aidata.LocationNames) == 0 {
				return errors.New("no location names")
			}
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
	algorithmEfficacy := make(map[string]map[string]models.BinaryStats)
	for alg := range predictionAnalysis {
		if _, ok := algorithmEfficacy[alg]; !ok {
			algorithmEfficacy[alg] = make(map[string]models.BinaryStats)
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
			algorithmEfficacy[alg][correctLocation] = models.NewBinaryStats(tp, fp, tn, fn)
		}
	}

	correct := 0
	ProbabilitiesOfBestGuess := make([]float64, len(aidatas))
	accuracyBreakdown := make(map[string]float64)
	accuracyBreakdownTotal := make(map[string]float64)
	for i := range aidatas {
		if _, ok := accuracyBreakdownTotal[datas[i].Location]; !ok {
			accuracyBreakdownTotal[datas[i].Location] = 0
			accuracyBreakdown[datas[i].Location] = 0
		}
		accuracyBreakdownTotal[datas[i].Location]++
		bestGuess := determineBestGuess(aidatas[i], algorithmEfficacy)
		if len(bestGuess) == 0 {
			continue
		}
		if bestGuess[0].Location == datas[i].Location {
			accuracyBreakdown[datas[i].Location]++
			correct++
			ProbabilitiesOfBestGuess[i] = bestGuess[0].Probability
		} else {
			ProbabilitiesOfBestGuess[i] = -1 * bestGuess[0].Probability
		}
	}
	logger.Log.Infof("[%s] total correct: %d/%d", datas[0].Family, correct, len(aidatas))

	for loc := range accuracyBreakdown {
		accuracyBreakdown[loc] = accuracyBreakdown[loc] / accuracyBreakdownTotal[loc]
		logger.Log.Infof("[%s] %s accuracy: %2.0f%%", datas[0].Family, loc, accuracyBreakdown[loc]*100)
	}

	// gather the data
	db, err := database.Open(datas[0].Family)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	defer db.Close()
	err = db.Set("ProbabilitiesOfBestGuess", ProbabilitiesOfBestGuess)
	if err != nil {
		logger.Log.Error(err)
	}
	err = db.Set("PercentCorrect", float64(correct)/float64(len(datas)))
	if err != nil {
		logger.Log.Error(err)
	}
	err = db.Set("AccuracyBreakdown", accuracyBreakdown)
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
	err = db.Set("LastCalibrationTime", time.Now())
	if err != nil {
		logger.Log.Error(err)
	}
	return
}

func dumpSensorsToCSV(datas []models.SensorData, csvFile string) (err error) {
	if len(datas) == 0 {
		err = errors.New("data is empty")
		return
	}
	logger.Log.Infof("[%s] dumping %d fingerprints to %s", datas[0].Family, len(datas), csvFile)
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

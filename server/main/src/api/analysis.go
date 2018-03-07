package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	cache "github.com/robfig/go-cache"
	"github.com/schollz/find3/server/main/src/database"
	"github.com/schollz/find3/server/main/src/models"
)

// AIPort designates the port for the AI processing
var AIPort = "8002"
var DataFolder = "."

var (
	httpClient *http.Client
	routeCache *cache.Cache
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 60
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
	routeCache = cache.New(5*time.Minute, 10*time.Minute)
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return client
}

type AnalysisResponse struct {
	Data    models.LocationAnalysis `json:"analysis"`
	Message string                  `json:"message"`
	Success bool                    `json:"success"`
}

func AnalyzeSensorData(s models.SensorData) (aidata models.LocationAnalysis, err error) {
	d, err := database.Open(s.Family)
	if err != nil {
		return
	}
	defer d.Close()

	// check if its already been classified
	// aidata, err = d.GetPrediction(s.Timestamp)
	// if err == nil {
	// 	return
	// }

	// inquire the AI
	var target AnalysisResponse
	type ClassifyPayload struct {
		Sensor     models.SensorData `json:"sensor_data"`
		DataFolder string            `json:"data_folder"`
	}
	var p2 ClassifyPayload
	p2.Sensor = s
	p2.DataFolder = DataFolder
	url := "http://localhost:" + AIPort + "/classify"
	bPayload, err := json.Marshal(p2)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bPayload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		return
	}
	if !target.Success {
		err = errors.New("unable to analyze: " + target.Message)
		return
	}
	if len(target.Data.Predictions) == 0 {
		err = errors.New("problem analyzing: no predictions")
		return
	}

	aidata = target.Data
	var algorithmEfficacy map[string]map[string]BinaryStats
	d.Get("AlgorithmEfficacy", &algorithmEfficacy)
	aidata.BestGuess = determineBestGuess(aidata, algorithmEfficacy)

	// add prediction to the database
	// adding predictions uses up a lot of space
	err = d.AddPrediction(s.Timestamp, aidata)

	return
}

func determineBestGuess(aidata models.LocationAnalysis, algorithmEfficacy map[string]map[string]BinaryStats) (b models.LocationPrediction) {
	bestEfficacy := float64(0)
	for _, prediction := range aidata.Predictions {
		if len(prediction.Locations) == 0 {
			continue
		}
		guessedLocation := aidata.LocationNames[prediction.Locations[0]]
		efficacy := prediction.Probabilities[0] * algorithmEfficacy[prediction.Name][guessedLocation].Informedness
		if efficacy > bestEfficacy {
			bestEfficacy = efficacy
			b.Location = guessedLocation
			b.Name = prediction.Name
			b.Probability = bestEfficacy
		}
	}

	// determine consensus
	locationScores := make(map[string]float64)
	for _, prediction := range aidata.Predictions {
		if len(prediction.Locations) == 0 {
			continue
		}
		guessedLocation := aidata.LocationNames[prediction.Locations[0]]
		efficacy := prediction.Probabilities[0] * algorithmEfficacy[prediction.Name][guessedLocation].Informedness
		if _, ok := locationScores[guessedLocation]; !ok {
			locationScores[guessedLocation] = float64(0)
		}
		if efficacy > 0 {
			locationScores[guessedLocation] += efficacy
		}
	}
	logger.Log.Infof("consensus: %+v", locationScores)

	b.Probability = 0
	b.Name = "consensus"
	total := float64(0)
	for key := range locationScores {
		total += locationScores[key]
	}
	for key := range locationScores {
		if locationScores[key]/total > b.Probability {
			b.Probability = locationScores[key] / total
			b.Location = key
		}
	}
	return b
}

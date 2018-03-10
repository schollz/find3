package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"sort"
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
	var algorithmEfficacy map[string]map[string]models.BinaryStats
	d.Get("AlgorithmEfficacy", &algorithmEfficacy)
	aidata.Guesses = determineBestGuess(aidata, algorithmEfficacy)

	// add prediction to the database
	// adding predictions uses up a lot of space
	err = d.AddPrediction(s.Timestamp, aidata.Guesses)

	return
}

func determineBestGuess(aidata models.LocationAnalysis, algorithmEfficacy map[string]map[string]models.BinaryStats) (b []models.LocationPrediction) {
	// determine consensus
	locationScores := make(map[string]float64)
	for _, prediction := range aidata.Predictions {
		if len(prediction.Locations) == 0 {
			continue
		}
		for i := range prediction.Locations {
			guessedLocation := aidata.LocationNames[prediction.Locations[i]]
			if prediction.Probabilities[i] <= 0 {
				continue
			}
			if len(guessedLocation) == 0 {
				continue
			}
			efficacy := prediction.Probabilities[i] * algorithmEfficacy[prediction.Name][guessedLocation].Informedness
			if _, ok := locationScores[guessedLocation]; !ok {
				locationScores[guessedLocation] = float64(0)
			}
			if efficacy > 0 {
				locationScores[guessedLocation] += efficacy
			}
		}
	}

	total := float64(0)
	for location := range locationScores {
		total += locationScores[location]
	}

	pl := make(PairList, len(locationScores))
	i := 0
	for k, v := range locationScores {
		pl[i] = Pair{k, v / total}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	b = make([]models.LocationPrediction, len(locationScores))
	for i := range pl {
		b[i].Location = pl[i].Key
		b[i].Probability = pl[i].Value
	}

	if len(locationScores) == 1 {
		b[0].Probability = 1
	}

	return b
}

type Pair struct {
	Key   string
	Value float64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

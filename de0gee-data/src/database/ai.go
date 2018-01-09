package database

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	cache "github.com/robfig/go-cache"
)

var (
	httpClient *http.Client
	routeCache *cache.Cache
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
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

type AIResponse struct {
	Data    AIData `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type AIData struct {
	LocationNames map[string]string `json:"location_names"`
	Predictions   []struct {
		Locations     []string  `json:"locations"`
		Name          string    `json:"name"`
		Probabilities []float64 `json:"probabilities"`
	} `json:"predictions"`
}

func (d *Database) Classify(s SensorData) (aidata AIData, err error) {
	// check if its already been classified
	aidata, err = d.GetPrediction(s.Timestamp)
	if err == nil {
		return
	}

	// inquire the AI
	var target AIResponse
	type ClassifyPayload struct {
		Sensor       SensorData `json:"sensor_data"`
		DataLocation string     `json:"data_location"`
	}
	var p2 ClassifyPayload
	p2.Sensor = s
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	p2.DataLocation = dir
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
	aidata = target.Data

	// add prediction to the database
	err = d.AddPrediction(s.Timestamp, aidata)
	if err != nil {
		return
	}

	return
}

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	log "github.com/cihub/seelog"
	"github.com/de0gee/de0gee-data/src/models"
)

var (
	httpClient *http.Client
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
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

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func postData(payload models.SensorData, route string) (err error) {
	log.Debug("posting data")
	if len(payload.Sensors) == 0 {
		return errors.New("no sensor data")
	}
	url := "http://localhost:8003" + route
	bPayload, err := json.Marshal(payload)
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

	var target Response
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		return
	}
	if !target.Success {
		err = errors.New("unable to analyze: " + target.Message)
	}
	log.Debug(target)
	return
}

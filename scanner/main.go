package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/de0gee/de0gee-data/src/models"
)

func main() {
	log.Println("Testing")
	payload := models.SensorData{}
	payload.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	payload.Device = "dell"
	payload.Family = "test"
	payload.Location = "kitchen"
	payload.Sensors = make(map[string]map[string]interface{})
	wifiData := iw()
	fmt.Println(wifiData)
	// wifiData = iwlist()
	// fmt.Println(wifiData)
	if len(wifiData) > 0 {
		payload.Sensors["wifi"] = wifiData
	}
	// bluetoothData := scanBluetooth()
	// if len(bluetoothData) > 0 {
	// 	payload.Sensors["bluetooth"] = bluetoothData
	// }
	bPayload, err := json.MarshalIndent(payload, "", " ")
	fmt.Println(string(bPayload), err)
	err = postData(payload)
	if err != nil {
		log.Fatal(err)
	}
}

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

func postData(payload models.SensorData) (err error) {
	log.Println("posting data")
	url := "http://localhost:8003/data"
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
	log.Println(target)
	return
}

// when breath becomes air

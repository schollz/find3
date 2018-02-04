package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/de0gee/de0gee-data/src/models"
	"github.com/montanaflynn/stats"
)

func main2() {
	t := time.Now()
	s, _ := runCommand(60*time.Second, "l2ping", "-c", "300", "-f", "0C:3E:9F:28:22:6A")
	milliseconds := make([]float64, 300)
	i := 0
	for _, line := range strings.Split(s, "\n") {
		if !strings.Contains(line, "ms") {
			continue
		}
		lineSplit := strings.Fields(line)
		msString := strings.TrimRight(lineSplit[len(lineSplit)-1], "ms")
		ms, err := strconv.ParseFloat(msString, 64)
		if err != nil {
			log.Fatal(err)
		}
		milliseconds[i] = ms
		i++
	}
	milliseconds = milliseconds[:i]
	median, err := stats.Median(milliseconds)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(median)
	fmt.Println(time.Since(t) / 300)
}

func main() {
	payload := models.SensorData{}
	payload.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	payload.Device = "dell"
	payload.Family = "test"
	payload.Location = ""
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
	if len(payload.Sensors) == 0 {
		log.Fatal(errors.New("collected no data"))
	}
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

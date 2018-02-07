package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	log "github.com/cihub/seelog"
	"github.com/de0gee/de0gee-data/src/models"
	"github.com/montanaflynn/stats"
)

func main() {
	defer log.Flush()
	setLogLevel("debug")
	log.Info("starting")
	// basicCapture()
	reverseCapture()
}

func reverseCapture() {
	sensors, err := ReverseScan("wlx98ded0151d38", "lf-testing2", "dell", 3*time.Second)
	if err != nil {
		log.Error(err)
		return
	}
	postData(sensors, "/reverse")
}

func basicCapture() {
	payload := models.SensorData{}
	payload.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	payload.Device = "dell"
	payload.Family = "test3"
	payload.Location = "kitchen table"
	payload.Sensors = make(map[string]map[string]interface{})
	wifiData := iw()
	fmt.Println(wifiData)
	// wifiData = iwlist()
	// fmt.Println(wifiData)
	if len(wifiData) > 0 {
		payload.Sensors["wifi"] = wifiData
	}
	bluetoothData := scanBluetooth()
	if len(bluetoothData) > 0 {
		payload.Sensors["bluetooth"] = bluetoothData
	}
	if len(payload.Sensors) == 0 {
		log.Error(errors.New("collected no data"))
	}
	bPayload, err := json.MarshalIndent(payload, "", " ")
	fmt.Println(string(bPayload), err)
	err = postData(payload, "/data")
	if err != nil {
		log.Error(err)
	}
}

// this doesn't work, just playing
func bluetoothTimeOfFlight() {
	t := time.Now()
	s, _ := RunCommand(60*time.Second, "l2ping -c 300 -f 0C:3E:9F:28:22:6A")
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
			log.Error(err)
		}
		milliseconds[i] = ms
		i++
	}
	milliseconds = milliseconds[:i]
	median, err := stats.Median(milliseconds)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(median)
	fmt.Println(time.Since(t) / 300)
}

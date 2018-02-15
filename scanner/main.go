package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	log "github.com/cihub/seelog"
	"github.com/montanaflynn/stats"
	"github.com/schollz/find3/server/main/src/models"
)

var (
	wifiInterface string

	server                   string
	family, device, location string

	scanSeconds            int
	doBluetooth            bool
	doReverse              bool
	doDebug                bool
	doSetPromiscuous       bool
	doNotModifyPromiscuity bool
	runForever             bool
)

func main() {
	defer log.Flush()
	flag.StringVar(&wifiInterface, "i", "wlan0", "wifi interface for scanning")
	flag.StringVar(&server, "server", "http://localhost:8003", "server to use")
	flag.StringVar(&family, "family", "", "family name")
	flag.StringVar(&device, "device", "", "device name")
	flag.StringVar(&location, "location", "", "location (optional)")
	flag.BoolVar(&doBluetooth, "bluetooth", false, "scan bluetooth")
	flag.BoolVar(&doReverse, "reverse", false, "reverse fingerprinting")
	flag.BoolVar(&doDebug, "debug", false, "enable debugging")
	flag.BoolVar(&doSetPromiscuous, "monitor-mode", false, "set promiscuous mode")
	flag.BoolVar(&doNotModifyPromiscuity, "no-modify", false, "disable changing wifi promiscuity mode")
	flag.BoolVar(&runForever, "forever", false, "run forever")
	flag.IntVar(&scanSeconds, "scantime", 3, "scan time")
	flag.Parse()

	if doDebug {
		setLogLevel("debug")
	} else {
		setLogLevel("info")
	}

	if device == "" {
		fmt.Println("device cannot be blank")
		flag.Usage()
		return
	}

	if doSetPromiscuous {
		PromiscuousMode(true)
		return
	}

	if family == "" {
		fmt.Println("family cannot be blank")
		flag.Usage()
		return
	}

	for {
		if !doReverse {
			log.Infof("scanning with %s", wifiInterface)
			basicCapture()
		} else {
			log.Infof("reverse scanning with %s", wifiInterface)
			reverseCapture()
		}
		if !runForever {
			break
		}
	}
}

func reverseCapture() {

	c := make(chan map[string]map[string]interface{})
	if doBluetooth {
		go scanBluetooth(c)
	}

	if !doNotModifyPromiscuity {
		PromiscuousMode(true)
		time.Sleep(1 * time.Second)
	}
	payload, err := ReverseScan(time.Duration(scanSeconds) * time.Second)
	if !doNotModifyPromiscuity {
		PromiscuousMode(false)
		time.Sleep(1 * time.Second)
	}
	if doBluetooth {
		data := <-c
		log.Debugf("bluetooth data:%+v", data)
		for sensor := range data {
			payload.Sensors[sensor] = make(map[string]interface{})
			for device := range data[sensor] {
				payload.Sensors[sensor][device] = data[sensor][device]
			}
		}
	}
	bSensors, _ := json.MarshalIndent(payload, "", " ")
	log.Debug(string(bSensors))

	if err != nil {
		log.Error(err)
		return
	}
	err = postData(payload, "/reverse")
	if err != nil {
		log.Error(err)
	}
}

func basicCapture() {
	payload := models.SensorData{}
	payload.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	payload.Family = family
	payload.Device = device
	payload.Location = location
	payload.Sensors = make(map[string]map[string]interface{})

	// collect sensors asynchronously
	c := make(chan map[string]map[string]interface{})
	numSensors := 0

	go iw(c)
	numSensors++

	if doBluetooth {
		go scanBluetooth(c)
		numSensors++
	}

	for i := 0; i < numSensors; i++ {
		data := <-c
		for sensor := range data {
			payload.Sensors[sensor] = make(map[string]interface{})
			for device := range data[sensor] {
				payload.Sensors[sensor][device] = data[sensor][device]
			}
		}
	}

	if len(payload.Sensors) == 0 {
		log.Error(errors.New("collected no data"))
		return
	}
	bPayload, err := json.MarshalIndent(payload, "", " ")
	log.Debug(string(bPayload))
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

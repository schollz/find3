package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"

	log "github.com/cihub/seelog"
	"github.com/de0gee/de0gee-data/src/models"
)

type Packet struct {
	Mac       string    `json:"mac"`
	RSSI      float64   `json:"rssi"`
	Timestamp time.Time `json:"timestamp"`
}

func ReverseScan(scanTime time.Duration) (sensors models.SensorData, err error) {
	tempFileName := "tshark-" + RandomString(10)
	tempFile := path.Join("/tmp", tempFileName)
	log.Debugf("saving tshark data to %s", tempFile)

	command := fmt.Sprintf("tshark -I -i %s -a duration:%d -w %s", wifiInterface, int(scanTime.Seconds()), tempFile)
	log.Debug(command)
	RunCommand(scanTime+scanTime+scanTime+scanTime+scanTime, command)

	out, _ := RunCommand(scanTime+scanTime+scanTime+scanTime, "/usr/bin/tshark -r "+tempFile+" -T fields -e frame.time_epoch -e wlan.sa -e wlan.bssid -e radiotap.dbm_antsignal")
	lines := strings.Split(out, "\n")
	packets := make([]Packet, len(lines))
	i := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 4 {
			continue
		}

		// determine time
		timeSeconds, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			log.Error(err)
			continue
		}
		nanoSeconds := int64(timeSeconds * 1e9)

		// determine rssi
		rssi, err := strconv.ParseFloat(strings.Split(fields[3], ",")[0], 64)
		if err != nil {
			log.Error(err)
			continue
		}
		packet := Packet{}
		packet.Timestamp = time.Unix(0, nanoSeconds)
		packet.Mac = fields[1]
		packet.RSSI = rssi
		packets[i] = packet
		i++
	}
	packets = packets[:i]

	// merge packets
	strengths := make(map[string][]float64)
	for _, packet := range packets {
		if _, ok := strengths[packet.Mac]; !ok {
			strengths[packet.Mac] = []float64{}
		}
		strengths[packet.Mac] = append(strengths[packet.Mac], packet.RSSI)
	}
	mergedPackets := make(map[string]struct{})
	newPackets := make([]Packet, len(packets))
	i = 0
	for _, packet := range packets {
		if _, ok := mergedPackets[packet.Mac]; ok {
			continue
		}
		packet.RSSI = Average(strengths[packet.Mac])
		newPackets[i] = packet
		i++
		mergedPackets[packet.Mac] = struct{}{}
	}
	packets = newPackets[:i]
	log.Infof("collected %d packets", len(packets))
	if len(packets) == 0 {
		err = errors.New("no packets found")
		return
	}
	sensors = models.SensorData{}
	sensors.Family = family
	sensors.Device = device
	sensors.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	sensors.Sensors = make(map[string]map[string]interface{})
	sensors.Sensors["wifi"] = make(map[string]interface{})
	for _, packet := range packets {
		sensors.Sensors["wifi"][packet.Mac] = packet.RSSI
	}

	bSensors, _ := json.MarshalIndent(sensors, "", " ")
	log.Debug(string(bSensors))
	return
}

func PromiscuousMode(on bool) {
	sequence := []string{"ifconfig XX down", "iwconfig XX mode YY", "ifconfig XX up"}
	for _, command := range sequence {
		commandString := strings.Replace(command, "XX", wifiInterface, 1)
		if on {
			commandString = strings.Replace(commandString, "YY", "monitor", 1)
		} else {
			commandString = strings.Replace(commandString, "YY", "managed", 1)
		}
		s, t := RunCommand(30*time.Second, commandString)
		time.Sleep(3 * time.Second)
		if len(s) > 0 {
			log.Debugf("out: %s", s)
		}
		if len(t) > 0 {
			log.Debugf("err: %s", t)
		}
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/de0gee/de0gee-data/src/database"
)

func main() {
	payload := database.SensorData{}
	payload.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	payload.Device = "dell"
	payload.Family = "test"
	payload.Location = "couch"
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
	bPayload, err := json.MarshalIndent(payload, "", " ")
	fmt.Println(string(bPayload), err)
}

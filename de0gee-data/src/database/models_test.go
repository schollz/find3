package database

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModels(t *testing.T) {
	j := `{
		"t":1514034330040,
		"f":"familyname",
		"d":"username",
		"l":"living room",
		"s":{
			 "wifi":{
					"aa:bb:cc:dd:ee":-20,
					"ff:gg:hh:ii:jj":-80
			 },
			 "bluetooth":{
					"aa:00:cc:11:ee":-42,
					"ff:22:hh:33:jj":-50        
			 },
			 "temperature":{
					"sensor1":12,
					"sensor2":20       
			 },
			 "accelerometer":{
					"x":-1.11,
					"y":2.111,
					"z":1.23   
			 }      
		}
 }`
	var p SensorData
	err := json.Unmarshal([]byte(j), &p)
	assert.Nil(t, err)
	assert.Equal(t, -20.0, p.Sensors["wifi"]["aa:bb:cc:dd:ee"])
}

func TestBackwards(t *testing.T) {
	jsonTest := `{"username": "zack", "group": "Find", "wifi-fingerprint": [{"rssi": -45, "mac": "80:37:73:ba:f7:d8"}], "location": "zakhome floor 2 office", "timestamp": 1439596533831, "password": "frusciante_0128"}`
	var f FINDFingerprint
	err := json.Unmarshal([]byte(jsonTest), &f)
	assert.Nil(t, err)
	d := f.Convert()

	j := `{
		"t":1439596533831,
		"f":"Find",
		"d":"zack",
		"l":"zakhome floor 2 office",
		"s":{
			 "wifi":{
					"80:37:73:ba:f7:d8":-45
			 }  
		}
 }`
	var p SensorData
	json.Unmarshal([]byte(j), &p)
	assert.Equal(t, p, d)
	fmt.Println(d)
}

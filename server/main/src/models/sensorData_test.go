package models

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
	assert.Nil(t, p.Validate())
}

func TestModels2(t *testing.T) {
	j := `{
  "d": "device1",
  "f": "daimler",
  "t": 1520424248897,
  "l": "LOCATION",
  "s": {
    "bluetooth": {
      "20:25:64:b7:91:42": -72,
      "20:25:64:b8:06:38": -81
    },
    "wifi": {
      "20:25:64:b7:91:40": -73,
      "70:4d:7b:11:3a:c8": -81,
      "88:d7:f6:a7:2a:4c": -39,
      "8c:0f:6f:e7:2b:78": -42,
      "8c:0f:6f:e7:2b:80": -43,
      "92:0f:6f:e7:2b:80": -43,
      "96:0f:6f:e7:2b:78": -39,
      "9e:0f:6f:e7:2b:80": -43,
      "ac:9e:17:7f:38:a4": -55,
      "dc:fe:07:79:aa:c0": -90,
      "dc:fe:07:79:aa:c3": -89
    }
  },
  "gps": {
    "lat": 12.1,
    "lon": 10.1,
    "alt": 54
  }
}
`
	var p SensorData
	err := json.Unmarshal([]byte(j), &p)
	assert.Nil(t, err)
	assert.Equal(t, -73.0, p.Sensors["wifi"]["20:25:64:b7:91:40"])
	assert.Nil(t, p.Validate())
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

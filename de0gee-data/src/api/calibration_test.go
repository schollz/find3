package api

import (
	"encoding/json"
	"testing"

	"github.com/de0gee/de0gee-data/src/database"
	"github.com/de0gee/de0gee-data/src/models"
	"github.com/stretchr/testify/assert"
)

// for testing purposes
var j = `{
	"t":1514034330040,
	"f":"familyname",
	"d":"devicename",
	"l":"bathroom",
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

// for testing purposes
var j2 = `{
	"t":1514034335555,
	"f":"familyname",
	"d":"devicename",
	"l":"kitchen",
	"s":{
		 "wifi":{
				"22:bb:cc:dd:ee":-19,
				"ff:gg:hh:ii:jj":-77
		 },
		 "bluetooth":{
				"aa:00:cc:11:ee":-40,
				"ff:22:hh:33:jj":-45        
		 },
		 "temperature":{
				"sensor1":10,
				"sensor2":32       
		 },
		 "accelerometer":{
				"x":-2.11,
				"y":4.111,
				"z":0.23   
		 }      
	}
}`

func BenchmarkDumpToCSV(b *testing.B) {
	var s models.SensorData
	db, _ := database.Open("testing")
	defer db.Close()
	json.Unmarshal([]byte(j), &s)
	db.AddSensor(s)
	json.Unmarshal([]byte(j2), &s)
	db.AddSensor(s)
	ss, _ := db.GetAllForClassification()

	db.Debug(false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dumpSensorsToCSV(ss, "test.csv")
	}

}

func TestDumpSensorsToCSV(t *testing.T) {
	var s models.SensorData
	db, _ := database.Open("testing")
	defer db.Close()
	json.Unmarshal([]byte(j), &s)
	db.AddSensor(s)
	json.Unmarshal([]byte(j2), &s)
	db.AddSensor(s)
	ss, _ := db.GetAllForClassification()

	db.Debug(false)
	err := dumpSensorsToCSV(ss, "test.csv")
	assert.Nil(t, err)
}

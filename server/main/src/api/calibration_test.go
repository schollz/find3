package api

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/schollz/find3/server/main/src/database"
	"github.com/schollz/find3/server/main/src/models"
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

func TestRisingEfficacy(t *testing.T) {
	DataFolder, _ = filepath.Abs("../../data")
	database.DataFolder = DataFolder

	db, err := database.Open("pike5")
	assert.Nil(t, err)
	datas, err := db.GetAllForClassification()
	assert.Nil(t, err)
	db.Close()
	datas = datas[:2000]
	fmt.Println(len(datas))

	datasLearn, datasTest, err := splitDataForLearning(datas, true)
	assert.Nil(t, err)
	fmt.Println(len(datasLearn))
	fmt.Println(len(datasTest))

	err = learnFromData("pike5", datasLearn)
	assert.Nil(t, err)

	algorithmEfficacy, err := findBestAlgorithm(datasTest)
	assert.Nil(t, err)
	// bA, _ := json.MarshalIndent(algorithmEfficacy, "", " ")
	// fmt.Println(string(bA))
	bestInformedness := make(map[string][]float64)
	for alg := range algorithmEfficacy {
		for loc := range algorithmEfficacy[alg] {
			if _, ok := bestInformedness[loc]; !ok {
				bestInformedness[loc] = []float64{}
			}
			bestInformedness[loc] = append(bestInformedness[loc], algorithmEfficacy[alg][loc].Informedness)
		}
	}
	for loc := range bestInformedness {
		fmt.Println(loc, Max(bestInformedness[loc]))
	}
}

// Max returns the maximum value in the input slice. If the slice is empty, Max will panic.
func Max(s []float64) float64 {
	return s[MaxIdx(s)]
}

// MaxIdx returns the index of the maximum value in the input slice. If several
// entries have the maximum value, the first such index is returned. If the slice
// is empty, MaxIdx will panic.
func MaxIdx(s []float64) int {
	if len(s) == 0 {
		panic("floats: zero slice length")
	}
	max := s[0]
	var ind int
	for i, v := range s {
		if v > max {
			max = v
			ind = i
		}
	}
	return ind
}

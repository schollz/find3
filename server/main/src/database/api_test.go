package database

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/schollz/find2/server/main/src/models"
	"github.com/stretchr/testify/assert"
)

func init() {
	Debug(false)
}
func TestAddSensor(t *testing.T) {
	var s1 models.SensorData
	var s2 models.SensorData
	err := json.Unmarshal([]byte(j), &s1)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(j2), &s2)
	if err != nil {
		panic(err)
	}
	db, _ := Open("testing")
	defer db.Close()
	err = db.AddSensor(s1)
	assert.Nil(t, err)
	err = db.AddSensor(s2)
	assert.Nil(t, err)

	s1test, err := db.GetSensorFromTime(s1.Timestamp)
	assert.Nil(t, err)
	assert.Equal(t, s1test, s1)

	sLatest, err := db.GetLatest(s2.Device)
	assert.Nil(t, err)
	assert.Equal(t, s2, sLatest)

	sPrepared, err := db.GetAllFromPreparedQuery("select * from sensors where timestamp = ?", s1.Timestamp)
	assert.Nil(t, err)
	assert.Equal(t, s1, sPrepared[0])
}

func TestGetAllForClassification(t *testing.T) {
	os.Remove("test.csv")

	var err error
	var s models.SensorData
	db, _ := Open("testing")
	defer db.Close()
	json.Unmarshal([]byte(j), &s)
	err = db.AddSensor(s)
	assert.Nil(t, err)
	json.Unmarshal([]byte(j2), &s)
	err = db.AddSensor(s)
	assert.Nil(t, err)

	ss, err := db.GetAllForClassification()
	assert.Equal(t, 2, len(ss))
	assert.Nil(t, err)

}

func BenchmarkAddSensor(b *testing.B) {
	var s models.SensorData
	json.Unmarshal([]byte(j), &s)
	db, _ := Open("testing")
	defer db.Close()
	db.Debug(false)

	for i := 0; i < b.N; i++ {
		s.Timestamp = int64(i)
		err := db.AddSensor(s)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGetSensor(b *testing.B) {
	var s models.SensorData
	err := json.Unmarshal([]byte(j), &s)
	if err != nil {
		panic(err)
	}
	db, _ := Open("testing")
	defer db.Close()
	db.Debug(false)
	err = db.AddSensor(s)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := db.GetSensorFromTime(s.Timestamp)
		if err != nil {
			panic(err)
		}
	}
}
func BenchmarkKeystoreSet(b *testing.B) {
	db, _ := Open("testing")
	defer db.Close()
	db.Debug(false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := db.Set("human:"+strconv.Itoa(i), Human{"Dante", 5.4})
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkKeystoreOpenAndSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		db, _ := Open("testing")
		db.Debug(false)
		err := db.Set("human:"+strconv.Itoa(i), Human{"Dante", 5.4})
		if err != nil {
			panic(err)
		}
		db.Close()
	}
}

func BenchmarkKeystoreGet(b *testing.B) {
	db, _ := Open("testing")
	defer db.Close()
	db.Debug(false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var h2 Human
		db.Get("human:"+strconv.Itoa(i), &h2)
	}
}

func BenchmarkGetLatest(b *testing.B) {
	var s1 models.SensorData
	json.Unmarshal([]byte(j), &s1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		db, _ := Open("testing")
		db.Debug(false)
		db.GetLatest(s1.Device)
		db.Close()
	}
}

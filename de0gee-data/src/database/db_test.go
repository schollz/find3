package database

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/de0gee/datastore/src/sensor"
	"github.com/stretchr/testify/assert"
)

// Human is for testing purposes
type Human struct {
	Name   string
	Height float64
}

// for testing purposes
var j = `{
	"t":1514034330040,
	"f":"familyname",
	"u":"username",
	"s":{
		"location": {
			"bathroom":1
		},
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

func TestKeystore(t *testing.T) {
	db, err := Open("testing")
	assert.Nil(t, err)

	err = db.Set("hello", "world")
	assert.Nil(t, err)
	var s string
	err = db.Get("hello", &s)
	assert.Nil(t, err)
	assert.Equal(t, s, "world")

	h := Human{"Dante", 5.4}
	err = db.Set("human1", h)
	assert.Nil(t, err)
	var h2 Human
	err = db.Get("human1", &h2)
	assert.Nil(t, err)
	assert.Equal(t, h, h2)

	// check that key doesn't exist
	err = db.Get("human2", &h2)
	assert.NotNil(t, err)

	// check the table columns
	var columns []string
	columns, err = db.Columns()
	assert.Nil(t, err)
	assert.Equal(t, []string{"timestamp", "family", "user", "location"}, columns)

	err = db.Close()
	assert.Nil(t, err)
}

func TestConcurrency(t *testing.T) {
	errors := make(chan error)
	for i := 0; i < 3; i++ {
		go func(n int) {
			db, _ := Open("testing")
			defer db.Close()
			time.Sleep(time.Millisecond * 100)
			errors <- db.Set("concurrentHuman:"+strconv.Itoa(n), Human{"Dante", 5.4})
		}(i)
	}
	for i := 0; i < 3; i++ {
		assert.Nil(t, <-errors)
	}
}

func TestAddSensor(t *testing.T) {
	var s sensor.Data
	err := json.Unmarshal([]byte(j), &s)
	if err != nil {
		panic(err)
	}
	db, _ := Open("testing")
	defer db.Close()
	err = db.AddSensor(s)
	assert.Nil(t, err)

	s2, err := db.GetSensorFromTime(s.Timestamp)
	assert.Nil(t, err)
	fmt.Println(s2)
}

func BenchmarkAddSensor(b *testing.B) {
	var s sensor.Data
	json.Unmarshal([]byte(j), &s)
	db, _ := Open("testing")
	defer db.Close()
	for i := 0; i < b.N; i++ {
		s.Timestamp = i
		err := db.AddSensor(s)
		if err != nil {
			panic(err)
		}
	}
}
func BenchmarkKeystoreSet(b *testing.B) {
	db, _ := Open("testing")
	defer db.Close()
	Debug(false)
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
		Debug(false)
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
	Debug(false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var h2 Human
		db.Get("human:"+strconv.Itoa(i), &h2)
	}
}

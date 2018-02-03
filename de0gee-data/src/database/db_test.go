package database

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Remove("dGVzdGluZw==.sqlite3.db")
}

// Human is for testing purposes
type Human struct {
	Name   string
	Height float64
}

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
	assert.Equal(t, []string{"timestamp", "family", "device", "location"}, columns[0:4])

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

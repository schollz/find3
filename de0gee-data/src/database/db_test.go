package database

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Human is for testing purposes
type Human struct {
	Name   string
	Height float64
}

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

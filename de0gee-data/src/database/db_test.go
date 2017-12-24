package database

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	db, err := Open("testing")
	assert.Nil(t, err)
	fmt.Println(db)
	err = db.Close()
	assert.Nil(t, err)
}

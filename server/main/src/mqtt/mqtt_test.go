package mqtt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMQTT(t *testing.T) {
	Debug = true
	err := Setup()
	assert.Nil(t, err)
	err = Publish("labs", "fido", "hello, world")
	assert.Nil(t, err)
	password, err := AddFamily("labs")
	assert.Nil(t, err)
	fmt.Println(password)
}

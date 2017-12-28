package mqtt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMQTT(t *testing.T) {
	Debug(true)
	err := Setup()
	assert.Nil(t, err)

}

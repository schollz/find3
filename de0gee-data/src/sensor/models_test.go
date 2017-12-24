package sensor

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModels(t *testing.T) {
	j := `{
		"t":1514034330040,
		"f":"familyname",
		"u":"username",
		"a":"asdlkjf.alsdkfj.aiwejciwe234",
		"s":{
			 "location":{
				 "living room":1
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
	var p Data
	err := json.Unmarshal([]byte(j), &p)
	assert.Nil(t, err)
	assert.Equal(t, -20.0, p.Sensors["wifi"]["aa:bb:cc:dd:ee"])
}

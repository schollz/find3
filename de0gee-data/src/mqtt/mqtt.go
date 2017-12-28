package mqtt

import (
	"github.com/de0gee/de0gee-data/src/utils"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

var (
	// Server is the address of the broker to use for MQTT
	Server        = "localhost:1883"
	Existing      = false
	AdminUser     = "zack"
	AdminPassword = "1234"
)

var (
	adminClient MQTT.Client
)

func Setup() (err error) {
	logger := log.WithFields(log.Fields{
		"name": "Setup",
	})
	logger.Info("setting up")

	server := "tcp://" + Server
	opts := MQTT.NewClientOptions()
	if Existing {
		logger.Debug("using existing setup")
		opts.AddBroker(server).SetClientID(utils.RandomString(5)).SetCleanSession(true)
	} else {
		logger.Debug("using current setup")
		// updateMosquittoConfig()
		opts.AddBroker(server).SetClientID(utils.RandomString(5)).SetUsername(AdminUser).SetPassword(AdminPassword).SetCleanSession(true)
	}

	adminClient = MQTT.NewClient(opts)

	if token := adminClient.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
	}
	logger.Debug("finished setup")
	return
}

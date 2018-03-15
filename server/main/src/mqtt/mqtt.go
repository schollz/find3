package mqtt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	"github.com/schollz/find3/server/main/src/api"
	"github.com/schollz/find3/server/main/src/database"
	"github.com/schollz/find3/server/main/src/logging"
	"github.com/schollz/find3/server/main/src/models"
	"github.com/schollz/find3/server/main/src/utils"
)

var (
	// Server is the address of the broker to use for MQTT
	Server                   = "localhost:1883"
	Debug                    = false
	logger                   *logging.SeelogWrapper
	Existing                 = false
	IsSetup                  = false
	AdminUser                = "zack"
	AdminPassword            = "1234"
	MosquittoConfigDirectory = "mosquitto_config"
)

var (
	adminClient MQTT.Client
)

func Setup() (err error) {
	logger, _ = logging.New()
	if Debug {
		logger.SetLevel("debug")
	} else {
		logger.SetLevel("info")
	}

	logger.Log.Debug("setting up")

	server := "tcp://" + Server
	opts := MQTT.NewClientOptions()
	if Existing {
		logger.Log.Debug("using existing setup")
		opts.AddBroker(server).SetClientID(utils.RandomString(5)).SetCleanSession(true)
	} else {
		logger.Log.Debug("using current setup")
		err = updateMosquittoConfig()
		if err != nil {
			return
		}
		opts.AddBroker(server).SetClientID(utils.RandomString(5)).SetUsername(AdminUser).SetPassword(AdminPassword).SetCleanSession(true)
	}
	// subscribe
	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe("#", 1, messageReceived); token.Wait() && token.Error() != nil {
			err = token.Error()
			return
		}
	}

	adminClient = MQTT.NewClient(opts)

	if token := adminClient.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
	}
	logger.Log.Debug("finished setup")
	IsSetup = true
	return
}

func updateMosquittoConfig() (err error) {
	// open the database that stores the basic parameters
	logger.Log.Debug("opening mosquitto database")
	db, err := database.Open("mosquitto", false, true)
	if err != nil {
		return
	}
	defer db.Close()
	logger.Log.Debug("starting database")

	// check if the defaults exist, otherwise create them
	var errGet error
	var acl, passwd, conf string
	errGet = db.Get("acl", &acl)
	if errGet != nil {
		logger.Log.Debug("making acl")
		acl = fmt.Sprintf("user %s\ntopic readwrite #\n\n", AdminUser)
	}
	errGet = db.Get("passwd", &passwd)
	if errGet != nil {
		logger.Log.Debug("making passwd")
		passwd = fmt.Sprintf("%s:%s\n", AdminUser, AdminPassword)
	}
	errGet = db.Get("conf", &conf)
	if errGet != nil {
		logger.Log.Debug("making conf")
		conf = fmt.Sprintf("allow_anonymous false\n\nacl_file %s/acl\n\npassword_file %s/passwd\n\npid_file %s/pid", MosquittoConfigDirectory, MosquittoConfigDirectory, MosquittoConfigDirectory)
	}

	var passes map[string]string
	errGet = db.Get("passes", &passes)
	if errGet == nil {
		for user := range passes {
			acl = acl + fmt.Sprintf("user %s\ntopic readwrite %s/#\n\n", user, user)
			passwd = passwd + fmt.Sprintf("%s:%s\n", user, passes[user])
		}
	}

	os.MkdirAll(MosquittoConfigDirectory, 0755)
	err = ioutil.WriteFile(path.Join(MosquittoConfigDirectory, "acl"), []byte(acl), 0644)
	if err != nil {
		err = errors.Wrap(err, "could not open to write file")
		return
	}
	err = ioutil.WriteFile(path.Join(MosquittoConfigDirectory, "passwd"), []byte(passwd), 0644)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path.Join(MosquittoConfigDirectory, "mosquitto.conf"), []byte(conf), 0644)
	if err != nil {
		err = errors.Wrap(err, "could not write conf")
		return
	}

	// generate passwords
	cmd := "mosquitto_passwd"
	args := []string{"-U", path.Join(MosquittoConfigDirectory, "passwd")}
	err = exec.Command(cmd, args...).Run()
	if err != nil {
		err = errors.Wrap(err, "problem writing new passwd")
		return
	}

	// regenerate mosquitto
	bPID, errPID := ioutil.ReadFile(path.Join(MosquittoConfigDirectory, "pid"))
	if errPID != nil {
		logger.Log.Debug("could not get PID, running")
		// try running by itself
		cmd = "mosquitto"
		args = []string{"-c", fmt.Sprintf("%s/mosquitto.conf", MosquittoConfigDirectory), "-d"}
		if err = exec.Command(cmd, args...).Run(); err != nil {
			err = errors.Wrap(err, "problem runnign")
		}
		return
	}
	cmd = "kill"
	args = []string{"-HUP", string(bPID)}
	if err = exec.Command(cmd, args...).Run(); err != nil {
		err = errors.Wrap(err, "problem giving HUP")
		return
	}
	logger.Log.Debug("setup mosquitto and gave HUP signal")
	return
}

// add pushes a new family into the database
func add(family string) (password string, err error) {
	// open the database that stores the basic parameters
	db, err := database.Open("mosquitto", false, true)
	if err != nil {
		return
	}
	defer db.Close()

	var passes map[string]string
	errGet := db.Get("passes", &passes)
	if errGet != nil {
		passes = make(map[string]string)
	}
	password = utils.RandomString(5)
	passes[family] = password
	err = db.Set("passes", passes)
	return
}

func AddFamily(family string) (password string, err error) {
	password, err = add(family)
	if err != nil {
		return
	}
	err = updateMosquittoConfig()
	return
}

func Publish(family, device, message string) (err error) {
	if !IsSetup {
		return errors.New("mqtt not setup")
	}
	pubTopic := strings.Join([]string{family, "/location/", device}, "")

	if token := adminClient.Publish(pubTopic, 1, false, message); token.Wait() && token.Error() != nil {
		err = fmt.Errorf("Failed to send message")
	}
	return
}

func messageReceived(client MQTT.Client, msg MQTT.Message) {
	jsonFingerprint, route, err := mqttBuildFingerprint(msg.Topic(), msg.Payload())
	if err != nil {
		return
	}
	logger.Log.Debug("Got valid MQTT request for group " + jsonFingerprint.Group + ", user " + jsonFingerprint.Username)
	if route == "track" {
		jsonFingerprint.Location = ""
	}
	d := jsonFingerprint.Convert()
	err = api.SaveSensorData(d)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	_, err = sendOutData(d)
	if err != nil {
		logger.Log.Error(err)
		return
	}
}

func sendOutData(p models.SensorData) (analysis models.LocationAnalysis, err error) {
	analysis, _ = api.AnalyzeSensorData(p)
	type Payload struct {
		Sensors models.SensorData           `json:"sensors"`
		Guesses []models.LocationPrediction `json:"guesses"`
	}
	payload := Payload{
		Sensors: p,
		Guesses: analysis.Guesses,
	}
	bTarget, err := json.Marshal(payload)
	if err != nil {
		return
	}
	logger.Log.Debugf("[%s] sending data over mqtt (%s)", p.Family, p.Device)
	Publish(p.Family, p.Device, string(bTarget))
	return
}

// backwards compatible with FIND
func mqttBuildFingerprint(topic string, message []byte) (jsonFingerprint models.FINDFingerprint, route string, err error) {
	err = nil
	route = "track"
	topics := strings.Split(strings.ToLower(topic), "/")
	jsonFingerprint.Location = ""
	if len(topics) < 3 || (topics[1] != "track" && topics[1] != "learn") {
		err = fmt.Errorf("Must define track or learn")
		return
	}
	route = topics[1]
	if route == "track" && len(topics) != 3 {
		err = fmt.Errorf("Track needs a user name")
		return
	}
	if route == "learn" {
		if len(topics) != 4 {
			err = fmt.Errorf("Track needs a user name and location")
			return
		} else {
			jsonFingerprint.Location = topics[3]
		}
	}
	jsonFingerprint.Group = topics[0]
	jsonFingerprint.Username = topics[2]
	routers := []models.Router{}
	for i := 0; i < len(message); i += 14 {
		if (i + 14) > len(message) {
			break
		}
		mac := string(message[i:i+2]) + ":" + string(message[i+2:i+4]) + ":" + string(message[i+4:i+6]) + ":" + string(message[i+6:i+8]) + ":" + string(message[i+8:i+10]) + ":" + string(message[i+10:i+12])
		val, _ := strconv.Atoi(string(message[i+12 : i+14]))
		rssi := -1 * val
		routers = append(routers, models.Router{Mac: mac, Rssi: rssi})
	}
	jsonFingerprint.WifiFingerprint = routers
	return
}

package mqtt

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/de0gee/de0gee-data/src/database"
	"github.com/de0gee/de0gee-data/src/utils"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	// Server is the address of the broker to use for MQTT
	Server                   = "localhost:1883"
	Existing                 = false
	AdminUser                = "zack"
	AdminPassword            = "1234"
	MosquittoConfigDirectory = "mosquitto_config"
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
		err = updateMosquittoConfig()
		if err != nil {
			return
		}
		opts.AddBroker(server).SetClientID(utils.RandomString(5)).SetUsername(AdminUser).SetPassword(AdminPassword).SetCleanSession(true)
	}

	adminClient = MQTT.NewClient(opts)

	if token := adminClient.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
	}
	logger.Debug("finished setup")
	return
}

type Pass struct {
	User     string
	Password string
}

func updateMosquittoConfig() (err error) {
	logger := log.WithFields(log.Fields{
		"name": "updateMosquittoConfig",
	})
	// open the database that stores the basic parameters
	db, err := database.Open("mosquitto")
	if err != nil {
		return
	}
	defer db.Close()

	// check if the defaults exist, otherwise create them
	var errGet error
	var acl, passwd, conf string
	errGet = db.Get("acl", &acl)
	if errGet != nil {
		logger.Debug("making acl")
		acl = fmt.Sprintf("user %s\ntopic readwrite #\n\n", AdminUser)
	}
	errGet = db.Get("passwd", &passwd)
	if errGet != nil {
		logger.Debug("making passwd")
		passwd = fmt.Sprintf("%s:%s\n", AdminUser, AdminPassword)
	}
	errGet = db.Get("conf", &conf)
	if errGet != nil {
		logger.Debug("making conf")
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

	// give mosquitto HUP signal
	bPID, errPID := ioutil.ReadFile(path.Join(MosquittoConfigDirectory, "pid"))
	if errPID != nil {
		logger.Debug("could not get PID, running")
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
	logger.Debug("setup mosquitto and gave HUP signal")
	return
}

func add(family string) (password string, err error) {
	// open the database that stores the basic parameters
	db, err := database.Open("mosquitto")
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
	pubTopic := strings.Join([]string{family, "/location/", device}, "")

	if token := adminClient.Publish(pubTopic, 1, false, message); token.Wait() && token.Error() != nil {
		err = fmt.Errorf("Failed to send message")
	}
	return
}

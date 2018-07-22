package wigle

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/schollz/find3/server/main/src/logging"
)

var (
	DataFolder = "data"
	DebugMode  bool
	logger     *logging.SeelogWrapper
)

type WigleResponse struct {
	Success   bool          `json:"success"`
	Cdma      bool          `json:"cdma"`
	Gsm       bool          `json:"gsm"`
	Wifi      bool          `json:"wifi"`
	Addresses []interface{} `json:"addresses"`
	Results   []struct {
		Trilat       float64     `json:"trilat"`
		Trilong      float64     `json:"trilong"`
		Ssid         string      `json:"ssid"`
		Qos          float64     `json:"qos"`
		Transid      string      `json:"transid"`
		Firsttime    time.Time   `json:"firsttime"`
		Lasttime     time.Time   `json:"lasttime"`
		Lastupdt     time.Time   `json:"lastupdt"`
		Housenumber  interface{} `json:"housenumber"`
		Road         interface{} `json:"road"`
		City         interface{} `json:"city"`
		Region       interface{} `json:"region"`
		Country      interface{} `json:"country"`
		Netid        string      `json:"netid"`
		Name         interface{} `json:"name"`
		Type         string      `json:"type"`
		Comment      interface{} `json:"comment"`
		Wep          string      `json:"wep"`
		Channel      float64     `json:"channel"`
		Bcninterval  float64     `json:"bcninterval"`
		Freenet      string      `json:"freenet"`
		Dhcp         string      `json:"dhcp"`
		Paynet       string      `json:"paynet"`
		Userfound    interface{} `json:"userfound"`
		LocationData []struct {
			Alt             float64     `json:"alt"`
			Accuracy        float64     `json:"accuracy"`
			Lastupdt        time.Time   `json:"lastupdt"`
			Latitude        float64     `json:"latitude"`
			Longitude       float64     `json:"longitude"`
			Month           string      `json:"month"`
			Ssid            string      `json:"ssid"`
			Time            time.Time   `json:"time"`
			Signal          float64     `json:"signal"`
			Name            interface{} `json:"name"`
			NetID           string      `json:"netId"`
			Noise           float64     `json:"noise"`
			Snr             float64     `json:"snr"`
			Wep             string      `json:"wep"`
			EncryptionValue string      `json:"encryptionValue"`
		} `json:"locationData"`
		Encryption string `json:"encryption"`
	} `json:"results"`
}

func Debug(debugMode bool) {
	DebugMode = debugMode
	if debugMode {
		logger.SetLevel("debug")
	} else {
		logger.SetLevel("info")
	}
}

func init() {
	err := start()
	if err != nil {
		panic(err)
	}
	logger, err = logging.New()
	if err != nil {
		panic(err)
	}
	Debug(false)
}

func start() (err error) {
	db, err := bolt.Open(path.Join(DataFolder, "wigle.boltdb"), 0600, nil)
	if err != nil {
		return
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("jsons"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return
	}
	err = db.Close()
	return
}

// GetLocation attempts to find the lat/lon for a given set of macs
func GetLocation(macs []string) (latitude, longitude float64) {
	if os.Getenv("WIGLE_NAME") == "" && os.Getenv("WIGLE_TOKEN") == "" {
		return
	}
	logger.Log.Debugf("getting location for %d macs", len(macs))
	accounted := 0.0
	for _, mac := range macs {
		wr, _ := GetWifi(mac)
		if len(wr.Results) > 0 {
			if len(wr.Results[0].LocationData) > 0 {
				latitude += wr.Results[0].LocationData[0].Latitude
				longitude += wr.Results[0].LocationData[0].Longitude
				accounted++
			}
		}
	}
	if accounted == 0 {
		return
	}
	latitude = latitude / accounted
	longitude = longitude / accounted
	return
}

// GetWiFi uses the Wigle API to get the information for a mac
func GetWifi(mac string) (wr WigleResponse, err error) {
	mac = strings.TrimSpace(strings.ToLower(mac))
	db, err := bolt.Open(path.Join(DataFolder, "wigle.boltdb"), 0600, nil)
	if err != nil {
		return
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("jsons"))
		v := b.Get([]byte(mac))
		if v != nil {
			return json.Unmarshal(v, &wr)
		}
		return errors.New("not found")
	})
	if err == nil {
		return
	}

	logger.Log.Debugf("requesting %s from wigle.net", mac)
	req, err := http.NewRequest("GET", "https://api.wigle.net/api/v2/network/detail?netid="+mac, nil)
	if err != nil {
		return
	}
	req.SetBasicAuth(os.Getenv("WIGLE_NAME"), os.Getenv("WIGLE_TOKEN"))
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&wr)
	if err == nil {
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("jsons"))
			bj, err := json.Marshal(wr)
			if err != nil {
				return err
			}
			return b.Put([]byte(mac), bj)
		})
	}
	return
}

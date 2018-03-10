package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/schollz/find3/server/main/src/api"
	"github.com/schollz/find3/server/main/src/database"
	"github.com/schollz/find3/server/main/src/models"
	"github.com/schollz/find3/server/main/src/mqtt"
	"github.com/schollz/find3/server/main/src/utils"
)

// Port defines the public port
var Port = "8003"
var UseSSL = false
var UseMQTT = false
var MinimumPassive = -1

// Run will start the server listening on the specified port
func Run() (err error) {
	defer logger.Log.Flush()

	if UseMQTT {
		// setup MQTT
		err = mqtt.Setup()
		if err != nil {
			logger.Log.Warn(err)
		}
		logger.Log.Debug("setup mqtt")
	}

	logger.Log.Debug("current families: ", database.GetFamilies())

	// setup gin server
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// Standardize logs
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.Use(middleWareHandler(), gin.Recovery())
	r.HEAD("/", func(c *gin.Context) { // handler for the uptime robot
		c.String(http.StatusOK, "OK")
	})
	r.GET("/", func(c *gin.Context) { // handler for the uptime robot
		c.String(http.StatusOK, "see www.internalpositioning.com/doc/api.md")
	})
	r.GET("/view/location/:family/:device", func(c *gin.Context) {
		family := c.Param("family")
		device := c.Param("device")
		c.HTML(http.StatusOK, "view_location.tmpl", gin.H{
			"Family": family,
			"Device": device,
			"SSL":    UseSSL,
		})
	})
	r.GET("/api/v1/devices/*family", handlerApiV1Devices)
	r.GET("/api/v1/location/:family/*device", handlerApiV1Location)
	r.GET("/api/v1/locations/:family", handlerApiV1Locations)
	r.GET("/api/v1/by_location/:family", handlerApiV1ByLocation)
	r.GET("/api/v1/calibrate/*family", handlerApiV1Calibrate)
	r.POST("/api/v1/settings/passive", handlerReverseSettings)
	r.GET("/ping", ping)
	r.GET("/test", handleTest)
	r.GET("/ws", wshandler) // handler for the web sockets (see websockets.go)
	if UseMQTT {
		r.POST("/mqtt", handlerMQTT) // handler for setting MQTT
	}
	r.POST("/data", handlerData)       // typical data handler
	r.POST("/gps", handlerGPS)         // typical GPS handler
	r.POST("/passive", handlerReverse) // typical data handler
	r.POST("/learn", handlerFIND)      // backwards-compatible with FIND for learning
	r.POST("/track", handlerFIND)      // backwards-compatible with FIND for tracking
	logger.Log.Infof("Running on 0.0.0.0:%s", Port)

	err = r.Run(":" + Port) // listen and serve on 0.0.0.0:8080
	return
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func handleTest(c *gin.Context) {
	d, _ := database.Open("testdb", true)
	err := d.Dump()
	if err != nil {
		fmt.Println(err)
	}
	d.Close()
	c.String(http.StatusOK, "ok")
}

func handlerApiV1Devices(c *gin.Context) {
	err := func(c *gin.Context) (err error) {
		family := strings.TrimSpace(c.Param("family")[1:])
		d, err := database.Open(family, true)
		if err != nil {
			return
		}
		defer d.Close()
		s, err := d.GetDevices()
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "got devices", "success": true, "devices": s})
		return
	}(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": false})
	}
}

func handlerApiV1Locations(c *gin.Context) {
	type Location struct {
		Device   string                  `json:"device"`
		Sensors  models.SensorData       `json:"sensors"`
		Analysis models.LocationAnalysis `json:"analysis"`
	}

	locations, err := func(c *gin.Context) (locations []Location, err error) {
		family := strings.TrimSpace(c.Param("family"))

		d, err := database.Open(family, true)
		if err != nil {
			return
		}
		devices, err := d.GetDevices()
		d.Close()
		if err != nil {
			return
		}
		locations = make([]Location, len(devices))
		for i, device := range devices {
			d, err = database.Open(family, true)
			if err != nil {
				return
			}
			locations[i] = Location{Device: device}
			locations[i].Sensors, err = d.GetLatest(device)
			d.Close()
			if err != nil {
				return
			}
			locations[i].Analysis, err = api.AnalyzeSensorData(locations[i].Sensors)
			if err != nil {
				return
			}
		}

		return
	}(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": err == nil})
	} else {

		c.JSON(http.StatusOK, gin.H{"message": "got locations", "success": err == nil, "locations": locations})
	}
}

func handlerApiV1ByLocation(c *gin.Context) {
	type Location struct {
		Device      string    `json:"device"`
		Timestamp   time.Time `json:"timestamp"`
		Probability float64   `json:"probability"`
		Randomized  bool      `json:"random_mac"`
	}

	locations, err := func(c *gin.Context) (locations map[string][]Location, err error) {
		locations = make(map[string][]Location)
		family := strings.TrimSpace(c.Param("family"))
		minutesAgo := strings.TrimSpace(c.DefaultQuery("history", "120"))
		allowRandomization := c.DefaultQuery("random", "1") == "1"

		minutesAgoInt, err := strconv.Atoi(minutesAgo)
		if err != nil {
			return
		}
		millisecondsAgo := int64(minutesAgoInt * 60 * 1000)

		d, err := database.Open(family, true)
		if err != nil {
			return
		}
		sensors, err := d.GetSensorFromGreaterTime(millisecondsAgo)

		preAnalyzed := make(map[int64][]models.LocationPrediction)
		for _, sensor := range sensors {
			a, errGet := d.GetPrediction(sensor.Timestamp)
			if errGet != nil {
				continue
			}
			preAnalyzed[sensor.Timestamp] = a
		}
		d.Close()
		if err != nil {
			return
		}
		logger.Log.Debugf("[%s] got %d sensors", family, len(sensors))
		for _, s := range sensors {
			isRandomized := utils.IsMacRandomized(s.Device)
			if allowRandomization == false && isRandomized {
				continue
			}
			var a []models.LocationPrediction
			if _, ok := preAnalyzed[s.Timestamp]; ok {
				logger.Log.Debugf("[%s] using pre-analyzed for %d", family, s.Timestamp)
				a = preAnalyzed[s.Timestamp]
			} else {
				var aidata models.LocationAnalysis
				aidata, err = api.AnalyzeSensorData(s)
				if err != nil {
					return
				}
				a = aidata.Guesses
			}
			if _, ok := locations[a[0].Location]; !ok {
				locations[a[0].Location] = []Location{}
			}
			locations[a[0].Location] = append(locations[a[0].Location], Location{
				Device:      s.Device,
				Timestamp:   time.Unix(0, s.Timestamp*1000000).UTC(),
				Probability: a[0].Probability,
				Randomized:  isRandomized,
			})
		}
		return
	}(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": err == nil})
	} else {

		c.JSON(http.StatusOK, gin.H{"message": "got locations", "success": err == nil, "locations": locations})
	}
}

func handlerApiV1Location(c *gin.Context) {
	s, analysis, err := func(c *gin.Context) (s models.SensorData, analysis models.LocationAnalysis, err error) {
		family := strings.TrimSpace(c.Param("family"))
		device := strings.TrimSpace(c.Param("device")[1:])

		d, err := database.Open(family, true)
		if err != nil {
			return
		}
		s, err = d.GetLatest(device)
		d.Close()
		if err != nil {
			return
		}
		analysis, err = api.AnalyzeSensorData(s)
		if err != nil {
			return
		}
		return
	}(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": err == nil})
	} else {

		c.JSON(http.StatusOK, gin.H{"message": "got location", "success": err == nil, "sensors": s, "analysis": analysis})
	}
}

func handlerApiV1Calibrate(c *gin.Context) {
	family := strings.TrimSpace(c.Param("family")[1:])
	var err error
	if family == "" {
		err = errors.New("invalid family")
	} else {
		err = api.Calibrate(family, true)
	}
	message := "calibrated data"
	if err != nil {
		message = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{"message": message, "success": err == nil})
}

func handlerMQTT(c *gin.Context) {
	type Payload struct {
		Family string `json:"family" binding:"required"`
		//	OTP    string `json:"otp" binding:"required"`
	}
	success := false
	var message string
	var p Payload
	if errBind := c.ShouldBindJSON(&p); errBind == nil {
		// TODO: authenticate p.OTP
		passphrase, err := mqtt.AddFamily(p.Family)
		if err != nil {
			message = err.Error()
		} else {
			message = fmt.Sprintf("Added '%s' for mqtt. Your passphrase is '%s'", p.Family, passphrase)
		}
	} else {
		message = errBind.Error()
	}
	c.JSON(http.StatusOK, gin.H{"message": message, "success": success})
}

func sendOutLocation(family, device string) (s models.SensorData, analysis models.LocationAnalysis, err error) {
	d, err := database.Open(family, true)
	if err != nil {
		return
	}
	s, err = d.GetLatest(device)
	d.Close()
	if err != nil {
		return
	}
	analysis, err = sendOutData(s)
	return
}

func handlerData(c *gin.Context) {
	message, err := func(c *gin.Context) (message string, err error) {
		var d models.SensorData
		err = c.BindJSON(&d)
		if err != nil {
			return
		}

		err = d.Validate()
		if err != nil {
			return
		}

		// process data
		err = processSensorData(d)
		if err != nil {
			return
		}
		message = "inserted data"
		hasGPS, _ := api.HasGPS(d)
		if !hasGPS {
			message += " [need GPS]"
		}

		logger.Log.Debugf("[%s] /data %+v", d.Family, d)
		return
	}(c)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": message, "success": true})
	}
}

func handlerGPS(c *gin.Context) {
	message, err := func(c *gin.Context) (message string, err error) {
		var data models.FingerprintWithGPS
		err = c.BindJSON(&data)
		if err != nil {
			return
		}
		err = api.AddGPSData(data)
		message = "added data"
		return
	}(c)

	if err != nil {
		logger.Log.Warn(err)
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": message, "success": true})
	}
}

func handlerReverseSettings(c *gin.Context) {
	message, err := func(c *gin.Context) (message string, err error) {
		// bind sensor data
		type ReverseSettings struct {
			// Minimum number of passive
			MinimumPassive int `json:"minimum_passive"`
			// Timespan of window
			Window int64 `json:"window"`
			// Family is a group of devices
			Family string `json:"family" binding:"required"`
			// Device are unique within a family
			Device string `json:"device"`
			// Location is optional, used for designating learning
			Location string `json:"location"`
		}
		var d ReverseSettings
		err = c.BindJSON(&d)
		if err != nil {
			return
		}
		d.Family = strings.TrimSpace(strings.ToLower(d.Family))
		d.Device = strings.TrimSpace(strings.ToLower(d.Device))
		d.Location = strings.TrimSpace(strings.ToLower(d.Location))

		// open database
		db, err := database.Open(d.Family)
		if err != nil {
			return
		}
		defer db.Close()

		var rollingData models.ReverseRollingData
		err = db.Get("ReverseRollingData", &rollingData)
		if err != nil {
			rollingData = models.ReverseRollingData{
				Family:         d.Family,
				DeviceLocation: make(map[string]string),
				TimeBlock:      20 * time.Second,
			}
		}
		if rollingData.TimeBlock.Seconds() == 0 {
			rollingData.TimeBlock = 20 * time.Second
		}

		// set tracking information
		if d.Device != "" {
			if d.Location != "" {
				message = fmt.Sprintf("Set location to '%s' for %s for learning with device '%s'", d.Location, d.Family, d.Device)
				rollingData.DeviceLocation[d.Device] = d.Location
			} else {
				message = fmt.Sprintf("switched to tracking for %s", d.Family)
				delete(rollingData.DeviceLocation, d.Device)
			}
			message += ". "
		}
		message += fmt.Sprintf("Now learning on %d devices: %+v", len(rollingData.DeviceLocation), rollingData.DeviceLocation)

		// set time block information
		if d.Window > 0 {
			rollingData.TimeBlock = time.Duration(d.Window) * time.Second
		}
		message += fmt.Sprintf("with time block of %2.0f seconds", rollingData.TimeBlock.Seconds())

		if d.MinimumPassive != 0 {
			rollingData.MinimumPassive = d.MinimumPassive
			message += fmt.Sprintf(" and set minimum passive to %d", rollingData.MinimumPassive)
		}

		err = db.Set("ReverseRollingData", rollingData)
		logger.Log.Debugf("[%s] %s", d.Family, message)
		return
	}(c)

	if err != nil {
		logger.Log.Warn(err)
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": message, "success": true})
	}
}

func handlerReverse(c *gin.Context) {
	message, err := func(c *gin.Context) (message string, err error) {
		// bind sensor data
		var d models.SensorData
		err = c.BindJSON(&d)
		if err != nil {
			return
		}

		// validate sensor data
		err = d.Validate()
		if err != nil {
			return
		}

		// open database
		db, err := database.Open(d.Family)
		if err != nil {
			return
		}
		defer db.Close()

		var rollingData models.ReverseRollingData
		err = db.Get("ReverseRollingData", &rollingData)
		if err != nil {
			// defaults
			rollingData = models.ReverseRollingData{
				Family:         d.Family,
				DeviceLocation: make(map[string]string),
				TimeBlock:      20 * time.Second,
			}
		}
		if rollingData.TimeBlock.Seconds() == 0 {
			rollingData.TimeBlock = 20 * time.Second
		}

		if !rollingData.HasData {
			rollingData.Timestamp = time.Now().UTC()
			rollingData.Datas = []models.SensorData{}
			rollingData.HasData = true
		}
		if len(d.Sensors) == 0 {
			err = errors.New("no fingerprints")
			return
		}

		rollingData.Datas = append(rollingData.Datas, d)
		numFingerprints := 0
		for sensor := range d.Sensors {
			numFingerprints += len(d.Sensors[sensor])
		}
		err = db.Set("ReverseRollingData", rollingData)
		message = fmt.Sprintf("inserted %d fingerprints for %s", numFingerprints, d.Family)

		if err == nil {
			go parseRollingData(d.Family)
		}
		return
	}(c)

	if err != nil {
		logger.Log.Warn(err)
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": message, "success": true})
	}

}

func parseRollingData(family string) (err error) {
	db, err := database.Open(family)
	if err != nil {
		return
	}
	defer db.Close()

	var rollingData models.ReverseRollingData
	err = db.Get("ReverseRollingData", &rollingData)
	if err != nil {
		return
	}

	sensorMap := make(map[string]models.SensorData)
	if rollingData.HasData && time.Since(rollingData.Timestamp) > rollingData.TimeBlock {
		logger.Log.Debugf("[%s] New data arrived %s", family, time.Since(rollingData.Timestamp))
		// merge data
		for _, data := range rollingData.Datas {
			for sensor := range data.Sensors {
				for mac := range data.Sensors[sensor] {
					rssi := data.Sensors[sensor][mac]
					trackedDeviceName := sensor + "-" + mac
					if _, ok := sensorMap[trackedDeviceName]; !ok {
						location := ""
						// if there is a device+location in map, then it is currently doing learning
						if loc, hasMac := rollingData.DeviceLocation[trackedDeviceName]; hasMac {
							location = loc
						}
						sensorMap[trackedDeviceName] = models.SensorData{
							Family:    family,
							Device:    trackedDeviceName,
							Timestamp: time.Now().UTC().UnixNano() / int64(time.Millisecond),
							Sensors:   make(map[string]map[string]interface{}),
							Location:  location,
						}
						time.Sleep(10 * time.Millisecond)
						sensorMap[trackedDeviceName].Sensors[sensor] = make(map[string]interface{})
					}
					sensorMap[trackedDeviceName].Sensors[sensor][data.Device+"-"+sensor] = rssi
				}
			}
		}
		rollingData.HasData = false
	}
	db.Set("ReverseRollingData", rollingData)
	db.Close()
	for sensor := range sensorMap {
		logger.Log.Debugf("[%s] reverse sensor data: %+v", family, sensorMap[sensor])
		numPassivePoints := 0
		for sensorType := range sensorMap[sensor].Sensors {
			numPassivePoints += len(sensorMap[sensor].Sensors[sensorType])
		}
		if numPassivePoints < rollingData.MinimumPassive {
			logger.Log.Debugf("[%s] skipped saving reverse sensor data for %s, not enough points (< %d)", family, sensor, rollingData.MinimumPassive)
			continue
		}
		err := processSensorData(sensorMap[sensor])
		if err != nil {
			logger.Log.Warnf("[%s] problem saving: %s", family, err.Error())
		}
		logger.Log.Debugf("[%s] saved reverse sensor data for %s", family, sensor)
	}

	return
}

func handlerFIND(c *gin.Context) {
	var j models.FINDFingerprint
	var err error
	var message string
	err = c.BindJSON(&j)
	if err == nil {
		if c.Request.URL.Path == "/track" {
			j.Location = ""
		}
		d := j.Convert()
		err2 := processSensorData(d)
		if err2 == nil {
			message = "inserted data"
		} else {
			err = err2
		}
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": message, "success": true})
	}
}

func processSensorData(p models.SensorData) (err error) {
	err = api.SaveSensorData(p)
	if err != nil {
		return
	}

	go sendOutData(p)
	return
}

func sendOutData(p models.SensorData) (analysis models.LocationAnalysis, err error) {
	analysis, _ = api.AnalyzeSensorData(p)

	type Payload struct {
		Sensors  models.SensorData       `json:"sensors"`
		Analysis models.LocationAnalysis `json:"analysis"`
	}
	payload := Payload{
		Sensors:  p,
		Analysis: analysis,
	}
	bTarget, err := json.Marshal(payload)
	if err != nil {
		return
	}

	// logger.Log.Debugf("sending data over websockets (%s/%s):%s", p.Family, p.Device, bTarget)
	SendMessageOverWebsockets(p.Family, p.Device, bTarget)

	if UseMQTT {
		// logger.Log.Debugf("sending data over websockets (%s/%s):%s", p.Family, p.Device, bTarget)
		mqtt.Publish(p.Family, p.Device, string(bTarget))
	}
	return
}

func middleWareHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now().UTC()
		// Add base headers
		addCORS(c)
		// Run next function
		c.Next()
		// Log request
		logger.Log.Infof("%v %v %v %s", c.Request.RemoteAddr, c.Request.Method, c.Request.URL, time.Since(t))
	}
}

func addCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
}

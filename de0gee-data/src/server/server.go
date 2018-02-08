package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/de0gee/de0gee-data/src/api"
	"github.com/de0gee/de0gee-data/src/database"
	"github.com/de0gee/de0gee-data/src/models"
	"github.com/de0gee/de0gee-data/src/mqtt"
	"github.com/gin-gonic/gin"
)

// Port defines the public port
var Port = "8003"

// Run will start the server listening on the specified port
func Run() (err error) {
	// setup MQTT
	err = mqtt.Setup()
	if err != nil {
		logger.Log.Warn(err)
	}
	logger.Log.Debug("setup mqtt")
	logger.Log.Debug("current families: ", database.GetFamilies())

	// setup gin server
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// Standardize logs
	r.Use(middleWareHandler(), gin.Recovery())
	r.HEAD("/", func(c *gin.Context) { // handler for the uptime robot
		c.String(http.StatusOK, "OK")
	})
	r.GET("/ping", ping)
	r.GET("/ws", wshandler)                  // handler for the web sockets (see websockets.go)
	r.POST("/mqtt", handlerMQTT)             // handler for setting MQTT
	r.POST("/data", handlerData)             // typical data handler
	r.POST("/reverse", handlerReverse)       // typical data handler
	r.POST("/learn", handlerFIND)            // backwards-compatible with FIND for learning
	r.POST("/track", handlerFIND)            // backwards-compatible with FIND for tracking
	r.GET("/location", handlerLocation)      // get the latest location
	r.POST("/calibrate", handlerCalibration) // calibrate to get the latest location
	logger.Log.Infof("Running on 0.0.0.0:%s", Port)

	// start goroutines
	logger.Log.Debugf("starting background processes")
	go checkRolingData()

	err = r.Run(":" + Port) // listen and serve on 0.0.0.0:8080
	return
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
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

func handleLocation(c *gin.Context) (err error) {
	t := time.Now()
	type Payload struct {
		Family string `json:"family" binding:"required"`
		Device string `json:"device" binding:"required"`
	}
	var p Payload
	err = c.ShouldBindJSON(&p)
	if err != nil {
		return
	}
	d, err := database.Open(p.Family, true)
	if err != nil {
		return
	}
	s, err := d.GetLatest(p.Device)
	d.Close()
	if err != nil {
		return
	}
	analysis, err := sendOutData(s)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "got latest in " + time.Since(t).String(), "success": true, "analysis": analysis, "sensors": s})
	return
}

func handleCalibration(c *gin.Context) (err error) {
	type Payload struct {
		Family string `json:"family" binding:"required"`
	}
	var p Payload
	err = c.BindJSON(&p)
	if err != nil {
		return
	}
	err = api.Calibrate(p.Family)
	return
}

func handlerCalibration(c *gin.Context) {
	t := time.Now()
	err := handleCalibration(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "calibrated data in " + time.Since(t).String(), "success": true})
	}
}

func handlerLocation(c *gin.Context) {
	err := handleLocation(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": false})
	}
}

func handlerData(c *gin.Context) {
	var err error
	var message string
	var d models.SensorData
	err = c.BindJSON(&d)
	if err == nil {
		err2 := processFingerprint(d)
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

func handlerReverse(c *gin.Context) {
	err := handleReverse(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": false})
	}
}
func handleReverse(c *gin.Context) (err error) {
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
		rollingData = models.ReverseRollingData{
			Family:         d.Family,
			DeviceLocation: make(map[string]string),
		}
	}
	var message string
	success := true
	if d.Timestamp == 1 {
		if d.Location != "" {
			message = fmt.Sprintf("set location to '%s' for %s for learning with device '%s'", d.Location, d.Family, d.Device)
			rollingData.DeviceLocation[d.Device] = d.Location
		} else {
			message = fmt.Sprintf("switched to tracking for %s", d.Family)
			delete(rollingData.DeviceLocation, d.Device)
		}
	} else {
		if !rollingData.HasData {
			rollingData.Timestamp = time.Now()
			rollingData.Datas = []models.SensorData{}
			rollingData.HasData = true
		}
		if len(d.Sensors["wifi"]) == 0 {
			return errors.New("no fingerprints")
		} else {
			rollingData.Datas = append(rollingData.Datas, d)
			message = fmt.Sprintf("inserted %d fingerprints for %s", len(d.Sensors["wifi"]), d.Family)
		}
	}
	err = db.Set("ReverseRollingData", rollingData)
	if err != nil {
		return
	}
	logger.Log.Debugf("success: %v, %s", success, message)
	c.JSON(http.StatusOK, gin.H{"message": message, "success": success})
	return
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
	if rollingData.HasData && time.Since(rollingData.Timestamp) > 13*time.Second {
		logger.Log.Debugf("%s has new data, %s", family, time.Since(rollingData.Timestamp))
		// merge data
		for _, data := range rollingData.Datas {
			for mac := range data.Sensors["wifi"] {
				rssi := data.Sensors["wifi"][mac]
				if _, ok := sensorMap[mac]; !ok {
					location := ""
					// if there is a device+location in map, then it is currently doing learning
					if len(rollingData.DeviceLocation) > 0 {
						if loc, hasMac := rollingData.DeviceLocation[mac]; hasMac {
							location = loc
						} else {
							continue
						}
					}
					sensorMap[mac] = models.SensorData{
						Family:    family,
						Device:    mac,
						Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
						Sensors:   make(map[string]map[string]interface{}),
						Location:  location,
					}
					time.Sleep(10 * time.Millisecond)
					sensorMap[mac].Sensors["wifi"] = make(map[string]interface{})
				}
				sensorMap[mac].Sensors["wifi"][data.Device] = rssi
			}
		}
		rollingData.HasData = false
	}
	db.Set("ReverseRollingData", rollingData)
	db.Close()
	for sensor := range sensorMap {
		logger.Log.Debugf("saving reverse sensor data for %s", sensor)
		logger.Log.Debugf("%+v", sensorMap[sensor])
		err := api.SaveSensorData(sensorMap[sensor])
		if err != nil {
			logger.Log.Warnf("problem saving: %s", err.Error())
		}
	}

	return
}

func checkRolingData() {
	for {
		time.Sleep(1 * time.Second)
		for _, family := range database.GetFamilies() {
			parseRollingData(family)
		}
	}
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
		err2 := processFingerprint(d)
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

func processFingerprint(p models.SensorData) (err error) {
	err = api.SaveSensorData(p)
	if err != nil {
		return
	}

	go sendOutData(p)
	return
}

func sendOutData(p models.SensorData) (analysis models.LocationAnalysis, err error) {
	analysis, err = api.AnalyzeSensorData(p)
	if err != nil {
		return
	}

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
	SendMessageOverWebsockets(p.Family, bTarget)
	mqtt.Publish(p.Family, p.Device, string(bTarget))
	return
}

func middleWareHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log request
		logger.Log.Infof("%v %v %v", c.Request.RemoteAddr, c.Request.Method, c.Request.URL)
		// Add base headers
		addCORS(c)
		// Run next function
		c.Next()
	}
}

func addCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
}

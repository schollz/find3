package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/schollz/find2/server/main/src/api"
	"github.com/schollz/find2/server/main/src/database"
	"github.com/schollz/find2/server/main/src/models"
	"github.com/schollz/find2/server/main/src/mqtt"
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
	r.GET("/api/v1/devices/*family", handlerApiV1Devices)
	r.GET("/api/v1/location/:family/*device", handlerApiV1Location)
	r.GET("/api/v1/calibrate/*family", handlerApiV1Calibrate)
	r.GET("/ping", ping)
	r.GET("/test", handleTest)
	r.GET("/ws", wshandler)            // handler for the web sockets (see websockets.go)
	r.POST("/mqtt", handlerMQTT)       // handler for setting MQTT
	r.POST("/data", handlerData)       // typical data handler
	r.POST("/reverse", handlerReverse) // typical data handler
	r.POST("/learn", handlerFIND)      // backwards-compatible with FIND for learning
	r.POST("/track", handlerFIND)      // backwards-compatible with FIND for tracking
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

func handleTest(c *gin.Context) {
	go api.Calibrate("pike", true)
	c.String(http.StatusOK, "ok")
}

func handlerApiV1Devices(c *gin.Context) {
	err := func(c *gin.Context) (err error) {
		family := strings.TrimSpace(c.Param("family")[1:])
		d, err := database.Open(family, true)
		defer d.Close()
		if err != nil {
			return
		}
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
	var err error
	var message string
	var d models.SensorData
	err = c.BindJSON(&d)
	if err == nil {
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
		message += fmt.Sprintf(", now learning on %d devices: %+v", len(rollingData.DeviceLocation), rollingData.DeviceLocation)
	} else {
		if !rollingData.HasData {
			rollingData.Timestamp = time.Now()
			rollingData.Datas = []models.SensorData{}
			rollingData.HasData = true
		}
		if len(d.Sensors) == 0 {
			return errors.New("no fingerprints")
		} else {
			rollingData.Datas = append(rollingData.Datas, d)
			numFingerprints := 0
			for sensor := range d.Sensors {
				numFingerprints += len(d.Sensors[sensor])
			}
			message = fmt.Sprintf("inserted %d fingerprints for %s", numFingerprints, d.Family)
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
	if rollingData.HasData && time.Since(rollingData.Timestamp) > 20*time.Second {
		logger.Log.Debugf("%s has new data, %s", family, time.Since(rollingData.Timestamp))
		// merge data
		for _, data := range rollingData.Datas {
			for sensor := range data.Sensors {
				for mac := range data.Sensors[sensor] {
					rssi := data.Sensors[sensor][mac]
					trackedDeviceName := sensor + "-" + mac
					if _, ok := sensorMap[trackedDeviceName]; !ok {
						location := ""
						// if there is a device+location in map, then it is currently doing learning
						if loc, hasMac := rollingData.DeviceLocation[mac]; hasMac {
							location = loc
						}
						sensorMap[trackedDeviceName] = models.SensorData{
							Family:    family,
							Device:    trackedDeviceName,
							Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
							Sensors:   make(map[string]map[string]interface{}),
							Location:  location,
						}
						time.Sleep(10 * time.Millisecond)
						sensorMap[trackedDeviceName].Sensors[sensor] = make(map[string]interface{})
					}
					sensorMap[trackedDeviceName].Sensors[sensor][data.Device] = rssi
				}
			}
		}
		rollingData.HasData = false
	}
	db.Set("ReverseRollingData", rollingData)
	db.Close()
	for sensor := range sensorMap {
		logger.Log.Debugf("saving reverse sensor data for %s", sensor)
		logger.Log.Debugf("%+v", sensorMap[sensor])
		err := processSensorData(sensorMap[sensor])
		if err != nil {
			logger.Log.Warnf("problem saving: %s", err.Error())
		}
	}

	return
}

func checkRolingData() {
	for {
		time.Sleep(6 * time.Second)
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
	SendMessageOverWebsockets(p.Family, p.Device, bTarget)
	mqtt.Publish(p.Family, p.Device, string(bTarget))
	return
}

func middleWareHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
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

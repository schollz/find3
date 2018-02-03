package server

import (
	"encoding/json"
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
func Run() {
	// setup MQTT
	mqtt.Setup()

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
	r.POST("/learn", handlerFIND)            // backwards-compatible with FIND for learning
	r.POST("/track", handlerFIND)            // backwards-compatible with FIND for tracking
	r.GET("/location", handlerLocation)      // get the latest location
	r.POST("/calibrate", handlerCalibration) // calibrate to get the latest location
	logger.Log.Infof("Running on 0.0.0.0:%s", Port)
	r.Run(":" + Port) // listen and serve on 0.0.0.0:8080
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

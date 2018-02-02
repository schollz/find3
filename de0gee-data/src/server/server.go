package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/de0gee/de0gee-data/src/database"
	"github.com/de0gee/de0gee-data/src/logging"
	"github.com/de0gee/de0gee-data/src/mqtt"
	"github.com/gin-gonic/gin"
)

// Port defines the public port
var Port = "8003"

var log = logging.Log

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
	r.GET("/ws", wshandler)             // handler for the web sockets (see websockets.go)
	r.POST("/mqtt", handlerMQTT)        // handler for setting MQTT
	r.POST("/data", handlerData)        // typical data handler
	r.POST("/learn", handlerFIND)       // backwards-compatible with FIND for learning
	r.POST("/track", handlerFIND)       // backwards-compatible with FIND for tracking
	r.GET("/location", handlerLocation) // get the latest location
	r.Run(":" + Port)                   // listen and serve on 0.0.0.0:8080
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

func handlerLocation(c *gin.Context) {
	type Payload struct {
		Family string `json:"family" binding:"required"`
		Device string `json:"device" binding:"required"`
	}
	success := false
	var message string
	var p Payload
	if errBind := c.ShouldBindJSON(&p); errBind == nil {
		d, err := database.Open(p.Family, true)
		if err != nil {
			message = err.Error()
		} else {
			defer d.Close()
			s, err := d.GetLatest(p.Device)
			if err != nil {
				message = err.Error()
			} else {
				target, err := d.Classify(s)
				if err != nil {
					message = err.Error()
				} else {

					bTarget, err := json.Marshal(target)
					if err != nil {
						fmt.Println(err)
					}
					SendMessageOverWebsockets(p.Family, bTarget)
					mqtt.Publish(p.Family, p.Device, string(bTarget))

					c.JSON(http.StatusOK, gin.H{"message": "got latest", "success": true, "response": target})
					return
				}
			}
		}
	} else {
		message = errBind.Error()
	}
	c.JSON(http.StatusOK, gin.H{"message": message, "success": success})
}

func handlerData(c *gin.Context) {
	var err error
	var message string
	var d database.SensorData
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
	var j database.FINDFingerprint
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

func processFingerprint(d database.SensorData) (err error) {
	err = d.Save()
	if err != nil {
		return
	}

	// TODO: use MQTT to push the latest classification
	return
}

func middleWareHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log request
		log.Debug(fmt.Sprintf("%v %v %v", c.Request.RemoteAddr, c.Request.Method, c.Request.URL))
		// Add base headers
		AddCORS(c)
		// Run next function
		c.Next()
	}
}

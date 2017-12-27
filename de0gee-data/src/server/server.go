package server

import (
	"net/http"

	"github.com/de0gee/de0gee-data/src/database"
	"github.com/gin-gonic/gin"
)

func Run(port string) {
	r := gin.Default()
	r.GET("/ping", ping)
	r.HEAD("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.POST("/", handlerData)      // typical data handler
	r.POST("/learn", handlerFIND) // backwards-compatible with FIND
	r.POST("/track", handlerFIND) // backwards-compatible with FIND
	r.Run(":" + port)             // listen and serve on 0.0.0.0:8080
}

func handlerData(c *gin.Context) {
	AddCORS(c)
	var err error
	var message string
	var d database.SensorData
	err = c.BindJSON(&d)
	if err == nil {
		err2 := d.Save()
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
	AddCORS(c)
	var j database.FINDFingerprint
	var err error
	var message string
	err = c.BindJSON(&j)
	if err == nil {
		if c.Request.URL.Path == "/track" {
			j.Location = ""
		}
		d := j.Convert()
		err2 := d.Save()
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

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

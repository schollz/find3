package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(port string) {
	r := gin.Default()
	r.GET("/ping", ping)
	r.HEAD("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.POST("/learn", learn) // backwards-compatible with FIND
	r.Run(":" + port)       // listen and serve on 0.0.0.0:8080
}

func learn(c *gin.Context) {
	var j FINDFingerprint
	var err error
	var message string
	err = c.BindJSON(&j)
	if err == nil {
		fmt.Println(j)
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

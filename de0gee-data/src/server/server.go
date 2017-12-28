package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/de0gee/de0gee-data/src/database"
	"github.com/gin-gonic/gin"
	cache "github.com/robfig/go-cache"
	log "github.com/sirupsen/logrus"
)

var (
	httpClient   *http.Client
	aiClientPort string
	routeCache   *cache.Cache
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
	routeCache = cache.New(5*time.Minute, 10*time.Minute)
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return client
}

func Run(port string, aiPort string) {
	aiClientPort = aiPort
	r := gin.Default()
	r.GET("/ping", ping)
	r.HEAD("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.POST("/data", handlerData)  // typical data handler
	r.POST("/learn", handlerFIND) // backwards-compatible with FIND
	r.POST("/track", handlerFIND) // backwards-compatible with FIND
	r.GET("/location", handlerLocation)
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}

func handlerLocation(c *gin.Context) {
	logger := log.WithFields(log.Fields{
		"name": "handlerLocation",
	})
	start := time.Now()
	AddCORS(c)
	type Payload struct {
		Family string `json:"family" binding:"required"`
		Device string `json:"device" binding:"required"`
	}
	success := false
	var message string
	var p Payload
	if errBind := c.ShouldBindJSON(&p); errBind == nil {
		d, err := database.Open(p.Family)
		defer d.Close()
		if err != nil {
			message = err.Error()
		} else {
			s, err := d.GetLatest(p.Device)
			if err != nil {
				message = err.Error()
			} else {
				type AIResponse struct {
					Data struct {
						LocationNames map[string]string `json:"location_names"`
						Predictions   []struct {
							Locations     []string  `json:"locations"`
							Name          string    `json:"name"`
							Probabilities []float64 `json:"probabilities"`
						} `json:"predictions"`
					} `json:"data"`
					Message string `json:"message"`
					Success bool   `json:"success"`
				}
				var target AIResponse

				cachedName := fmt.Sprintf("%s-%s-%f", p.Family, p.Device, s.Timestamp)
				responseCache, found := routeCache.Get(cachedName)
				logger.Info(found)
				if found {
					logger.Info("using cache")
					target = responseCache.(AIResponse)
				} else {
					type ClassifyPayload struct {
						Sensor       database.SensorData `json:"sensor_data"`
						DataLocation string              `json:"data_location"`
					}
					var p2 ClassifyPayload
					p2.Sensor = s
					dir, err := os.Getwd()
					if err != nil {
						logger.Fatal(err)
					}
					p2.DataLocation = dir
					logger.Info(time.Since(start))
					url := "http://localhost:" + aiClientPort + "/classify"
					bPayload, err := json.Marshal(p2)
					if err != nil {
						panic(err)
					}
					logger.Info(time.Since(start))
					req, err := http.NewRequest("POST", url, bytes.NewBuffer(bPayload))
					req.Header.Set("Content-Type", "application/json")
					resp, err := httpClient.Do(req)
					if err != nil {
						panic(err)
					}
					defer resp.Body.Close()
					logger.Info(time.Since(start))
					err = json.NewDecoder(resp.Body).Decode(&target)
					if err != nil {
						panic(err)
					}
					routeCache.Set(cachedName, target, 10*time.Second)
				}
				c.JSON(http.StatusOK, gin.H{"message": "got latest", "success": true, "response": target})
				logger.Info(time.Since(start))
				return
			}
		}
	} else {
		message = errBind.Error()
	}
	c.JSON(http.StatusOK, gin.H{"message": message, "success": success})
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

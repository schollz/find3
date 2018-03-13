package main

import (
	"flag"
	"log"
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"time"

	"fmt"

	"github.com/schollz/find3/server/main/src/api"
	"github.com/schollz/find3/server/main/src/database"
	"github.com/schollz/find3/server/main/src/mqtt"
	"github.com/schollz/find3/server/main/src/server"
)

func main() {
	aiPort := flag.String("ai", "8002", "port for the AI server")
	port := flag.String("port", "8003", "port for the data (this) server")
	debug := flag.Bool("debug", false, "turn on debug mode")
	mqttServer := flag.String("mqtt-server", "", "add MQTT server")
	mqttAdmin := flag.String("mqtt-admin", "admin", "name for mqtt admin")
	mqttPass := flag.String("mqtt-pass", "1234", "password for mqtt admin")
	dump := flag.String("dump", "", "family database to dump")
	memprofile := flag.Bool("memprofile", false, "whether to profile memory")
	var dataFolder string
	flag.StringVar(&dataFolder, "data", "", "location to store data")

	flag.Parse()

	if dataFolder == "" {
		dataFolder, _ = os.Getwd()
		dataFolder = path.Join(dataFolder, "data")
	}
	os.MkdirAll(dataFolder, 0775)

	// setup folders
	database.DataFolder = dataFolder
	api.DataFolder = dataFolder

	// setup debugging
	database.Debug(*debug)
	api.Debug(*debug)
	server.Debug(*debug)
	mqtt.Debug = *debug

	if os.Getenv("MQTT_ADMIN") != "" {
		mqtt.AdminUser = os.Getenv("MQTT_ADMIN")
	} else {
		mqtt.AdminUser = *mqttAdmin
	}
	if os.Getenv("MQTT_PASS") != "" {
		mqtt.AdminPassword = os.Getenv("MQTT_PASS")
	} else {
		mqtt.AdminPassword = *mqttPass
	}
	if os.Getenv("MQTT_SERVER") != "" {
		mqtt.Server = os.Getenv("MQTT_SERVER")
	} else {
		mqtt.Server = *mqttServer
	}

	api.AIPort = *aiPort
	server.Port = *port
	server.UseMQTT = mqtt.Server != ""

	if *memprofile {
		memprofilePath := path.Join(dataFolder, "memprofile")
		os.MkdirAll(memprofilePath, 0755)
		go func() {
			for {
				time.Sleep(1 * time.Second)
				log.Println("profiling memory")
				f, err := os.Create(path.Join(memprofilePath, fmt.Sprintf("%d.memprofile", time.Now().UnixNano()/int64(time.Millisecond))))
				if err != nil {
					log.Fatal("could not create memory profile: ", err)
				}
				runtime.GC() // get up-to-date statistics
				if err := pprof.WriteHeapProfile(f); err != nil {
					log.Fatal("could not write memory profile: ", err)
				}
				f.Close()
				time.Sleep(60 * time.Second)
			}
		}()
	}
	var err error
	if *dump != "" {
		err = api.Dump(*dump)
	} else {
		err = server.Run()
	}
	if err != nil {
		fmt.Print("error: ")
		fmt.Println(err)
	}
}

package main

import (
	"flag"
	"os"
	"path"

	"fmt"

	"github.com/de0gee/de0gee-data/src/mqtt"
	"github.com/schollz/find3/server/main/src/api"
	"github.com/schollz/find3/server/main/src/database"
	"github.com/schollz/find3/server/main/src/server"
)

func main() {
	externalAddress := flag.String("external", "127.0.0.1:8003", "external address")
	aiPort := flag.String("ai", "8002", "port for the AI server")
	port := flag.String("port", "8003", "port for the data (this) server")
	debug := flag.Bool("debug", false, "turn on debug mode")
	mqttFlag := flag.Bool("mqtt", false, "turn on mqtt")
	sslFlag := flag.Bool("ssl", false, "using SSL")
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

	api.AIPort = *aiPort
	server.Port = *port
	if os.Getenv("EXTERNAL_ADDRESS") != "" {
		server.ExternalServerAddress = os.Getenv("EXTERNAL_ADDRESS")
	} else {
		server.ExternalServerAddress = *externalAddress
	}

	server.UseSSL = *sslFlag
	server.UseMQTT = *mqttFlag
	err := server.Run()
	if err != nil {
		fmt.Print("error: ")
		fmt.Println(err)
	}
}

package main

import (
	"flag"

	"fmt"

	"github.com/de0gee/de0gee-data/src/api"
	"github.com/de0gee/de0gee-data/src/database"
	"github.com/de0gee/de0gee-data/src/mqtt"
	"github.com/de0gee/de0gee-data/src/server"
)

func main() {
	aiPort := flag.String("ai", "8002", "port for the AI server")
	port := flag.String("port", "8003", "port for the data (this) server")
	debug := flag.Bool("debug", false, "turn on debug mode")
	flag.Parse()

	// setup debugging
	database.Debug(*debug)
	api.Debug(*debug)
	server.Debug(*debug)
	mqtt.Debug = *debug

	api.AIPort = *aiPort
	server.Port = *port
	err := server.Run()
	if err != nil {
		fmt.Print("error: ")
		fmt.Println(err)
	}
}

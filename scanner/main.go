package main

import (
	log "github.com/cihub/seelog"
)

func main() {
	defer log.Flush()
	setLogLevel("debug")
	log.Info("starting")
}

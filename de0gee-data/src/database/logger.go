package database

import "github.com/de0gee/de0gee-data/src/logging"

var DebugMode bool
var logger *logging.SeelogWrapper

func init() {
	var err error
	logger, err = logging.New()
	if err != nil {
		panic(err)
	}
	Debug(false)
}

func Debug(debugMode bool) {
	if debugMode {
		logger.SetLevel("debug")
	} else {
		logger.SetLevel("info")
	}
}

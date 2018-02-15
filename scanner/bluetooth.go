package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/cihub/seelog"
)

// sudo apt-get install bluez
// use btmgmt find instead
var negativeNumberRegex = regexp.MustCompile(`-\d+`)
var macAddressRegex = regexp.MustCompile(`([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`)

func scanBluetooth(out chan map[string]map[string]interface{}) {
	log.Info("scanning bluetooth")
	s := btmgmtFind()
	data := make(map[string]map[string]interface{})
	data["bluetooth"] = make(map[string]interface{})
	for _, line := range strings.Split(s, "\n") {
		if negativeNumberRegex.MatchString(line) && macAddressRegex.MatchString(line) {
			rssi, err := strconv.Atoi(negativeNumberRegex.FindString(line))
			if err != nil {
				log.Warn(err)
				continue
			}
			data["bluetooth"][strings.ToLower(macAddressRegex.FindString(line))] = rssi
		}
	}
	out <- data
}

func hcitoolLescan() {
	RunCommand(4000*time.Millisecond, "hcitool lescan")
	log.Debug("finished lescan")
}

func btmgmtFind() string {
	for i := 0; i < 30; i++ {
		stdOut, stdErr := RunCommand(20*time.Second, "btmgmt find")
		if !strings.Contains(stdErr, "Unable to start") && len(stdOut) != 0 {
			log.Debug("finished btmgmt find")
			return stdOut
		}
		RunCommand(20*time.Second, "service bluetooth restart")
		time.Sleep(2 * time.Second)
	}
	return ""
}

func btmon(out chan string) {
	s, t := RunCommand(8000*time.Millisecond, "btmon")
	log.Debug("finished btmon")
	out <- s
	out <- t
}

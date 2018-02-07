package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	log "github.com/cihub/seelog"
)

// sudo apt-get install bluez
// use btmgmt find instead

func scanBluetooth() map[string]interface{} {
	// log.Println("reseting bluetooth")
	// log.Println(RunCommand(1*time.Second, "service", "bluetooth", "restart"))
	// time.Sleep(2 * time.Second)
	c := make(chan string)
	log.Debug("starting btmon")
	go btmon(c)
	time.Sleep(1500 * time.Millisecond)
	log.Debug("starting btmgmt")
	go btmgmtFind()
	s, _ := <-c, <-c
	ioutil.WriteFile("out", []byte(s), 0644)
	name := ""
	rssi := 0
	datas := make(map[string]interface{})
	for _, line := range strings.Split(s, "\n") {
		if strings.Contains(line, "Address:") {
			foo := strings.Split(line, "Address:")[1]
			foo = strings.Split(foo, "(")[0]
			foo = strings.ToLower(foo)
			foo = strings.TrimSpace(foo)
			name = foo
		} else if strings.Contains(line, "RSSI:") {
			foo := strings.Split(line, "RSSI:")[1]
			foo = strings.Split(foo, "dB")[0]
			foo = strings.TrimSpace(foo)
			var err error
			rssi, err = strconv.Atoi(foo)
			if err != nil {
				panic(err)
			}
		}
		if name != "" && rssi != 0 {
			datas[name] = rssi
			name = ""
			rssi = 0
		}
	}
	return datas
}

func hcitoolLescan() {
	RunCommand(4000*time.Millisecond, "hcitool lescan")
	log.Debug("finished lescan")
}

func btmgmtFind() {
	RunCommand(6000*time.Millisecond, "btmgmt find")
	log.Debug("finished btmgmt find")
}

func btmon(out chan string) {
	s, t := RunCommand(8000*time.Millisecond, "btmon")
	log.Debug("finished btmon")
	out <- s
	out <- t
}

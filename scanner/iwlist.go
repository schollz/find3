package main

import (
	"strconv"
	"strings"
	"time"
)

func iwlist() map[string]interface{} {
	s, _ := RunCommand(10*time.Second, "/sbin/iwlist "+wifiInterface+" scan")
	name := ""
	signal := 0
	datas := make(map[string]interface{})
	for _, line := range strings.Split(s, "\n") {
		if strings.Contains(line, "Address:") {
			name = strings.Split(line, "Address:")[1]
			name = strings.ToLower(name)
			name = strings.TrimSpace(name)
		} else if strings.Contains(line, "Signal level=") {
			foo := strings.Split(line, "Signal level=")[1]
			foo = strings.Split(foo, "dBm")[0]
			foo = strings.TrimSpace(foo)
			var err error
			signal, err = strconv.Atoi(foo)
			if err != nil {
				panic(err)
			}
		}
		if name != "" && signal != 0 {
			datas[name] = signal
		}
	}
	return datas
}

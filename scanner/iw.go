package main

import (
	"strconv"
	"strings"
	"time"
)

func iw(out chan map[string]map[string]interface{}) {
	s, _ := RunCommand(10*time.Second, "/sbin/iw dev "+wifiInterface+" scan -u")
	name := ""
	signal := 0
	datas := make(map[string]map[string]interface{})
	datas["wifi"] = make(map[string]interface{})
	for _, line := range strings.Split(s, "\n") {
		if strings.Contains(line, "(on") {
			name = strings.Split(strings.Split(line, "(")[0], "BSS")[1]
			name = strings.ToLower(name)
			name = strings.TrimSpace(name)
		} else if strings.Contains(line, "signal:") {
			foo := strings.Split(line, "signal:")[1]
			foo = strings.Split(foo, ".")[0]
			foo = strings.TrimSpace(foo)
			var err error
			signal, err = strconv.Atoi(foo)
			if err != nil {
				panic(err)
			}
		}
		if name != "" && signal != 0 {
			datas["wifi"][name] = signal
		}
	}
	out <- datas
}

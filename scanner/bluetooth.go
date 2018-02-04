package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// sudo apt-get install bluez
// use btmgmt find instead

func scanBluetooth() map[string]interface{} {
	// log.Println("reseting bluetooth")
	// log.Println(runCommand(1*time.Second, "service", "bluetooth", "restart"))
	// time.Sleep(2 * time.Second)
	c := make(chan string)
	log.Println("starting btmon")
	go btmon(c)
	time.Sleep(1500 * time.Millisecond)
	log.Println("starting lescan")
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
	runCommand(4000*time.Millisecond, "hcitool", "lescan")
	log.Println("finished lescan")
}

func btmgmtFind() {
	runCommand(6000*time.Millisecond, "btmgmt", "find")
	log.Println("finished btmgmt find")
}

func btmon(out chan string) {
	s, t := runCommand(8000*time.Millisecond, "btmon")
	log.Println("finished btmon")
	log.Println(s, t)
	out <- s
	out <- t
}

func runCommand(tDuration time.Duration, command ...string) (string, string) {
	cmd := exec.Command(command[0])
	if len(command) > 0 {
		cmd = exec.Command(command[0], command[1:]...)
	}
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(tDuration):
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill: ", err)
		}
		log.Println("process killed as timeout reached")
	case err := <-done:
		if err != nil {
			log.Printf("process done with error = %v", err)
		} else {
			log.Print("process done gracefully without error")
		}
	}
	return strings.TrimSpace(outb.String()), strings.TrimSpace(errb.String())
}

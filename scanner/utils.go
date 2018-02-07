package main

import (
	"bytes"
	"os/exec"
	"strings"
	"time"

	log "github.com/cihub/seelog"
)

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
		log.Error(err)
	}
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(tDuration):
		if err := cmd.Process.Kill(); err != nil {
			log.Error("failed to kill: ", err)
		}
		log.Debug("process killed as timeout reached")
	case err := <-done:
		if err != nil {
			log.Warn("process done with error = %v", err)
		} else {
			log.Debug("process done gracefully without error")
		}
	}
	return strings.TrimSpace(outb.String()), strings.TrimSpace(errb.String())
}

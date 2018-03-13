package logging

import (
	"fmt"
	"os"

	seelog "github.com/cihub/seelog"
)

type SeelogWrapper struct {
	Log   seelog.LoggerInterface
	Level string
}

func (self *SeelogWrapper) init() (err error) {
	if "" == self.Level {
		self.Level = "debug"
	}

	self.Log = seelog.Disabled

	// https://en.wikipedia.org/wiki/ANSI_escape_code#3/4_bit
	// https://github.com/cihub/seelog/wiki/Log-levels
	appConfig := `
	<seelog minlevel="` + self.Level + `">
	<outputs formatid="stdout">
	<filter levels="debug,trace">
		<console formatid="debug"/>
	</filter>
	<filter levels="info">
		<console formatid="info"/>
	</filter>
	<filter levels="critical,error">
		<console formatid="error"/>
	</filter>
	<filter levels="warn">
		<console formatid="warn"/>
	</filter>
	</outputs>
	<formats>
		<format id="stdout"   format="%Date %Time [%LEVEL] [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />

		<format id="debug"   format="%Date %Time %EscM(37)[%LEVEL]%EscM(0) [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />
		<format id="info"    format="%Date %Time %EscM(36)[%LEVEL]%EscM(0) [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />
		<format id="warn"    format="%Date %Time %EscM(33)[%LEVEL]%EscM(0) [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />
		<format id="error"   format="%Date %Time %EscM(31)[%LEVEL]%EscM(0) [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />

	</formats>
	</seelog>
	`

	self.Log.Close()
	self.Log, err = seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	return
}

func (self *SeelogWrapper) isValidLevel(level string) bool {
	levels := [6]string{"debug", "trace", "info", "critical", "error", "warn"}
	for i := range levels {
		if levels[i] == level {
			return true
		}
	}
	return false
}

func (self *SeelogWrapper) SetLevel(level string) error {
	if !self.isValidLevel(level) {
		return fmt.Errorf("Mot a valid logging level")
	}
	self.Level = level
	return self.init()
}

func New() (*SeelogWrapper, error) {
	logger := new(SeelogWrapper)
	logger.Level = "debug"
	err := logger.init()
	return logger, err
}

// https://github.com/cihub/seelog/wiki/Custom-formatters
func pidLogFormatter(params string) seelog.FormatterFunc {
	return func(message string, level seelog.LogLevel, context seelog.LogContextInterface) interface{} {
		var pid = os.Getpid()
		return fmt.Sprintf("%v", pid)
	}
}

func init() {
	seelog.RegisterCustomFormatter("pidLogFormatter", pidLogFormatter)
}

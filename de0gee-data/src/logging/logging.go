package logging

import (
	"fmt"
	"os"

	seelog "github.com/cihub/seelog"
)

var (
	Verbose  bool = false
	Log      seelog.LoggerInterface
	LogLevel string = "debug"
)

// const (
// 	nocolor = 0
// 	red     = 31 // error critical
// 	green   = 32
// 	yellow  = 33 // warn
// 	blue    = 36 // info
// 	grey    = 37 // debug
// )

// https://github.com/cihub/seelog/wiki/Custom-formatters
func pidLogFormatter(params string) seelog.FormatterFunc {
	return func(message string, level seelog.LogLevel, context seelog.LogContextInterface) interface{} {
		var pid = os.Getpid()
		return fmt.Sprintf("%v", pid)
	}
}

func initLogging() {
	if Verbose {
		LogLevel = "trace"
	}

	Log = seelog.Disabled

	// https://en.wikipedia.org/wiki/ANSI_escape_code#3/4_bit
	// https://github.com/cihub/seelog/wiki/Log-levels
	appConfig := `
<seelog minlevel="` + LogLevel + `">
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

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	Log = logger

	// Logger.Trace("trace")
	// Logger.Debug("debug")
	// Logger.Info("info")
	// Logger.Warn("warn")
	// Logger.Error("error")
	// Logger.Critical("critical")

}

func init() {
	seelog.RegisterCustomFormatter("pidLogFormatter", pidLogFormatter)
	// seelog.RegisterCustomFormatter("LLogFormatter", LLogFormatter)
	initLogging()
}

func Debug(t bool) {
	if t {
		LogLevel = "debug"
	} else {
		LogLevel = "warn"
	}
	initLogging()
}

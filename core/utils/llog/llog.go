package llog

import "C"
import (
	"fmt"
	"log"
	"os"
	"time"

	"gitee.com/liumou_site/logger"
)

type Log struct {
	LogPath  string
	LogLevel string
	CLog     *logger.LocalLogger
	FLog     *log.Logger
	ToDay    string
}

func (l *Log) fileLog() *log.Logger {
	logFile := l.fileIO()
	defer logFile.Close()

	l.FLog.SetOutput(logFile)
	return l.FLog
}

func (l *Log) consoleLog() *logger.LocalLogger {
	return l.CLog
}

func (l *Log) fileIO() *os.File {
	date := time.Now().Format("20060102")
	if date != l.ToDay {
		l.ToDay = date
	}
	logFile, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", l.LogPath, l.ToDay),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)

	return logFile
}

func (l *Log) BotINFO(info_ string) {

}

func (l *Log) BotDEBUG(debug_ string) {

}

func (l *Log) BotWarning(warning_ string) {

}

func (l *Log) BotPANIC(panic_ string) {

}

func (l *Log) output(s string) string {

}

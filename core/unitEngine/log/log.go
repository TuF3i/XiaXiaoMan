package log

import (
	"XiaXiaoMan/core"
	"XiaXiaoMan/core/utils/llog"
	"log"
	"time"

	"gitee.com/liumou_site/logger"
)

func InitLogger() {
	flog := log.Logger{}
	flog.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	flog.SetPrefix("[FileLog]")

	clog := logger.NewLogger(1)

	date := time.Now().Format("20060102")

	core.Logger = &llog.Log{
		LogPath:  "core/data/log",
		LogLevel: "DEBUG",
		CLog:     clog,
		FLog:     &flog,
		ToDay:    date,
	}
}

package main

import (
	"XiaXiaoMan/core"
	"XiaXiaoMan/core/plugins"
	"XiaXiaoMan/core/unitEngine/aliBaiLian"
	"XiaXiaoMan/core/unitEngine/config"
	"XiaXiaoMan/core/unitEngine/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.InitConfig()
	log.InitLogger()
	//dao.InitDataBase()
	aliBaiLian.InitQwen()
	plugins.InitPlugins()

	//botEngine.InitEngine()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	//unitEngine.ReleaseResource()
	core.Logger.BotINFO("Bot Exited, Bye!")
}

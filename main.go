package main

import (
	"XiaXiaoMan/core"
	"XiaXiaoMan/core/unitEngine/config"
	"XiaXiaoMan/core/unitEngine/log"
)

func main() {
	config.InitConfig()
	log.InitLogger()

	core.Logger.BotINFO("114514")
	core.Logger.BotPANIC("good")
}

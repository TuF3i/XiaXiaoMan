package config

import (
	"XiaXiaoMan/core"
	"XiaXiaoMan/core/utils/config"
)

func InitConfig() {
	conf := config.ReadConfig(core.ConfigFilePath)
	credentials := config.ReadAuth(core.AuthFilePath)

	core.Credentials = credentials
	core.Conf = conf
}

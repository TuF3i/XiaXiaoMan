package core

import (
	"XiaXiaoMan/core/utils/llog"

	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"gorm.io/gorm"
)

var (
	AuthFilePath   string = "core/data/auth.yaml"
	ConfigFilePath string = "core/data/config.json"

	Credentials *token.QQBotCredentials
	Conf        Config
	Engine      openapi.OpenAPI
	DB          *gorm.DB
	Logger      *llog.Log

	ToDay string
)

type Config struct {
	EnableDebug  bool   `json:"ENABLEDEBUG"`
	MySQLIPAddr  string `json:"MYSQLIPADDR"`
	MySQLPort    string `json:"MYSQLPORT"`
	MySQLUser    string `json:"MYSQLUSER"`
	MySQLPasswd  string `json:"MYSQLPASSWD"`
	MySQLDB      string `json:"MYSQLDB"`
	ListenIPAddr string `json:"ListenIPAddr"`
	ListenPort   string `json:"ListenPort"`
	CallBackPath string `json:"CallBackPath"`
	LogFilePath  string `json:"LOGPATH"`
}

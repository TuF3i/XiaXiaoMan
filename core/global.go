package core

import (
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
)

var (
	AuthFilePath   string = "core/data/auth.yaml"
	ConfigFilePath string = "core/data/config.json"

	Credentials *token.QQBotCredentials
	Conf        Config
	Engine      openapi.OpenAPI
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
}

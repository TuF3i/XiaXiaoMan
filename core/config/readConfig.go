package config

import (
	"XiaXiaoMan/core"
	"encoding/json"
	"log"
	"os"

	"github.com/tencent-connect/botgo/token"
	"gopkg.in/yaml.v2"
)

func InitConfig() {
	readAuth()
	readConfig()
}

func readAuth() {
	//读取认证文件
	auth, err := os.ReadFile(core.AuthFilePath)
	if err != nil {
		log.Fatalf("Read AuthFile Error: %v", err)
		os.Exit(1)
	}

	//解析认证文件
	credentials := &token.QQBotCredentials{
		AppID:     "",
		AppSecret: "",
	}

	err = yaml.Unmarshal(auth, &credentials)
	if err != nil {
		log.Fatalf("Parse AuthFile Error: %v", err)
		os.Exit(1)
	}

	log.Println("credentials:", credentials)
	core.Credentials = credentials
}

func readConfig() {
	//创建Config结构体
	var conf core.Config

	//读取Config文件
	data, err := os.ReadFile(core.ConfigFilePath)
	if err != nil {
		log.Fatalf("Read Config File Error: %v", err)
		os.Exit(1)
	}

	//解析Json到结构体
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("Unmarshal Error: %v", err)
		os.Exit(1)
	}

	core.Conf = conf
}

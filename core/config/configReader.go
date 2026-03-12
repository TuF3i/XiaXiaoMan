package config

import (
	"strings"

	"github.com/spf13/viper"
)

func setDefaultSettings(v *viper.Viper) {
	v.SetDefault(XIAXIAOMAN_LLBOTCONFIG_URL, "")
	v.SetDefault(XIAXIAOMAN_LLBOTCONFIG_TOKEN, "")
}

func configLoader() (*Config, error) {
	// 初始化配置结构体
	conf := new(Config)
	// 初始化viper
	v := viper.New()
	//这样环境变量需要以 XIAXIAOMAN_ 开头，如 XIAXIAOMAN_HERTZ_LISTENADDR
	v.SetEnvPrefix("XIAXIAOMAN")
	// 加载默认配置
	setDefaultSettings(v)
	// 加载环境变量
	v.AutomaticEnv()
	//设置键名转换器（将环境变量中的 _ 映射到结构体的嵌套字段）
	//例如：DANMU_HERTZ_LISTEN_ADDR -> Hertz.ListenAddr
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// 解析配置到结构体
	if err := v.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}

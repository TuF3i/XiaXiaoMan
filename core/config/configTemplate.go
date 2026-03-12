package config

type Config struct {
	LLBotConfig LLBotConfig
}

type LLBotConfig struct {
	Url   string
	Token string
}

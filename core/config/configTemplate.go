package config

type Config struct {
	LLBotConfig    LLBotConfig
	LLMChatConfig  LLMChatConfig
	DataBaseConfig DataBaseConfig
}

type LLBotConfig struct {
	Url   string
	Token string
}

type LLMChatConfig struct {
	AliyunAPIKey  string
	AliyunBaseURL string
	ChatModel     string
}

type DataBaseConfig struct {
	RedisConfig RedisConfig
	PgdbConfig  PgdbConfig
}

type RedisConfig struct {
	Addr     string
	Password string
}

type PgdbConfig struct {
	Addr      string
	Port      string
	UserName  string
	Password  string
	DefaultDB string
}

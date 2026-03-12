package config

func SetupConfig() (*Config, error) {
	return configLoader()
}

package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	FrontendUrl string `toml:"frontend_url"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

// Конфигурация сервера
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}

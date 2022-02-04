package rest_api

// Config General instance for App
type Config struct {
	BindAddr string	`toml:bind_addr`
	LoggerLevel string `toml:logger_level`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LoggerLevel: ":debug",
	}
}


package config

type Config struct {
	TelegramAPI struct {
		Token   string
		BaseURL string
	}
	DB struct {
		Driver string
		DSN    string
	}
}

func NewConfig() *Config {
	cfg := &Config{}

	cfg.TelegramAPI.Token = ""
	cfg.TelegramAPI.BaseURL = ""

	cfg.DB.Driver = "sqlite3"
	cfg.DB.DSN = "telegram.db"
	return cfg
}

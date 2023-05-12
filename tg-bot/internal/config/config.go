package config

type Config struct {
	Bot    Bot
	SQLite SQLite
}

func NewConfig() (*Config, error) {
	return nil, nil
}

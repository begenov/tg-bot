package config

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

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

const path = "./.env"

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := gotenv.Load(path)

	if err != nil {
		return nil, err
	}

	cfg.TelegramAPI.Token = os.Getenv("TELEGRAM_API_TOKEN")
	cfg.TelegramAPI.BaseURL = os.Getenv("TELEGRAM_API_BASE_URL")

	cfg.DB.Driver = os.Getenv("DRIVER")
	cfg.DB.DSN = os.Getenv("DSN")
	log.Println(cfg)
	return cfg, nil
}

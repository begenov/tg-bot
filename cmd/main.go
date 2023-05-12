package main

import (
	"log"

	"github.com/begenov/tg-bot/internal/api/telegram"
	"github.com/begenov/tg-bot/internal/app"
	"github.com/begenov/tg-bot/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		return
	}
	telegramAPI, err := telegram.NewTelegramAPI(cfg.Bot.Token)
	if err != nil {
		return
	}

	app := app.NewApp(cfg, telegramAPI)

	if err := app.Run(); err != nil {
		log.Println(err)
	}
}

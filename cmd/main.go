package main

import (
	"log"

	"github.com/begenov/tg-bot/internal/app"
	"github.com/begenov/tg-bot/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		return
	}

	app := app.NewApp(cfg)

	if err := app.Run(); err != nil {
		log.Println(err)
	}
}

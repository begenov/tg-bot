package main

import (
	"log"

	"github.com/begenov/tg-bot/internal/api/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBotAPI(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}

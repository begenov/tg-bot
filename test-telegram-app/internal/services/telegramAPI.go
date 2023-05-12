package services

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAPI struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramAPI(token string) (*TelegramAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		return nil, err
	}

	bot.Debug = true
	return &TelegramAPI{bot: bot}, nil
}

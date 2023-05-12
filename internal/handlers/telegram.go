package handlers

import (
	"github.com/begenov/tg-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAPI struct {
	bot      *tgbotapi.BotAPI
	services *services.Service
}

func NewTelegramAPI(bot *tgbotapi.BotAPI, servces *services.Service) *TelegramAPI {

	return &TelegramAPI{
		bot:      bot,
		services: servces,
	}
}

func (api *TelegramAPI) StartTelegramAPI() error {
	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60

	updates := api.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := api.handleCommand(update.Message); err != nil {
				api.errorHandler()
			}
			continue
		}

		if err := api.hadnleMessage(update.Message); err != nil {
			api.errorHandler()
		}

	}

	return nil
}

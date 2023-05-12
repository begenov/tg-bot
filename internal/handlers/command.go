package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (api *TelegramAPI) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return api.handleStartCommand()
	default:
		return api.handleUnknownCommand()
	}
}

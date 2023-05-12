package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	commandStart = "start"
)

func (api *TelegramAPI) handleStartCommand() error {
	return nil
}

func (api *TelegramAPI) handleUnknownCommand() error {
	return nil
}

func (api *TelegramAPI) hadnleMessage(message tgbotapi.Message) error {
	return nil
}

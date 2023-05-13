package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
)

func (api *TelegramAPI) handleStartCommand(message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, "Здравствуйте, это Telegram-bot по поиску работы и сотрудников.\nВыберите язык:")
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("казахский", "kazakh"),
			tgbotapi.NewInlineKeyboardButtonData("русский", "russian"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard

	api.bot.Send(msg)

	return nil
}

func (api *TelegramAPI) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hello User")
	api.bot.Send(msg)
	return nil
}

func (api *TelegramAPI) hadnleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hello User")
	api.bot.Send(msg)
	return nil
}

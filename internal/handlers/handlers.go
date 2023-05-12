package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
)

func (api *TelegramAPI) handleStartCommand(message *tgbotapi.Message) error {
	// start
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Здравствуйте, это Telegram-bot по поиску работы и сотрудников.\nВыберите язык: русский/казахский")
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("русский", "russian"),
			tgbotapi.NewInlineKeyboardButtonData("казахский", "kazakh"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	api.bot.Send(msg)

	if message.Text == "русский" {
		// Handle Russian language
		// ...
	} else if message.Text == "казахский" {
		// Handle Kazakh language
		// ...
	}
	return nil
}

func (api *TelegramAPI) handleUnknownCommand(message *tgbotapi.Message) error {

	return nil
}

func (api *TelegramAPI) hadnleMessage(message *tgbotapi.Message) error {

	return nil
}

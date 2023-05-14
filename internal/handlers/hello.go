package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (api *TelegramAPI) Hello(message *tgbotapi.Message, chatId int64) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Здравствуйте, это Telegram-bot по поиску работы и сотрудников.\nВыберите язык:")
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("казахский", "kazakh"),
			tgbotapi.NewInlineKeyboardButtonData("русский", "russian"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard

	api.usermapa[chatId] = &User{Stage: 0}

	if _, err := api.bot.Send(msg); err != nil {
		log.Fatal()
	}
}

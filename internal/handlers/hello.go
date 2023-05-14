package handlers

import (
	"github.com/begenov/tg-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (api *TelegramAPI) Hello(message *tgbotapi.Message, chatId int64) {
	msg := tgbotapi.NewMessage(chatId, models.InfoTelega)
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(models.Kazakh, models.Kazakh),
			tgbotapi.NewInlineKeyboardButtonData(models.Russian, models.Russian),
		),
	)
	msg.ReplyMarkup = inlineKeyboard

	api.usermapa[chatId] = &User{Stage: 0}

	api.bot.Send(msg)
}

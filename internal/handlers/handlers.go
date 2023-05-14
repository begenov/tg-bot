package handlers

import (
	"github.com/begenov/tg-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (api *TelegramAPI) choseKazakh(update tgbotapi.Update, msg tgbotapi.MessageConfig, chatId int64) {
	lang := update.CallbackQuery.Data
	api.usermapa[chatId].Stage = 1

	msg2 := tgbotapi.NewMessage(chatId, "")
	if lang == models.Kazakh {
		msg.Text = models.ChoseKazakh
		msg2.Text = models.KazakhName
	} else if lang == models.Russian {
		msg.Text = models.ChoseRussian
		msg2.Text = models.RussianName
	}

	api.bot.Send(msg)
	api.bot.Send(msg2)

}

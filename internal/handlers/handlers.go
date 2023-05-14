package handlers

import (
	"github.com/begenov/tg-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (api *TelegramAPI) choseKazakhHandler(update tgbotapi.Update, msg tgbotapi.MessageConfig, chatId int64) {
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

// share telefon

// Share Name

func (api *TelegramAPI) nameHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	name := update.Message.Text
	api.usermapa[chatId].name = name
	api.usermapa[chatId].Stage = 2

	shareButton := tgbotapi.NewKeyboardButtonContact("")
	msg2 := tgbotapi.NewMessage(chatId, "")

	if api.usermapa[chatId].lang == models.Kazakh {
		msg.Text = models.KazakhHello + name

		shareButton.Text = models.KazakhNumberButton
	} else {
		msg.Text = models.RussianHello + name
		shareButton.Text = models.RussianNumberButton

	}

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(shareButton),
	)

	api.bot.Send(msg)

	msg2.ReplyMarkup = keyboard
	api.bot.Send(msg2)
}

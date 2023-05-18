package handlers

import (
	"strconv"

	"github.com/begenov/tg-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (api *TelegramAPI) jobSeekersHandler(update tgbotapi.Update, msg tgbotapi.MessageConfig, chatId int64) {
	switch api.usermapa[chatId].Stage {
	case 1:
		if update.CallbackQuery != nil {
			api.workFieldHandler(update, chatId, msg)
		}
	case 2:
		if update.CallbackQuery != nil {
			api.jobHandler(update, chatId, msg)
		}
	case 10:
		api.anotherArea(update, chatId, msg)
	default:
		msg.Text = "В какой сфере вы бы хотели найти работу?"
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Торговля", "1"),
				tgbotapi.NewInlineKeyboardButtonData("Общепит", "2"),
				tgbotapi.NewInlineKeyboardButtonData("Другое", "3"),
				tgbotapi.NewInlineKeyboardButtonData("Пропустить шаг", "4"),
			),
		)

		msg.ReplyMarkup = inlineKeyboard
		api.bot.Send(msg)
		api.usermapa[chatId].Stage = 1
	}
	// api.workFieldHandler(update, chatId, msg)
}

func (api *TelegramAPI) workFieldHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	workField, _ := strconv.Atoi(update.CallbackQuery.Data)

	switch workField {
	case 1:
		api.usermapa[chatId].Stage = 3
	case 2:
		api.usermapa[chatId].Stage = 3
	case 3:
		msg.Text = "Напиши сферу"

		api.bot.Send(msg)
		api.usermapa[chatId].Stage = 10
		return
	case 4:
		api.usermapa[chatId].Stage = 3

	}

	msg.Text = "На какой работе хотите работать"

	Jobs := models.Field[workField]
	keyboard := tgbotapi.NewInlineKeyboardRow()
	for _, job := range Jobs {
		button := tgbotapi.NewInlineKeyboardButtonData(job, job)
		keyboard = append(keyboard, button)
	}
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(keyboard)
	msg.ReplyMarkup = inlineKeyboard
	api.bot.Send(msg)
	api.usermapa[chatId].Stage = 2
}

func (api *TelegramAPI) jobHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
}

func (api *TelegramAPI) anotherArea(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	api.usermapa[chatId].Field = msg.Text
}

package handlers

import (
	"log"
	"strconv"

	"github.com/begenov/tg-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (api *TelegramAPI) jobSeekersHandler(update tgbotapi.Update, msg tgbotapi.MessageConfig, chatId int64) {
	switch api.usermapa[chatId].Stage {
	case 0:

		// api.usermapa[chatId].Stage = 8
		msg.Text = "В какой сфере вы бы хотели найти работу?"
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Торговля", "0"),
				tgbotapi.NewInlineKeyboardButtonData("Общепит", "1"),
				tgbotapi.NewInlineKeyboardButtonData("Другое", "2"),
				tgbotapi.NewInlineKeyboardButtonData("Пропустить шаг", "3"),
			),
		)

		msg.ReplyMarkup = inlineKeyboard
		api.bot.Send(msg)
		api.usermapa[chatId].Stage = 1

	case 1:
		if update.CallbackQuery != nil {
			api.workFieldHandler(update, chatId, msg)
		}
	case 2:
		if update.CallbackQuery != nil {
			api.jobHandler(update, chatId, msg)
		}
	default:
		log.Println("-----------------------")
	}

	// api.workFieldHandler(update, chatId, msg)
}

func (api *TelegramAPI) workFieldHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	workField, _ := strconv.Atoi(update.CallbackQuery.Data)
	if workField == 2 {
		api.usermapa[chatId].Stage = 6
		return
	} else if workField == 3 {
		api.usermapa[chatId].Stage = 2
		return
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
	job := update.CallbackQuery.Data
	log.Fatal(job)
}

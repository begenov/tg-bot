package handlers

import (
	"fmt"
	"log"
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
	case 3:
		if update.CallbackQuery != nil {
			api.salaryHandler(update, chatId, msg)
		}
	case 4:
		if update.CallbackQuery != nil {
			api.salaryHandler(update, chatId, msg)
		}
	case 5:
		if update.CallbackQuery != nil {
			api.jobFinder(update, chatId, msg)
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
	api.usermapa[chatId].FieldId = workField
	if workField >= 0 && workField <= 2 {
		api.usermapa[chatId].Field = models.Fields[workField]
	} else {
		api.usermapa[chatId].Field = "Нужно указать! (Другое или что-нибудь еще)"
	}
	msg.Text = "Кем бы вы хотели работать?"

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
	api.usermapa[chatId].Job = job

	msg.Text = "Какую зарплату вы хотели бы получить?"
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("50,000 - 150,000", "50,000 - 150,000"),
			tgbotapi.NewInlineKeyboardButtonData("150,000 - 250,000", "150,000 - 250,000"),
			tgbotapi.NewInlineKeyboardButtonData("250,000 - 350,000", "250,000 - 350,000"),
			tgbotapi.NewInlineKeyboardButtonData("350,000 - 500,000", "350,000 - 500,000"),
			tgbotapi.NewInlineKeyboardButtonData("500,000 - 700,000", "500,000 - 700,000"),
			tgbotapi.NewInlineKeyboardButtonData("700,000 < ", "700,000 < "),
		),
	)

	msg.ReplyMarkup = inlineKeyboard
	api.bot.Send(msg)
	api.usermapa[chatId].Stage = 3

	// log.Fatal(job)
}

func (api *TelegramAPI) salaryHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	salary := update.CallbackQuery.Data
	api.usermapa[chatId].Salary = salary

	msg.Text = "Начать поиск подходящих вакансий?"
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ДА", "1"),
			tgbotapi.NewInlineKeyboardButtonData("НЕТ", "0"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	api.bot.Send(msg)
	api.usermapa[chatId].Stage = 5
}

func (api *TelegramAPI) jobFinder(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	answer, err := strconv.Atoi(update.CallbackQuery.Data)
	if err != nil {
		log.Fatal(err)
	}
	if answer == 0 {
		return
	}
	message := fmt.Sprintf("По вашему запросу: «%s – %s – %s тг.» найдено __ вакансий", api.usermapa[chatId].Field, api.usermapa[chatId].Job, api.usermapa[chatId].Salary)
	msg.Text = message
	api.bot.Send(msg)
}

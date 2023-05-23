package handlers

import (
	"context"
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
}

func (api *TelegramAPI) workFieldHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	workField, _ := strconv.Atoi(update.CallbackQuery.Data)

	user := api.usermapa[chatId]

	switch workField {
	case 1:
		api.usermapa[chatId].Stage = 3
	case 2:
		api.usermapa[chatId].Stage = 3
	case 3:
		text := "Напиши сферу"
		if user.Lang == models.Kazakh {
			text = "Сфераны жазыңыз"
		}
		msg.Text = text
		api.bot.Send(msg)
		api.usermapa[chatId].Stage = 5
		return
	case 4:
		api.usermapa[chatId].Stage = 3

	}
	api.usermapa[chatId].FieldId = workField
	if workField >= 0 && workField <= 2 {
		api.usermapa[chatId].Field = models.Fields[workField]
	} else {
		api.usermapa[chatId].Field = "Нужно указать! (Другое или что-нибудь еще)"
	}
	var Jobs []string

	if user.Lang == models.Kazakh {
		Jobs = models.FieldKZ[workField]
	} else {
		Jobs = models.Field[workField]
	}
	keyboard := tgbotapi.NewInlineKeyboardRow()
	for _, job := range Jobs {
		button := tgbotapi.NewInlineKeyboardButtonData(job, job)
		keyboard = append(keyboard, button)
	}
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(keyboard)
	msg.ReplyMarkup = inlineKeyboard
	text := "Кем бы вы хотели работать?"
	if user.Lang == models.Kazakh {
		text = "Сіз кім болып жұмыс істегіңіз келеді?"
	}
	msg.Text = text
	api.bot.Send(msg)
	api.usermapa[chatId].Stage = 2
}

func (api *TelegramAPI) jobHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	user := api.usermapa[chatId]

	job := update.CallbackQuery.Data
	api.usermapa[chatId].Job = job

	text := "Какую зарплату вы хотели бы получить?"
	if user.Lang == models.Kazakh {
		text = "Сіз қандай жалақы алғыңыз келеді?"
	}
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
	msg.Text = text

	msg.ReplyMarkup = inlineKeyboard
	api.bot.Send(msg)

	api.usermapa[chatId].Stage = 3

	// log.Fatal(job)
}

func (api *TelegramAPI) salaryHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	user := api.usermapa[chatId]

	salary := update.CallbackQuery.Data
	api.usermapa[chatId].Salary = salary

	if err := api.services.JobSeeker.CreateJobSeeker(context.Background(), models.JobSeeker{
		Sphere:     api.usermapa[chatId].Field,
		Profession: api.usermapa[chatId].Job,
		ChatID:     int(chatId),
		Salary:     api.usermapa[chatId].Salary,
	}); err != nil {
		log.Fatalln(err)
	}
	text := "Начать поиск подходящих вакансий?"
	confirm := "Да"
	decline := "Нет"
	if user.Lang == models.Kazakh {
		text = "Жұмыс орындарын іздеуді бастайық па?"
		confirm = "Иә"
		decline = "Жоқ"
	}
	msg.Text = text
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(confirm, "1"),
			tgbotapi.NewInlineKeyboardButtonData(decline, "0"),
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

	user := api.usermapa[chatId]

	text := "По вашему запросу: «%s – %s – %s тг.» найдено __ вакансий"

	if user.Lang == models.Kazakh {
		text = "Сіздің сұранысыңыз бойынша: «%s - %s- % S тг.» __жұмыс орны табылды"
	}

	message := fmt.Sprintf(text, api.usermapa[chatId].Field, api.usermapa[chatId].Job, api.usermapa[chatId].Salary)
	msg.Text = message
	api.bot.Send(msg)
}

func (api *TelegramAPI) anotherJob(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	api.usermapa[chatId].Field = msg.Text
	api.usermapa[chatId].Stage = 2
}

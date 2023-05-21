package handlers

import (
	"log"
	"strconv"

	"github.com/begenov/tg-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (api *TelegramAPI) jobPostingHandlers(update tgbotapi.Update, msg tgbotapi.MessageConfig, chatId int64) {
	log.Println(api.usermapa[chatId].Stage)
	switch api.usermapa[chatId].Stage {
	case 0:
		if update.CallbackQuery != nil {
			api.fillinGoutTheDataHandler(msg, chatId)
		}

	case 1:
		if update.CallbackQuery != nil {
			api.jobPostingWorkFieldHandler(update, msg, chatId)
		}
	case 2:
		if update.CallbackQuery != nil {
			api.handleSalarySelection(update, chatId, msg)
		}
	case 3:
		if update.CallbackQuery != nil {
			api.handleCompanyNameInput(update, chatId, msg)
		}
	case 4:
		if update.Message != nil {
			api.handleBINInput(update, chatId, msg)
		}
	case 5:
		if update.Message != nil {
			api.handleCandidateRequirementsInput(update, chatId, msg)
		}
	case 6:
		if update.Message != nil {
			api.handleCandidateResponsibilitiesInput(update, chatId, msg)
		}
	case 7:
		if update.Message != nil {
			api.handleVacancyPublication(update, chatId, msg)
		}
	default:
		api.createJobPostingHandler(msg, chatId)
	}
}

func (api *TelegramAPI) createJobPostingHandler(msg tgbotapi.MessageConfig, chatId int64) {
	msg.Text = "Отлично! Давайте приступим к поиску сотрудников!"
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Создать вакансию", "1"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	api.bot.Send(msg)
}

func (api *TelegramAPI) fillinGoutTheDataHandler(msg tgbotapi.MessageConfig, chatId int64) {
	msg.Text = "В какой сфере Вы работаете?"

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

func (api *TelegramAPI) jobPostingWorkFieldHandler(update tgbotapi.Update, msg tgbotapi.MessageConfig, chatId int64) {
	workField, _ := strconv.Atoi(update.CallbackQuery.Data)

	switch workField {
	case 1:
		api.usermapa[chatId].Stage = 3
	case 2:
		api.usermapa[chatId].Stage = 3
	case 3:
		msg.Text = "Напиши сферу"
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
	msg.Text = "На какую должность вы ищите сотрудника?"

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

func (api *TelegramAPI) handleSalarySelection(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	job := update.CallbackQuery.Data
	api.usermapa[chatId].Job = job

	msg.Text = "Какую зарплату Вы готовы предложить??"
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

func (api *TelegramAPI) handleCompanyNameInput(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	salary := update.CallbackQuery.Data
	api.usermapa[chatId].Salary = salary

	msg.Text = "Введите, пожалуйста, наименование вашей компании?"
	api.bot.Send(msg)
	api.usermapa[chatId].Stage = 4
}

func (api *TelegramAPI) handleBINInput(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	company := msg.Text
	api.usermapa[chatId].Company = company
	msg.Text = "Введите, пожалуйста, БИН вашей компании:"
	msg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true, Selective: true}
	api.bot.Send(msg)
	api.usermapa[chatId].Stage = 5
}

func (api *TelegramAPI) handleCandidateRequirementsInput(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	bin := msg.Text
	api.usermapa[chatId].BIN = bin
	msg.Text = "Какие ваши требования к кандидату?"
	msg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true, Selective: true}
	api.bot.Send(msg)
	api.usermapa[chatId].Stage = 6
}

func (api *TelegramAPI) handleCandidateResponsibilitiesInput(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	requirement := msg.Text
	api.usermapa[chatId].Requirement = requirement
	msg.Text = "Какие функциональные обязанности кандидата?"
	msg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true, Selective: true}
	api.bot.Send(msg)
	api.usermapa[chatId].Stage = 7
}

func (api *TelegramAPI) handleVacancyPublication(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	responsibilities := msg.Text
	api.usermapa[chatId].Responsibilities = responsibilities
	msg.Text = "Спасибо, ваша вакансия сохранена. Она будет обработана модератором. В скором времени вы получите уведомление."
	api.bot.Send(msg)
	api.usermapa[chatId].Stage = 8
}

package handlers

import (
	"log"

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

		}
	case 3:
		if update.CallbackQuery != nil {

		}
	case 4:
		if update.CallbackQuery != nil {

		}
	case 5:
		if update.CallbackQuery != nil {

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

}

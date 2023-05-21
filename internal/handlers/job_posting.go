package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (api *TelegramAPI) jobPostingHandlers(update tgbotapi.Update, msg tgbotapi.MessageConfig, chatId int64) {
	switch api.usermapa[chatId].Stage {
	case 0:
		api.createJobPostingHandler(msg, chatId)
	case 1:
		if update.CallbackQuery != nil {

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
	api.usermapa[chatId].Stage = 1

}

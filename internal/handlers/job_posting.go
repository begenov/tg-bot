package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (api *TelegramAPI) jobPostingHandlers(update tgbotapi.Update, msg tgbotapi.MessageConfig, chatId int64) {
	switch api.usermapa[chatId].Stage {
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

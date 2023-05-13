package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (api *TelegramAPI) handleRussianlanguage(message *tgbotapi.MessageConfig) {
	msg := tgbotapi.NewMessage(message.ChatID, "Выбрали русский язык")
	api.bot.Send(msg)
}

func (api *TelegramAPI) handleKazakhlanguage(message tgbotapi.MessageConfig) {
	msg := tgbotapi.NewMessage(message.ChatID, "Выбрали казахский язык")
	api.bot.Send(msg)
}

func (api *TelegramAPI) handleCallbackQuery(callbackQuery *tgbotapi.CallbackQuery) error {
	if callbackQuery.Data == "kazakh" {
		// пользователь выбрал казахский язык
	} else if callbackQuery.Data == "russian" {
		// пользователь выбрал русский язык
	} else {
		// обработка ошибки, если значение Data не соответствует ожидаемым значениям
	}

	// отправка подтверждения обработки callback-запроса

	return nil
}

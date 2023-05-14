package handlers

import (
	"fmt"

	"github.com/begenov/tg-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (api *TelegramAPI) choseKazakhHandler(update tgbotapi.Update, msg tgbotapi.MessageConfig, chatId int64) {
	lang := update.CallbackQuery.Data
	api.usermapa[chatId].Stage = 1
	api.usermapa[chatId].lang = lang

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

func (api *TelegramAPI) phoneNumberHandler1(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	phoneNumber := update.Message.Contact.PhoneNumber
	// log.Fatal(phoneNumber)
	api.usermapa[chatId].phone = phoneNumber
	api.usermapa[chatId].Stage = 4
	if api.usermapa[chatId].lang == "kazakh" {
		msg.Text = "Сізге коды бар SMS хабарлама келеді. Оны енгізіңіз."
	} else {
		msg.Text = "Вы получите SMS-уведомление с кодом. Введите его, пожалуйста."
	}
	api.bot.Send(msg)
}

func (api *TelegramAPI) phoneNumberHandler2(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	phoneNumber := update.Message.Text
	api.usermapa[chatId].phone = phoneNumber
	api.usermapa[chatId].Stage = 4
	if api.usermapa[chatId].lang == "kazakh" {
		msg.Text = "Сізге коды бар SMS хабарлама келеді. Оны енгізіңіз."
	} else {
		msg.Text = "Вы получите SMS-уведомление с кодом. Введите его, пожалуйста."
	}
	api.bot.Send(msg)
}

// Share Name

func (api *TelegramAPI) checkPhoneNumberHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	code := update.Message.Text
	if code == "0000" {
		if api.usermapa[chatId].lang == "kazakh" {
			msg.Text = "Сізге коды бар SMS хабарлама келеді. Оны енгізіңіз."
		} else {
			msg.Text = "Вы получите SMS-уведомление с кодом. Введите его, пожалуйста."
		}
		// api.usermapa[chatId].Stage = 5
		// // api.bot.Send(msg)
		// continue
	}
	msg.Text = "Ваш код получен, спасибо!"
	api.bot.Send(msg)
	msg1 := tgbotapi.NewMessage(chatId, "")
	msg1.Text = fmt.Sprintf("%s, спасибо за регистрацию! Вы ищите работу или сотрудника?", api.usermapa[chatId].name)
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ищу работу", "work"),
			tgbotapi.NewInlineKeyboardButtonData("ищу сотрудника", "employee"),
		),
	)
	msg1.ReplyMarkup = inlineKeyboard
	api.usermapa[chatId].Stage = 5
	api.bot.Send(msg1)
}

func (api *TelegramAPI) nameHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	name := update.Message.Text
	api.usermapa[chatId].name = name

	shareButton := tgbotapi.NewKeyboardButtonContact(models.RussianNumberButton)
	msg2 := tgbotapi.NewMessage(chatId, "")

	if api.usermapa[chatId].lang == models.Kazakh {

		msg.Text = models.KazakhHello + name
		msg2.Text = models.KazakhNumberButton
		shareButton.Text = models.KazakhNumberButton
	} else {
		msg.Text = models.RussianHello + name
		msg2.Text = models.RussianNumberButton
		shareButton.Text = models.RussianNumberButton

	}

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(shareButton),
	)

	api.bot.Send(msg)

	msg2.ReplyMarkup = keyboard
	api.bot.Send(msg2)
	api.usermapa[chatId].Stage = 2
}

func (api *TelegramAPI) coverLetterHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	api.usermapa[chatId].aim = update.CallbackQuery.Data
	api.usermapa[chatId].Stage = 6
	// fmt.Println("--------------The aim is ", api.usermapa[chatId].aim)
	if api.usermapa[chatId].lang == "kazakh" {
		msg.Text = "Сопроводительное письмо"
	} else {
		msg.Text = "Сопроводительное письмо"
	}
	api.bot.Send(msg)
}

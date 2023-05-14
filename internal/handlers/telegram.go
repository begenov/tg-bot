package handlers

import (
	"fmt"

	"github.com/begenov/tg-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAPI struct {
	bot      *tgbotapi.BotAPI
	services *services.Service
	usermapa map[int64]*User
}

type User struct {
	Stage int
	lang  string
	name  string
	phone string
	aim   string
}

func NewTelegramAPI(bot *tgbotapi.BotAPI, servces *services.Service) *TelegramAPI {
	return &TelegramAPI{
		bot:      bot,
		services: servces,
		usermapa: make(map[int64]*User),
	}
}

func (api *TelegramAPI) StartTelegramAPI() error {
	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60
	updates := api.bot.GetUpdatesChan(u)

	for update := range updates {
		chatId := update.FromChat().ID
		msg := tgbotapi.NewMessage(chatId, "")

		if _, exi := api.usermapa[chatId]; !exi {
			api.Hello(update.Message, chatId)
			continue
		}

		switch api.usermapa[chatId].Stage {
		case 0:
			if update.CallbackQuery != nil {
				api.choseKazakhHandler(update, msg, chatId)
				continue
			}
		case 1:
			if update.Message != nil {
				api.nameHandler(update, chatId, msg)
				continue
			}
		case 2:
			if update.Message.Contact != nil {
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
				continue
			}
			if update.Message != nil {
				phoneNumber := update.Message.Text
				api.usermapa[chatId].phone = phoneNumber
				api.usermapa[chatId].Stage = 4
				if api.usermapa[chatId].lang == "kazakh" {
					msg.Text = "Сізге коды бар SMS хабарлама келеді. Оны енгізіңіз."
				} else {
					msg.Text = "Вы получите SMS-уведомление с кодом. Введите его, пожалуйста."
				}
				api.bot.Send(msg)
				continue
			}
		case 4:
			if update.Message != nil {
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
				continue

			}
			// fallthrough
		case 5:
			if update.CallbackQuery != nil {
				api.usermapa[chatId].aim = update.CallbackQuery.Data
				api.usermapa[chatId].Stage = 6
				// fmt.Println("--------------The aim is ", api.usermapa[chatId].aim)
				if api.usermapa[chatId].lang == "kazakh" {
					msg.Text = "Сопроводительное письмо"
				} else {
					msg.Text = "Сопроводительное письмо"
				}
				api.bot.Send(msg)
				continue
			}
			// if update.Message != nil {
			// 	api.usermapa[chatId].aim = update.Message.Text
			// 	api.usermapa[chatId].Stage = 6
			// 	if api.usermapa[chatId].lang == "kazakh" {
			// 		msg.Text = "Сопроводительное письмо"
			// 	} else {
			// 		msg.Text = "Сопроводительное письмо"
			// 	}
			// 	api.bot.Send(msg)
			// 	continue
			// }

		}
	}

	return nil
}

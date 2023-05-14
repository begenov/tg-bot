package handlers

import (
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
				api.phoneNumberHandler1(update, chatId, msg)
				continue
			}
			if update.Message != nil {
				api.phoneNumberHandler2(update, chatId, msg)
				continue
			}
		case 4:
			if update.Message != nil {
				api.checkPhoneNumberHandler(update, chatId, msg)
				continue

			}
			// fallthrough
		case 5:
			if update.CallbackQuery != nil {
				api.coverLetterHandler(update, chatId, msg)
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

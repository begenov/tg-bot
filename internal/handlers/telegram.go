package handlers

import (
	"github.com/begenov/tg-bot/internal/models"
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
				api.choseKazakh(update, msg, chatId)
				continue
			}
		case 1:
			if update.Message != nil {
				name := update.Message.Text
				api.usermapa[chatId].name = name
				api.usermapa[chatId].Stage = 2

				shareButton := tgbotapi.NewKeyboardButtonContact("")
				msg2 := tgbotapi.NewMessage(chatId, "")

				if api.usermapa[chatId].lang == models.Kazakh {
					msg.Text = models.KazakhHello + name

					shareButton.Text = models.KazakhNumberButton
				} else {
					msg.Text = models.RussianHello + name
					shareButton.Text = models.RussianNumberButton

				}

				keyboard := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(shareButton),
				)

				api.bot.Send(msg)

				msg2.ReplyMarkup = keyboard
				api.bot.Send(msg2)
				continue
			}
		}
	}

	return nil
}

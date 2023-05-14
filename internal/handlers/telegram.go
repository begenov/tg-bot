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

		fmt.Println(api.usermapa[chatId], "oooooooooooooooooo")
		switch api.usermapa[chatId].Stage {
		case 0:
			if update.CallbackQuery != nil {
				lang := update.CallbackQuery.Data
				api.usermapa[chatId].Stage = 1
				fmt.Println(api.usermapa[chatId], "----------")

				msg2 := tgbotapi.NewMessage(chatId, "")
				if lang == "kazakh" {
					msg.Text = "сіз Қазақ тілің таңдадыңыз"
					msg2.Text = "Атыңызды енгізіңіз: "
				} else {
					msg.Text = "Вы выбрали русский язык"
					msg2.Text = "Введите имя: "
				}

				api.bot.Send(msg)
				api.bot.Send(msg2)
				continue
			}
		case 1:
			if update.Message != nil {
				name := update.Message.Text
				api.usermapa[chatId].name = name
				api.usermapa[chatId].Stage = 2

				shareButton := tgbotapi.NewKeyboardButtonContact("")
				msg2 := tgbotapi.NewMessage(chatId, "")

				if api.usermapa[chatId].lang == "kazakh" {
					msg.Text = "Сәлем" + name

					shareButton.Text = "Нөмірмен бөлісу"

					msg2.Text = "Ботқа тіркелу үшін бізге телефон нөміріңіз қажет."
				} else {
					msg.Text = "Здравствуйте, " + name
					shareButton.Text = "Для регистрации в боте нам нужен ваш номер телефона."
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

package handlers

import (
	"context"
	"log"

	"github.com/begenov/tg-bot/internal/models"
	"github.com/begenov/tg-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAPI struct {
	bot      *tgbotapi.BotAPI
	services *services.Service
	usermapa map[int64]*models.User
}

func NewTelegramAPI(bot *tgbotapi.BotAPI, servces *services.Service) *TelegramAPI {
	return &TelegramAPI{
		bot:      bot,
		services: servces,
		usermapa: make(map[int64]*models.User),
	}
}

func (api *TelegramAPI) StartTelegramAPI() error {
	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60
	updates := api.bot.GetUpdatesChan(u)

	for update := range updates {
		chatId := update.FromChat().ID
		user, err := api.services.User.UserByChatID(context.Background(), int(chatId))
		log.Println(user, err, "------------------------------------------")
		msg := tgbotapi.NewMessage(chatId, "")

		if user == (&models.User{}) {
			if _, exi := api.usermapa[chatId]; !exi {
				api.Hello(update.Message, chatId)
				continue
			}
			api.profileUser(update, chatId, msg)
			continue
		}

	}

	return nil
}

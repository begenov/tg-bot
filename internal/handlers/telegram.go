package handlers

import (
	"context"
	"fmt"

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
		if update.FromChat() != nil {

			chatId := update.FromChat().ID

			fmt.Printf("api.usermapa[chatId]: %v\n", api.usermapa[chatId])
			user, _ := api.services.User.UserByChatID(context.Background(), int(chatId))
			msg := tgbotapi.NewMessage(chatId, "")

			if user == nil {
				if _, exi := api.usermapa[chatId]; !exi {
					api.Hello(update.Message, chatId)
					continue
				}
				api.profileUser(update, chatId, msg)
				continue
			}

			if _, exi := api.usermapa[chatId]; !exi {
				api.usermapa[chatId] = user
			}

			if api.usermapa[chatId].Aim == 1 {

				api.jobSeekersHandler(update, msg, chatId)
				continue
			}
		}
	}

	return nil
}

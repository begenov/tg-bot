package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAPI struct {
	bot     *tgbotapi.BotAPI
	Message *TelegramMessage
}

func NewTelegramAPI(token string) (*TelegramAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		return nil, err
	}

	bot.Debug = true
	return &TelegramAPI{bot: bot,
		Message: NewTelegramMessage(bot)}, nil
}

func (api *TelegramAPI) StartTelegramAPI() error {

	log.Printf("Authorized on account %s", api.bot.Self.UserName)

	updates := api.initUpdatesChannell()

	api.handleUpdates(updates)

	return nil
}

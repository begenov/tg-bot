package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/*
type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBotAPI(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdatesChannell()

	b.handleUpdates(updates)

	return nil
}
*/

type TelegramAPI struct {
	bot     *tgbotapi.BotAPI
	Message *TelegramMessage
}

func NewTelegramAPI(token string) (*TelegramAPI, error) {
	bot, err := tgbotapi.NewBotAPI("")

	if err != nil {
		return nil, err
	}

	bot.Debug = true
	return &TelegramAPI{bot: bot,
		Message: NewTelegramMessage(bot)}, nil
}

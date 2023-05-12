package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TelegramMessage struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramMessage(bot *tgbotapi.BotAPI) *TelegramMessage {
	return &TelegramMessage{bot: bot}
}

func (m *TelegramMessage) Send() error {
	return nil
}

func (m *TelegramMessage) SendWithOptions(options *tgbotapi.MessageConfig) error {
	_, err := m.bot.Send(options)
	if err != nil {
		return err
	}
	return nil
}

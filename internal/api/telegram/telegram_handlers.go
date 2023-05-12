package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const commandStart = "start"

func (api *TelegramAPI) handlerCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды :(")
	switch message.Command() {
	case commandStart:
		msg.Text = "Ты ввел команду /start"
		_, err := api.bot.Send(msg)
		return err
	default:
		_, err := api.bot.Send(msg)
		return err
	}
}

func (api *TelegramAPI) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil { // If we got a message

			if update.Message.IsCommand() {
				api.handlerCommand(update.Message)
				continue
			}

			api.handleMessage(update.Message)
		}
	}
}

func (api *TelegramAPI) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	api.bot.Send(msg)
}

func (api *TelegramAPI) initUpdatesChannell() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := api.bot.GetUpdatesChan(u)
	return updates

}

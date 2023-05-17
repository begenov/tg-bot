package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/begenov/tg-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (api *TelegramAPI) jobSeekersHandler(update tgbotapi.Update, user models.User, msg tgbotapi.MessageConfig) {
	switch user.Stage {
	case 0:

		// api.usermapa[chatId].Stage = 8
		msg.Text = "В какой сфере вы бы хотели найти работу?"
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Торговля", "0"),
				tgbotapi.NewInlineKeyboardButtonData("Общепит", "1"),
				tgbotapi.NewInlineKeyboardButtonData("Другое", "2"),
				tgbotapi.NewInlineKeyboardButtonData("Пропустить шаг", "3"),
			),
		)

		msg.ReplyMarkup = inlineKeyboard
		api.bot.Send(msg)

	case 1:
		if update.CallbackQuery != nil {
			api.workFieldHandler(update, int64(user.ChatID), msg)
		}
	default:
		log.Println("-----------------------")
	}

	// api.workFieldHandler(update, chatId, msg)
}

func (api *TelegramAPI) workFieldHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	workField, err := strconv.Atoi(update.CallbackQuery.Data)
	if err != nil {
		fmt.Println("WorkField handler error --------")
		log.Fatal(err)
	}
	api.usermapa[chatId].Stage = 2
	log.Println(workField, "------------------------------------------")
	msg.Text = "Дальше будеь выбор профессиии, пока в разработке"
	api.bot.Send(msg)
}

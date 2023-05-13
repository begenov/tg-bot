package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/begenov/tg-bot/internal/config"
	"github.com/begenov/tg-bot/internal/models"
	"github.com/begenov/tg-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAPI struct {
	bot      *tgbotapi.BotAPI
	services *services.Service
}

func NewTelegramAPI(bot *tgbotapi.BotAPI, servces *services.Service) *TelegramAPI {

	return &TelegramAPI{
		bot:      bot,
		services: servces,
	}
}

var (
	name string
)

/*
func (api *TelegramAPI) StartTelegramAPI() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := api.bot.GetUpdatesChan(u)
	for update := range updates {
		fmt.Println(name)
		if update.CallbackQuery != nil {
			if err := api.handleCallbackQuery(update.CallbackQuery); err != nil {
				log.Fatal("error")
			}
			continue
		}
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			// check in register
			if err := api.handleCommand(update.Message); err != nil {
				api.errorHandler()
			}
			continue
		}
		if err := api.hadnleMessage(update.Message); err != nil {
			api.errorHandler()
		}
	}
	return nil
}
*/

func (api *TelegramAPI) StartTelegramAPI(cfg *config.Config) error {
	fmt.Println("ooo")

	fmt.Println("botToken")
	offset := 0
	for {
		updates, err := getUpdates(cfg.TelegramAPI.BaseURL, offset)
		if err != nil {
			log.Println("error", err)
		}

		fmt.Println(updates)
		for _, update := range updates {
			respond(update, cfg.TelegramAPI.BaseURL)
			offset = update.UpdateID + 1
		}

	}

}

func getUpdates(botURl string, offset int) ([]models.Update, error) {
	resp, err := http.Get(botURl + "getUpdates" + "?offset=" + strconv.Itoa(offset))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var restResponse models.RestResponse

	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

func respond(update models.Update, botURL string) error {
	var botMessage models.BotMessage

	botMessage.ChatID = update.Message.Chat.ChatID
	botMessage.Text = update.Message.Text

	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}

	_, err = http.Post(botURL+"sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	return nil
}

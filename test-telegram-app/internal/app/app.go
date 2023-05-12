package app

import (
	"github.com/begenov/tg-bot/db"
	"github.com/begenov/tg-bot/test-telegram-app/internal/config"
	"github.com/begenov/tg-bot/test-telegram-app/internal/handlers"
	"github.com/begenov/tg-bot/test-telegram-app/internal/repository"
	"github.com/begenov/tg-bot/test-telegram-app/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run(cfg *config.Config) error {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramAPI.Token)

	if err != nil {
		return err
	}

	bot.Debug = true

	db, err := db.NewDB(cfg.DB.Driver, cfg.DB.DSN)
	if err != nil {
		return err
	}

	repository := repository.NewRepository(db)
	services := services.NewService(repository)

	telegramBot := handlers.NewTelegramAPI(bot, services)

	return nil
}

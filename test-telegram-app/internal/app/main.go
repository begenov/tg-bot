package app

import (
	"fmt"

	"github.com/begenov/tg-bot/db"
	"github.com/begenov/tg-bot/test-telegram-app/internal/config"
	"github.com/begenov/tg-bot/test-telegram-app/internal/handlers"
	"github.com/begenov/tg-bot/test-telegram-app/internal/repository"
	"github.com/begenov/tg-bot/test-telegram-app/internal/services"
)

func Run(cfg *config.Config) error {
	db, err := db.NewDB(cfg.DB.Driver, cfg.DB.DSN)
	if err != nil {
		return err
	}

	repository := repository.NewRepository(db)
	services, err := services.NewService(cfg.TelegramAPI.Token, repository)

	if err != nil {
		return err
	}

	handlers := handlers.NewHandler(services)

	fmt.Println(handlers)

	return nil
}

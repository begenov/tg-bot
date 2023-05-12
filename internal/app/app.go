package app

import (
	"fmt"

	"github.com/begenov/tg-bot/db"
	"github.com/begenov/tg-bot/internal/api/telegram"
	"github.com/begenov/tg-bot/internal/config"
	"github.com/begenov/tg-bot/internal/handler"
	"github.com/begenov/tg-bot/internal/repository"
	"github.com/begenov/tg-bot/internal/service"
)

type Application struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) *Application {
	return &Application{
		cfg: cfg,
	}
}

func (app *Application) Run() error {
	telegramAPI, err := telegram.NewTelegramAPI(app.cfg.Bot.Token)
	if err != nil {
		return err
	}

	db, err := db.NewDB(app.cfg.SQLite.Driver, app.cfg.SQLite.DSN)
	if err != nil {
		return err
	}

	repository := repository.NewRepository(db)

	service := service.NewService(repository)

	handler := handler.NewHandler(service, telegramAPI)

	fmt.Println(handler)

	return nil
}

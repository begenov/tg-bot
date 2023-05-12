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
	api *telegram.TelegramAPI
}

func NewApp(cfg *config.Config, api *telegram.TelegramAPI) *Application {
	return &Application{
		cfg: cfg,
		api: api,
	}
}

func (app *Application) Run() error {

	db, err := db.NewDB(app.cfg.SQLite.Driver, app.cfg.SQLite.DSN)
	if err != nil {
		return err
	}

	repository := repository.NewRepository(db)

	service := service.NewService(repository)

	handler := handler.NewHandler(service)
	fmt.Println(handler)
	return nil
}

package handler

import (
	"github.com/begenov/tg-bot/internal/api/telegram"
	"github.com/begenov/tg-bot/internal/service"
)

type Handler struct {
	service     *service.Service
	telegramAPI *telegram.TelegramAPI
}

func NewHandler(service *service.Service, API *telegram.TelegramAPI) *Handler {

	return &Handler{
		service:     service,
		telegramAPI: API,
	}
}

func (h *Handler) Start() error {
	return nil
}

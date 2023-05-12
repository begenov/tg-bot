package services

import "github.com/begenov/tg-bot/test-telegram-app/internal/repository"

type Service struct {
	TelegramAPI *TelegramAPI
	repository  *repository.Repository
}

func NewService(token string, repos *repository.Repository) (*Service, error) {
	telegramAPI, err := NewTelegramAPI(token)
	if err != nil {
		return nil, err
	}

	return &Service{
		TelegramAPI: telegramAPI,
		repository:  repos,
	}, err
}

package services

import (
	"context"

	"github.com/begenov/tg-bot/internal/models"
	"github.com/begenov/tg-bot/internal/repository"
)

type UserServiceIR interface {
	Create(ctx context.Context, user models.User) error
	UserByChatID(ctx context.Context, chatID int) (*models.User, error)
}

type Service struct {
	User *UserProfileService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserProfileService(&repos.User),
	}
}

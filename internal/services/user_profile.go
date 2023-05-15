package services

import (
	"context"

	"github.com/begenov/tg-bot/internal/models"
	"github.com/begenov/tg-bot/internal/repository"
)

type UserProfileService struct {
	user repository.UserIR
}

func NewUserProfileService(user repository.UserIR) *UserProfileService {
	return &UserProfileService{
		user: user,
	}
}

func (u *UserProfileService) Create(ctx context.Context, user models.User) error {
	return u.user.Create(ctx, user)
}

func (u *UserProfileService) UserByChatID(ctx context.Context, chatID int) (*models.User, error) {
	return u.user.UserByChatID(ctx, chatID)
}

package services

import (
	"context"
	"regexp"

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

func ValidatePhoneNumber(phoneNumber string) bool {
	// Паттерн для проверки формата номера телефона
	pattern := `^\+7\d{10}$`

	// Создание регулярного выражения
	regex := regexp.MustCompile(pattern)

	// Проверка соответствия паттерну
	return regex.MatchString(phoneNumber)
}

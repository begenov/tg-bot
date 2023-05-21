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

type JobSeeker interface {
	CreateJobSeeker(ctx context.Context, job models.JobSeeker) error
	JobSeekerByChatID(ctx context.Context, chatID int) (*models.JobSeeker, error)
}

type JobPosting interface {
	CreateJobPosting(ctx context.Context, v models.Vacancy) error
}

type Service struct {
	User       UserServiceIR
	JobSeeker  JobSeeker
	JobPosting JobPosting
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:       NewUserProfileService(repos.User),
		JobSeeker:  NewJobSeekerService(repos.JobSeeker),
		JobPosting: NewJobPostingService(repos.JobPosting),
	}
}

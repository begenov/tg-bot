package services

import (
	"context"

	"github.com/begenov/tg-bot/internal/models"
	"github.com/begenov/tg-bot/internal/repository"
)

type JobSeekerService struct {
	repo repository.JobSeeker
}

func NewJobSeekerService(repo repository.JobSeeker) *JobSeekerService {
	return &JobSeekerService{
		repo: repo,
	}
}

func (services *JobSeekerService) CreateJobSeeker(ctx context.Context, job models.JobSeeker) *JobSeekerService {
	return nil
}

func (services *JobSeekerService) JobSeekerByChatID(ctx context.Context, chatID int) (*models.JobSeeker, error) {
	return nil, nil
}

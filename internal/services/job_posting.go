package services

import (
	"context"

	"github.com/begenov/tg-bot/internal/models"
	"github.com/begenov/tg-bot/internal/repository"
)

type JobPostingService struct {
	repo repository.JobPosting
}

func NewJobPostingService(repo repository.JobPosting) *JobPostingService {
	return &JobPostingService{
		repo: repo,
	}
}

func (s *JobPostingService) CreateJobPosting(ctx context.Context, v models.Vacancy) error {
	return s.repo.CreateJobPosting(ctx, v)
}

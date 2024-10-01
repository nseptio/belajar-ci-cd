package services

import (
	"context"
	"errors"
	"time"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/repositories"
)

type SurveyService interface {
	CreateSurvey(ctx context.Context, survey *models.Survey) (*models.Survey, error)
	GetAllSurveys(ctx context.Context) ([]*models.Survey, error)
	GetSurveyByID(ctx context.Context, id string) (*models.Survey, error)
	UpdateSurvey(ctx context.Context, id string, survey *models.Survey) (*models.Survey, error)
	DeleteSurvey(ctx context.Context, id string) error
}

type surveyService struct {
	repository repositories.SurveyRepository
}

func NewSurveyService(repo repositories.SurveyRepository) SurveyService {
	return &surveyService{
		repository: repo,
	}
}

func (s *surveyService) CreateSurvey(ctx context.Context, survey *models.Survey) (*models.Survey, error) {
	// Ensure start date is before end date
	if survey.StartDate.After(survey.EndDate) {
		return nil, errors.New("start date cannot be after end date")
	}

	// Default value for publication status if not provided
	if !survey.IsPublished {
		survey.IsPublished = false
	}

	// Check if the year is valid (cannot be in the future)
	if survey.Year > time.Now().Year() {
		return nil, errors.New("survey year cannot be in the future")
	}

	return s.repository.CreateSurvey(ctx, survey)
}

func (s *surveyService) GetAllSurveys(ctx context.Context) ([]*models.Survey, error) {
	return s.repository.GetAllSurveys(ctx)
}

func (s *surveyService) GetSurveyByID(ctx context.Context, id string) (*models.Survey, error) {
	survey, err := s.repository.GetSurveyByID(ctx, id)
	if err != nil {
		return nil, errors.New("survey not found")
	}
	return survey, nil
}

func (s *surveyService) UpdateSurvey(ctx context.Context, id string, survey *models.Survey) (*models.Survey, error) {
	// Ensure the survey exists
	existingSurvey, err := s.repository.GetSurveyByID(ctx, id)
	if err != nil {
		return nil, errors.New("survey not found")
	}

	// Ensure start date is before end date
	if survey.StartDate.After(survey.EndDate) {
		return nil, errors.New("start date cannot be after end date")
	}

	// Keep original publication status if not updated
	if !survey.IsPublished {
		survey.IsPublished = existingSurvey.IsPublished
	}

	return s.repository.UpdateSurvey(ctx, id, survey)
}

func (s *surveyService) DeleteSurvey(ctx context.Context, id string) error {
	err := s.repository.DeleteSurvey(ctx, id)
	if err != nil {
		return errors.New("survey not found")
	}
	return nil
}

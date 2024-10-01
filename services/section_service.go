package services

import (
	"context"
	"errors"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/repositories"
)

type SectionService interface {
	CreateSection(ctx context.Context, section *models.Section) (*models.Section, error)
	GetAllSections(ctx context.Context) ([]*models.Section, error)
	GetSectionByID(ctx context.Context, id string) (*models.Section, error)
	UpdateSection(ctx context.Context, id string, section *models.Section) (*models.Section, error)
	DeleteSection(ctx context.Context, id string) error
}

type sectionService struct {
	repository repositories.SectionRepository
	surveyRepo repositories.SurveyRepository
}

func NewSectionService(repo repositories.SectionRepository, surveyRepo repositories.SurveyRepository) SectionService {
	return &sectionService{
		repository: repo,
		surveyRepo: surveyRepo,
	}
}

func (s *sectionService) CreateSection(ctx context.Context, section *models.Section) (*models.Section, error) {
	_, err := s.surveyRepo.GetSurveyByID(ctx, section.SurveyID.Hex())
	if err != nil {
		return nil, errors.New("invalid SurveyID")
	}
	return s.repository.CreateSection(ctx, section)
}

func (s *sectionService) GetAllSections(ctx context.Context) ([]*models.Section, error) {
	return s.repository.GetAllSections(ctx)
}

func (s *sectionService) GetSectionByID(ctx context.Context, id string) (*models.Section, error) {
	return s.repository.GetSectionByID(ctx, id)
}

func (s *sectionService) UpdateSection(ctx context.Context, id string, section *models.Section) (*models.Section, error) {
	_, err := s.surveyRepo.GetSurveyByID(ctx, section.SurveyID.Hex())
	if err != nil {
		return nil, errors.New("invalid SurveyID")
	}
	return s.repository.UpdateSection(ctx, id, section)
}

func (s *sectionService) DeleteSection(ctx context.Context, id string) error {
	return s.repository.DeleteSection(ctx, id)
}

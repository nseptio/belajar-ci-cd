package services

import (
	"context"
	"errors"
	"testing"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mock for SectionRepository
type MockSectionRepository struct {
	mock.Mock
}

func (m *MockSectionRepository) CreateSection(ctx context.Context, section *models.Section) (*models.Section, error) {
	args := m.Called(ctx, section)
	return args.Get(0).(*models.Section), args.Error(1)
}

func (m *MockSectionRepository) GetAllSections(ctx context.Context) ([]*models.Section, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Section), args.Error(1)
}

func (m *MockSectionRepository) GetSectionByID(ctx context.Context, id string) (*models.Section, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Section), args.Error(1)
}

func (m *MockSectionRepository) UpdateSection(ctx context.Context, id string, section *models.Section) (*models.Section, error) {
	args := m.Called(ctx, id, section)
	return args.Get(0).(*models.Section), args.Error(1)
}

func (m *MockSectionRepository) DeleteSection(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Mock for SurveyRepository
type MockSurveyRepository struct {
	mock.Mock
}

func (m *MockSurveyRepository) GetByID(id string) (*models.Survey, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Survey), args.Error(1)
}

func (m *MockSurveyRepository) Create(survey *models.Survey) (*mongo.InsertOneResult, error) {
	args := m.Called(survey)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockSurveyRepository) GetAll() ([]models.Survey, error) {
	args := m.Called()
	return args.Get(0).([]models.Survey), args.Error(1)
}

func (m *MockSurveyRepository) Update(id string, survey *models.Survey) error {
	args := m.Called(id, survey)
	return args.Error(0)
}

func (m *MockSurveyRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateSection_Success(t *testing.T) {
	mockSectionRepo := new(MockSectionRepository)
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSectionService(mockSectionRepo, mockSurveyRepo)

	section := &models.Section{
		Title:       "Section 1",
		Description: "Description of Section 1",
		SurveyID:    primitive.NewObjectID(),
	}

	mockSurveyRepo.On("GetSurveyByID", mock.Anything, section.SurveyID.Hex()).Return(&models.Survey{}, nil)
	mockSectionRepo.On("CreateSection", mock.Anything, section).Return(section, nil)

	createdSection, err := service.CreateSection(context.TODO(), section)

	assert.NoError(t, err)
	assert.Equal(t, section, createdSection)
}

func TestCreateSection_InvalidSurveyID(t *testing.T) {
	mockSectionRepo := new(MockSectionRepository)
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSectionService(mockSectionRepo, mockSurveyRepo)

	section := &models.Section{
		Title:       "Section 1",
		Description: "Description of Section 1",
		SurveyID:    primitive.NewObjectID(),
	}

	mockSurveyRepo.On("GetSurveyByID", mock.Anything, section.SurveyID.Hex()).Return((*models.Survey)(nil), errors.New("invalid SurveyID"))

	createdSection, err := service.CreateSection(context.TODO(), section)

	assert.Error(t, err)
	assert.Nil(t, createdSection)
}

func TestGetAllSections(t *testing.T) {
	mockSectionRepo := new(MockSectionRepository)
	service := NewSectionService(mockSectionRepo, nil)

	sections := []*models.Section{
		{Title: "Section 1"},
		{Title: "Section 2"},
	}

	mockSectionRepo.On("GetAllSections", mock.Anything).Return(sections, nil)

	result, err := service.GetAllSections(context.TODO())

	assert.NoError(t, err)
	assert.Equal(t, sections, result)
}

func TestGetAllSections_Error(t *testing.T) {
	mockSectionRepo := new(MockSectionRepository)
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSectionService(mockSectionRepo, mockSurveyRepo)

	mockSectionRepo.On("GetAllSections", mock.Anything).Return(([]*models.Section)(nil), errors.New("some error"))

	sections, err := service.GetAllSections(context.TODO())

	assert.Error(t, err)
	assert.Nil(t, sections)
}

func TestGetSectionByID_Success(t *testing.T) {
	mockSectionRepo := new(MockSectionRepository)
	service := NewSectionService(mockSectionRepo, nil)

	sectionID := primitive.NewObjectID().Hex()
	section := &models.Section{ID: primitive.NewObjectID(), Title: "Section 1"}

	mockSectionRepo.On("GetSectionByID", mock.Anything, sectionID).Return(section, nil)

	result, err := service.GetSectionByID(context.TODO(), sectionID)

	assert.NoError(t, err)
	assert.Equal(t, section, result)
}

func TestGetSectionByID_InvalidID(t *testing.T) {
	mockSectionRepo := new(MockSectionRepository)
	service := NewSectionService(mockSectionRepo, nil)

	sectionID := primitive.NewObjectID().Hex()

	mockSectionRepo.On("GetSectionByID", mock.Anything, sectionID).Return((*models.Section)(nil), errors.New("section not found"))

	result, err := service.GetSectionByID(context.TODO(), sectionID)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateSection_Success(t *testing.T) {
	mockSectionRepo := new(MockSectionRepository)
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSectionService(mockSectionRepo, mockSurveyRepo)

	sectionID := primitive.NewObjectID().Hex()
	section := &models.Section{ID: primitive.NewObjectID(), Title: "Updated Section", SurveyID: primitive.NewObjectID()}

	mockSurveyRepo.On("GetSurveyByID", mock.Anything, section.SurveyID.Hex()).Return(&models.Survey{}, nil)
	mockSectionRepo.On("UpdateSection", mock.Anything, sectionID, section).Return(section, nil)

	updatedSection, err := service.UpdateSection(context.TODO(), sectionID, section)

	assert.NoError(t, err)
	assert.Equal(t, section, updatedSection)
}

func TestUpdateSection_InvalidSurveyID(t *testing.T) {
	mockSectionRepo := new(MockSectionRepository)
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSectionService(mockSectionRepo, mockSurveyRepo)

	sectionID := primitive.NewObjectID().Hex()
	section := &models.Section{ID: primitive.NewObjectID(), Title: "Updated Section", SurveyID: primitive.NewObjectID()}

	mockSurveyRepo.On("GetSurveyByID", mock.Anything, section.SurveyID.Hex()).Return((*models.Survey)(nil), errors.New("invalid SurveyID"))

	updatedSection, err := service.UpdateSection(context.TODO(), sectionID, section)

	assert.Error(t, err)
	assert.Nil(t, updatedSection)
}

func TestDeleteSection_Success(t *testing.T) {
	mockSectionRepo := new(MockSectionRepository)
	service := NewSectionService(mockSectionRepo, nil)

	sectionID := primitive.NewObjectID().Hex()

	mockSectionRepo.On("DeleteSection", mock.Anything, sectionID).Return(nil)

	err := service.DeleteSection(context.TODO(), sectionID)

	assert.NoError(t, err)
}

func TestDeleteSection_InvalidID(t *testing.T) {
	mockSectionRepo := new(MockSectionRepository)
	service := NewSectionService(mockSectionRepo, nil)

	sectionID := primitive.NewObjectID().Hex()

	mockSectionRepo.On("DeleteSection", mock.Anything, sectionID).Return(errors.New("section not found"))

	err := service.DeleteSection(context.TODO(), sectionID)

	assert.Error(t, err)
}

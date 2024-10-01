package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *MockSurveyRepository) CreateSurvey(ctx context.Context, survey *models.Survey) (*models.Survey, error) {
	args := m.Called(ctx, survey)
	return args.Get(0).(*models.Survey), args.Error(1)
}

func (m *MockSurveyRepository) GetAllSurveys(ctx context.Context) ([]*models.Survey, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Survey), args.Error(1)
}

func (m *MockSurveyRepository) GetSurveyByID(ctx context.Context, id string) (*models.Survey, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Survey), args.Error(1)
}

func (m *MockSurveyRepository) UpdateSurvey(ctx context.Context, id string, survey *models.Survey) (*models.Survey, error) {
	args := m.Called(ctx, id, survey)
	return args.Get(0).(*models.Survey), args.Error(1)
}

func (m *MockSurveyRepository) DeleteSurvey(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Tests
func TestCreateSurvey_Success(t *testing.T) {
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSurveyService(mockSurveyRepo)

	survey := &models.Survey{
		Title:     "Survey 1",
		Year:      time.Now().Year(),
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 7),
	}

	mockSurveyRepo.On("CreateSurvey", mock.Anything, survey).Return(survey, nil)

	createdSurvey, err := service.CreateSurvey(context.TODO(), survey)

	assert.NoError(t, err)
	assert.Equal(t, survey, createdSurvey)
}

func TestCreateSurvey_InvalidStartDate(t *testing.T) {
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSurveyService(mockSurveyRepo)

	survey := &models.Survey{
		Title:     "Survey 1",
		Year:      time.Now().Year(),
		StartDate: time.Now().AddDate(0, 0, 7), // Start date is after end date
		EndDate:   time.Now(),
	}

	createdSurvey, err := service.CreateSurvey(context.TODO(), survey)

	assert.Error(t, err)
	assert.Nil(t, createdSurvey)
	assert.Equal(t, "start date cannot be after end date", err.Error())
}

func TestCreateSurvey_InvalidYear(t *testing.T) {
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSurveyService(mockSurveyRepo)

	survey := &models.Survey{
		Title:     "Survey 1",
		Year:      time.Now().Year() + 1, // Future year
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 7),
	}

	createdSurvey, err := service.CreateSurvey(context.TODO(), survey)

	assert.Error(t, err)
	assert.Nil(t, createdSurvey)
	assert.Equal(t, "survey year cannot be in the future", err.Error())
}

func TestGetAllSurveys_Success(t *testing.T) {
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSurveyService(mockSurveyRepo)

	surveys := []*models.Survey{
		{Title: "Survey 1"},
		{Title: "Survey 2"},
	}

	mockSurveyRepo.On("GetAllSurveys", mock.Anything).Return(surveys, nil)

	result, err := service.GetAllSurveys(context.TODO())

	assert.NoError(t, err)
	assert.Equal(t, surveys, result)
}

func TestGetAllSurveys_Error(t *testing.T) {
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSurveyService(mockSurveyRepo)

	mockSurveyRepo.On("GetAllSurveys", mock.Anything).Return(([]*models.Survey)(nil), errors.New("some error"))

	surveys, err := service.GetAllSurveys(context.TODO())

	assert.Error(t, err)
	assert.Nil(t, surveys)
}

func TestGetSurveyByID_Success(t *testing.T) {
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSurveyService(mockSurveyRepo)

	surveyID := primitive.NewObjectID().Hex()
	survey := &models.Survey{ID: primitive.NewObjectID(), Title: "Survey 1"}

	mockSurveyRepo.On("GetSurveyByID", mock.Anything, surveyID).Return(survey, nil)

	result, err := service.GetSurveyByID(context.TODO(), surveyID)

	assert.NoError(t, err)
	assert.Equal(t, survey, result)
}

func TestGetSurveyByID_NotFound(t *testing.T) {
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSurveyService(mockSurveyRepo)

	surveyID := primitive.NewObjectID().Hex()

	mockSurveyRepo.On("GetSurveyByID", mock.Anything, surveyID).Return((*models.Survey)(nil), errors.New("survey not found"))

	result, err := service.GetSurveyByID(context.TODO(), surveyID)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "survey not found", err.Error())
}

func TestUpdateSurvey_Success(t *testing.T) {
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSurveyService(mockSurveyRepo)

	surveyID := primitive.NewObjectID().Hex()
	survey := &models.Survey{ID: primitive.NewObjectID(), Title: "Updated Survey"}

	mockSurveyRepo.On("GetSurveyByID", mock.Anything, surveyID).Return(survey, nil)
	mockSurveyRepo.On("UpdateSurvey", mock.Anything, surveyID, survey).Return(survey, nil)

	updatedSurvey, err := service.UpdateSurvey(context.TODO(), surveyID, survey)

	assert.NoError(t, err)
	assert.Equal(t, survey, updatedSurvey)
}

func TestUpdateSurvey_NotFound(t *testing.T) {
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSurveyService(mockSurveyRepo)

	surveyID := primitive.NewObjectID().Hex()
	survey := &models.Survey{ID: primitive.NewObjectID(), Title: "Updated Survey"}

	mockSurveyRepo.On("GetSurveyByID", mock.Anything, surveyID).Return((*models.Survey)(nil), errors.New("survey not found"))

	updatedSurvey, err := service.UpdateSurvey(context.TODO(), surveyID, survey)

	assert.Error(t, err)
	assert.Nil(t, updatedSurvey)
	assert.Equal(t, "survey not found", err.Error())
}

func TestUpdateSurvey_StartDateAfterEndDate(t *testing.T) {
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSurveyService(mockSurveyRepo)

	surveyID := primitive.NewObjectID().Hex()
	survey := &models.Survey{
		ID:        primitive.NewObjectID(),
		Title:     "Invalid Survey",
		StartDate: time.Now().AddDate(0, 0, 1), // Start date is tomorrow
		EndDate:   time.Now(),                  // End date is today
	}

	mockSurveyRepo.On("GetSurveyByID", mock.Anything, surveyID).Return(survey, nil)

	updatedSurvey, err := service.UpdateSurvey(context.TODO(), surveyID, survey)

	assert.Error(t, err)
	assert.Nil(t, updatedSurvey)
	assert.Equal(t, "start date cannot be after end date", err.Error())
}

func TestDeleteSurvey_Success(t *testing.T) {
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSurveyService(mockSurveyRepo)

	surveyID := primitive.NewObjectID().Hex()

	mockSurveyRepo.On("DeleteSurvey", mock.Anything, surveyID).Return(nil)

	err := service.DeleteSurvey(context.TODO(), surveyID)

	assert.NoError(t, err)
}

func TestDeleteSurvey_NotFound(t *testing.T) {
	mockSurveyRepo := new(MockSurveyRepository)
	service := NewSurveyService(mockSurveyRepo)

	surveyID := primitive.NewObjectID().Hex()

	mockSurveyRepo.On("DeleteSurvey", mock.Anything, surveyID).Return(errors.New("survey not found"))

	err := service.DeleteSurvey(context.TODO(), surveyID)

	assert.Error(t, err)
	assert.Equal(t, "survey not found", err.Error())
}

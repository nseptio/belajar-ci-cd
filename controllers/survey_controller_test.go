package controllers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockSurveyService defines a mock service for surveys
type MockSurveyService struct {
	mock.Mock
}

func (m *MockSurveyService) CreateSurvey(ctx context.Context, survey *models.Survey) (*models.Survey, error) {
	args := m.Called(ctx, survey)
	return args.Get(0).(*models.Survey), args.Error(1)
}

func (m *MockSurveyService) GetAllSurveys(ctx context.Context) ([]*models.Survey, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Survey), args.Error(1)
}

func (m *MockSurveyService) GetSurveyByID(ctx context.Context, id string) (*models.Survey, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Survey), args.Error(1)
}

func (m *MockSurveyService) UpdateSurvey(ctx context.Context, id string, survey *models.Survey) (*models.Survey, error) {
	args := m.Called(ctx, id, survey)
	return args.Get(0).(*models.Survey), args.Error(1)
}

func (m *MockSurveyService) DeleteSurvey(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateSurvey_Success(t *testing.T) {
	e := echo.New()
	// Ensure the dates are in the correct format, e.g., ISO 8601
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"title":"Survey 1", "start_date":"2024-09-01T00:00:00Z", "end_date":"2024-09-30T23:59:59Z"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	// Mock the CreateSurvey service call
	mockService.On("CreateSurvey", mock.Anything, mock.AnythingOfType("*models.Survey")).Return(&models.Survey{Title: "Survey 1"}, nil)

	// Run the handler and check the response
	if assert.NoError(t, controller.CreateSurvey(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "Survey 1")
	}
}

func TestCreateSurvey_BindError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`invalid json`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	if assert.NoError(t, controller.CreateSurvey(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Invalid input")
	}
}

func TestCreateSurvey_MissingTitle(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"title":"","start_date":"2024-09-01T00:00:00Z","end_date":"2024-09-30T23:59:59Z"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	if assert.NoError(t, controller.CreateSurvey(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Title, Start Date, and End Date are required")
	}
}

func TestCreateSurvey_ServiceError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"title":"Survey 1", "start_date":"2024-09-01T00:00:00Z", "end_date":"2024-09-30T23:59:59Z"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	mockService.On("CreateSurvey", mock.Anything, mock.AnythingOfType("*models.Survey")).Return((*models.Survey)(nil), errors.New("service error"))

	if assert.NoError(t, controller.CreateSurvey(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "Could not create survey: service error")
	}
}

func TestGetAllSurveys_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	mockService.On("GetAllSurveys", mock.Anything).Return([]*models.Survey{{Title: "Survey 1"}}, nil)

	if assert.NoError(t, controller.GetAllSurveys(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Survey 1")
	}
}

func TestGetAllSurveys_Error(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	mockService.On("GetAllSurveys", mock.Anything).Return(([]*models.Survey)(nil), errors.New("service error"))

	if assert.NoError(t, controller.GetAllSurveys(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "Could not retrieve surveys: service error")
	}
}

func TestGetSurveyByID_Success(t *testing.T) {
	e := echo.New()
	surveyID := "validID"
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(surveyID)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	mockService.On("GetSurveyByID", mock.Anything, surveyID).Return(&models.Survey{Title: "Survey 1"}, nil)

	if assert.NoError(t, controller.GetSurveyByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Survey 1")
	}
}

func TestGetSurveyByID_NotFound(t *testing.T) {
	e := echo.New()
	surveyID := "invalidID"
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(surveyID)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	mockService.On("GetSurveyByID", mock.Anything, surveyID).Return((*models.Survey)(nil), errors.New("survey not found"))

	if assert.NoError(t, controller.GetSurveyByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Contains(t, rec.Body.String(), "Survey not found")
	}
}

func TestUpdateSurvey_Success(t *testing.T) {
	e := echo.New()
	surveyID := "validID"
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"title":"Updated Survey", "start_date":"2024-09-01T00:00:00Z", "end_date":"2024-09-30T23:59:59Z"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(surveyID)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	mockService.On("UpdateSurvey", mock.Anything, surveyID, mock.AnythingOfType("*models.Survey")).Return(&models.Survey{Title: "Updated Survey"}, nil)

	if assert.NoError(t, controller.UpdateSurvey(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Updated Survey")
	}
}

func TestUpdateSurvey_BindError(t *testing.T) {
	e := echo.New()
	surveyID := "validID"
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`invalid json`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(surveyID)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	if assert.NoError(t, controller.UpdateSurvey(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Invalid input")
	}
}

func TestUpdateSurvey_MissingTitle(t *testing.T) {
	e := echo.New()
	surveyID := "validID"
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"title":"", "start_date":"2024-09-01T00:00:00Z", "end_date":"2024-09-30T23:59:59Z"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(surveyID)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	if assert.NoError(t, controller.UpdateSurvey(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Title, Start Date, and End Date are required")
	}
}

func TestUpdateSurvey_NotFound(t *testing.T) {
	e := echo.New()
	surveyID := "validID"
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"title":"Updated Survey", "start_date":"2024-09-01T00:00:00Z", "end_date":"2024-09-30T23:59:59Z"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(surveyID)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	// Simulate a service error
	mockService.On("UpdateSurvey", mock.Anything, surveyID, mock.AnythingOfType("*models.Survey")).Return((*models.Survey)(nil), errors.New("survey not found"))

	if assert.NoError(t, controller.UpdateSurvey(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Contains(t, rec.Body.String(), "Could not update survey: survey not found")
	}
}

func TestDeleteSurvey_Success(t *testing.T) {
	e := echo.New()
	surveyID := "validID"
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(surveyID)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	mockService.On("DeleteSurvey", mock.Anything, surveyID).Return(nil)

	if assert.NoError(t, controller.DeleteSurvey(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Empty(t, rec.Body.String())
	}
}

func TestDeleteSurvey_NotFound(t *testing.T) {
	e := echo.New()
	surveyID := "invalidID"
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(surveyID)

	mockService := new(MockSurveyService)
	controller := NewSurveyController(mockService)

	mockService.On("DeleteSurvey", mock.Anything, surveyID).Return(errors.New("service error"))

	if assert.NoError(t, controller.DeleteSurvey(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Contains(t, rec.Body.String(), "Survey not found")
	}
}

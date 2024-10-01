package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockSurveyController is a mock implementation of SurveyControllerInterface
type MockSurveyController struct {
	mock.Mock
}

func (m *MockSurveyController) CreateSurvey(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockSurveyController) GetAllSurveys(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockSurveyController) GetSurveyByID(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockSurveyController) UpdateSurvey(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockSurveyController) DeleteSurvey(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func TestSurveyRoutes(t *testing.T) {
	e := echo.New()
	mockController := new(MockSurveyController)

	SurveyRoutes(e, mockController)

	tests := []struct {
		method       string
		target       string
		handler      func(c echo.Context) error
		expectedCode int
	}{
		{http.MethodPost, "/surveys", mockController.CreateSurvey, http.StatusOK},
		{http.MethodGet, "/surveys", mockController.GetAllSurveys, http.StatusOK},
		{http.MethodGet, "/surveys/1", mockController.GetSurveyByID, http.StatusOK},
		{http.MethodPut, "/surveys/1", mockController.UpdateSurvey, http.StatusOK},
		{http.MethodDelete, "/surveys/1", mockController.DeleteSurvey, http.StatusOK},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.target, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockController.On("CreateSurvey", c).Return(nil)
		mockController.On("GetAllSurveys", c).Return(nil)
		mockController.On("GetSurveyByID", c).Return(nil)
		mockController.On("UpdateSurvey", c).Return(nil)
		mockController.On("DeleteSurvey", c).Return(nil)

		if assert.NoError(t, tt.handler(c)) {
			assert.Equal(t, tt.expectedCode, rec.Code)
		}
	}
}

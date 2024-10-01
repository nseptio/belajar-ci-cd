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

// MockUserController defines a mock service for users
type MockUserController struct {
	mock.Mock
}

func (m *MockUserController) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	args := m.Called(ctx, student)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Student), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserController) GetAllStudents(ctx context.Context) ([]*models.Student, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Student), args.Error(1)
}

func (m *MockUserController) GetStudentByID(ctx context.Context, id string) (*models.Student, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Student), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserController) UpdateStudent(ctx context.Context, id string, student *models.Student) (*models.Student, error) {
	args := m.Called(ctx, id, student)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Student), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockUserController) DeleteStudent(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return args.Error(0)
}

func TestCreateStudent_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"John Doe", "email":"user1@example.com", "phone_number":"081234567890", "university_name":"University 1", "start_year":2021}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserController)
	controller := NewUserController(mockService)

	mockService.On("CreateStudent", mock.Anything, mock.AnythingOfType("*models.Student")).Return(&models.Student{Name: "John Doe"}, nil)

	err := controller.CreateStudent(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestCreateStudent_InvalidInput(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"John Doe", "email":"user1@example.com",
	"phone_number":"081234567890", "university_name":"University 1"}`))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserController)
	controller := NewUserController(mockService)

	err := controller.CreateStudent(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateStudent_ServiceError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"John Doe", "email":"user1@example.com",
	"phone_number":"081234567890", "university_name":"University 1", "start_year":2021}`))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserController)
	controller := NewUserController(mockService)

	// Return nil for the student object and an error
	mockService.On("CreateStudent", mock.Anything, mock.AnythingOfType("*models.Student")).Return(nil, errors.New("service error"))

	err := controller.CreateStudent(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestGetAllStudents_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserController)
	controller := NewUserController(mockService)

	mockService.On("GetAllStudents", mock.Anything).Return([]*models.Student{{Name: "John Doe"}}, nil)

	err := controller.GetAllStudents(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetAllStudents_Error(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserController)
	controller := NewUserController(mockService)

	mockService.On("GetAllStudents", mock.Anything).Return(([]*models.Student)(nil), errors.New("service error"))

	err := controller.GetAllStudents(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestGetStudentByID_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserController)
	controller := NewUserController(mockService)

	mockService.On("GetStudentByID", mock.Anything, mock.AnythingOfType("string")).Return(&models.Student{Name: "John Doe"}, nil)

	err := controller.GetStudentByID(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetStudentByID_NotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserController)
	controller := NewUserController(mockService)

	mockService.On("GetStudentByID", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("student not found"))

	err := controller.GetStudentByID(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestUpdateStudent_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"name":"John Doe", "email":"user1@example.com",
	"phone_number":"081234567890", "university_name":"University 1", "start_year":2021}`))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserController)
	controller := NewUserController(mockService)

	mockService.On("UpdateStudent", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("*models.Student")).Return(&models.Student{Name: "John Doe"}, nil)

	err := controller.UpdateStudent(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateStudent_ServiceError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"name":"John Doe", "email":"user1@example.com", "phone_number":"081234567890", "university_name":"University 1", "start_year":2021}`))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserController)
	controller := NewUserController(mockService)

	mockService.On("UpdateStudent", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("*models.Student")).Return(nil, errors.New("service error"))

	err := controller.UpdateStudent(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestDeleteStudent_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserController)
	controller := NewUserController(mockService)

	mockService.On("DeleteStudent", mock.Anything, mock.AnythingOfType("string")).Return(nil)

	err := controller.DeleteStudent(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestDeleteStudent_NotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserController)
	controller := NewUserController(mockService)

	mockService.On("DeleteStudent", mock.Anything, mock.AnythingOfType("string")).Return(errors.New("not found"))

	err := controller.DeleteStudent(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestNewUserController(t *testing.T) {
	mockService := new(MockUserController)
	controller := NewUserController(mockService)
	assert.NotNil(t, controller)
}

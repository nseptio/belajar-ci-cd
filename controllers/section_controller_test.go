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

type MockSectionService struct {
	mock.Mock
}

func (m *MockSectionService) CreateSection(ctx context.Context, section *models.Section) (*models.Section, error) {
	args := m.Called(ctx, section)
	return args.Get(0).(*models.Section), args.Error(1)
}

func (m *MockSectionService) GetAllSections(ctx context.Context) ([]*models.Section, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Section), args.Error(1)
}

func (m *MockSectionService) GetSectionByID(ctx context.Context, id string) (*models.Section, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Section), args.Error(1)
}

func (m *MockSectionService) UpdateSection(ctx context.Context, id string, section *models.Section) (*models.Section, error) {
	args := m.Called(ctx, id, section)
	return args.Get(0).(*models.Section), args.Error(1)
}

func (m *MockSectionService) DeleteSection(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateSection_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"title":"Section 1"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	mockService.On("CreateSection", mock.Anything, mock.AnythingOfType("*models.Section")).Return(&models.Section{Title: "Section 1"}, nil)

	if assert.NoError(t, controller.CreateSection(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "Section 1")
	}
}

func TestCreateSection_BindError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`invalid json`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	if assert.NoError(t, controller.CreateSection(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Invalid input")
	}
}

func TestCreateSection_MissingTitle(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"title":""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	if assert.NoError(t, controller.CreateSection(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Title is required")
	}
}

func TestCreateSection_ServiceError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"title":"Section 1"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	mockService.On("CreateSection", mock.Anything, mock.AnythingOfType("*models.Section")).Return((*models.Section)(nil), errors.New("service error"))

	if assert.NoError(t, controller.CreateSection(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "Could not create section: service error")
	}
}

func TestGetAllSections_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	mockService.On("GetAllSections", mock.Anything).Return([]*models.Section{{Title: "Section 1"}}, nil)

	if assert.NoError(t, controller.GetAllSections(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Section 1")
	}
}

func TestGetAllSections_Error(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	mockService.On("GetAllSections", mock.Anything).Return(([]*models.Section)(nil), errors.New("service error"))

	if assert.NoError(t, controller.GetAllSections(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "Could not retrieve sections: service error")
	}
}

func TestGetSectionByID_Success(t *testing.T) {
	e := echo.New()
	sectionID := "validID"
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(sectionID)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	mockService.On("GetSectionByID", mock.Anything, sectionID).Return(&models.Section{Title: "Section 1"}, nil)

	if assert.NoError(t, controller.GetSectionByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Section 1")
	}
}

func TestGetSectionByID_NotFound(t *testing.T) {
	e := echo.New()
	sectionID := "invalidID"
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(sectionID)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	mockService.On("GetSectionByID", mock.Anything, sectionID).Return((*models.Section)(nil), errors.New("section not found"))

	if assert.NoError(t, controller.GetSectionByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Contains(t, rec.Body.String(), "Section not found")
	}
}
func TestUpdateSection_Success(t *testing.T) {
	e := echo.New()
	sectionID := "validID"
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"title":"Updated Section"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(sectionID)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	mockService.On("UpdateSection", mock.Anything, sectionID, mock.AnythingOfType("*models.Section")).Return(&models.Section{Title: "Updated Section"}, nil)

	if assert.NoError(t, controller.UpdateSection(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Updated Section")
	}
}

func TestUpdateSection_BindError(t *testing.T) {
	e := echo.New()
	sectionID := "validID"
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`invalid json`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(sectionID)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	if assert.NoError(t, controller.UpdateSection(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Invalid input")
	}
}

func TestUpdateSection_MissingTitle(t *testing.T) {
	e := echo.New()
	sectionID := "validID"
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"title":""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(sectionID)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	if assert.NoError(t, controller.UpdateSection(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Title is required")
	}
}

func TestUpdateSection_InvalidID(t *testing.T) {
	e := echo.New()
	sectionID := "invalidID"
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"title":"Updated Section"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(sectionID)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	mockService.On("UpdateSection", mock.Anything, sectionID, mock.AnythingOfType("*models.Section")).Return((*models.Section)(nil), errors.New("invalid section ID"))

	if assert.NoError(t, controller.UpdateSection(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Contains(t, rec.Body.String(), "Could not update section: invalid section ID")
	}
}
func TestDeleteSection_Success(t *testing.T) {
	e := echo.New()
	sectionID := "validID"
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(sectionID)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	mockService.On("DeleteSection", mock.Anything, sectionID).Return(nil)

	if assert.NoError(t, controller.DeleteSection(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Contains(t, rec.Body.String(), "")
	}
}

func TestDeleteSection_NotFound(t *testing.T) {
	e := echo.New()
	sectionID := "invalidID"
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(sectionID)

	mockService := new(MockSectionService)
	controller := NewSectionController(mockService)

	mockService.On("DeleteSection", mock.Anything, sectionID).Return(errors.New("Section not found"))

	if assert.NoError(t, controller.DeleteSection(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Contains(t, rec.Body.String(), "Section not found")
	}
}

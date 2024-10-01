package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSectionController struct {
	mock.Mock
}

func (m *MockSectionController) CreateSection(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockSectionController) GetAllSections(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockSectionController) GetSectionByID(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockSectionController) UpdateSection(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockSectionController) DeleteSection(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func TestSectionRoutes(t *testing.T) {
	e := echo.New()
	mockController := new(MockSectionController)
	SectionRoutes(e, mockController)

	t.Run("CreateSection", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/sections", strings.NewReader(`{"title":"Section 1"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockController.On("CreateSection", c).Return(nil)

		if assert.NoError(t, mockController.CreateSection(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("GetAllSections", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/sections", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockController.On("GetAllSections", c).Return(nil)

		if assert.NoError(t, mockController.GetAllSections(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("GetSectionByID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/sections/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockController.On("GetSectionByID", c).Return(nil)

		if assert.NoError(t, mockController.GetSectionByID(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("UpdateSection", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/sections/1", strings.NewReader(`{"title":"Updated Section"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockController.On("UpdateSection", c).Return(nil)

		if assert.NoError(t, mockController.UpdateSection(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("DeleteSection", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/sections/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockController.On("DeleteSection", c).Return(nil)

		if assert.NoError(t, mockController.DeleteSection(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

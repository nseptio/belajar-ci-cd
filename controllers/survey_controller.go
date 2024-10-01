package controllers

import (
	"net/http"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/services"
	"github.com/labstack/echo/v4"
)

type SurveyControllerInterface interface {
	CreateSurvey(ctx echo.Context) error
	GetAllSurveys(ctx echo.Context) error
	GetSurveyByID(ctx echo.Context) error
	UpdateSurvey(ctx echo.Context) error
	DeleteSurvey(ctx echo.Context) error
}

type SurveyController struct {
	service services.SurveyService
}

func NewSurveyController(service services.SurveyService) *SurveyController {
	return &SurveyController{service: service}
}

// CreateSurvey handles POST /surveys
func (c *SurveyController) CreateSurvey(ctx echo.Context) error {
	var survey models.Survey
	if err := ctx.Bind(&survey); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Validation for required fields
	if survey.Title == "" || survey.StartDate.IsZero() || survey.EndDate.IsZero() {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Title, Start Date, and End Date are required"})
	}

	createdSurvey, err := c.service.CreateSurvey(ctx.Request().Context(), &survey)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create survey: " + err.Error()})
	}

	return ctx.JSON(http.StatusCreated, createdSurvey)
}

// GetAllSurveys handles GET /surveys
func (c *SurveyController) GetAllSurveys(ctx echo.Context) error {
	surveys, err := c.service.GetAllSurveys(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not retrieve surveys: " + err.Error()})
	}
	return ctx.JSON(http.StatusOK, surveys)
}

// GetSurveyByID handles GET /surveys/:id
func (c *SurveyController) GetSurveyByID(ctx echo.Context) error {
	id := ctx.Param("id")
	survey, err := c.service.GetSurveyByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": "Survey not found"})
	}
	return ctx.JSON(http.StatusOK, survey)
}

// UpdateSurvey handles PUT /surveys/:id
func (c *SurveyController) UpdateSurvey(ctx echo.Context) error {
	id := ctx.Param("id")
	var survey models.Survey
	if err := ctx.Bind(&survey); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Validation for required fields
	if survey.Title == "" || survey.StartDate.IsZero() || survey.EndDate.IsZero() {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Title, Start Date, and End Date are required"})
	}

	updatedSurvey, err := c.service.UpdateSurvey(ctx.Request().Context(), id, &survey)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": "Could not update survey: " + err.Error()})
	}
	return ctx.JSON(http.StatusOK, updatedSurvey)
}

// DeleteSurvey handles DELETE /surveys/:id
func (c *SurveyController) DeleteSurvey(ctx echo.Context) error {
	id := ctx.Param("id")
	err := c.service.DeleteSurvey(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": "Survey not found"})
	}
	return ctx.NoContent(http.StatusNoContent)
}

package controllers

import (
	"net/http"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/services"
	"github.com/labstack/echo/v4"
)

type SectionControllerInterface interface {
	CreateSection(ctx echo.Context) error
	GetAllSections(ctx echo.Context) error
	GetSectionByID(ctx echo.Context) error
	UpdateSection(ctx echo.Context) error
	DeleteSection(ctx echo.Context) error
}

type SectionController struct {
	service services.SectionService
}

func NewSectionController(service services.SectionService) SectionControllerInterface {
	return &SectionController{service: service}
}

func (c *SectionController) CreateSection(ctx echo.Context) error {
	var section models.Section
	if err := ctx.Bind(&section); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	if section.Title == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Title is required"})
	}

	createdSection, err := c.service.CreateSection(ctx.Request().Context(), &section)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create section: " + err.Error()})
	}

	return ctx.JSON(http.StatusCreated, createdSection)
}

func (c *SectionController) GetAllSections(ctx echo.Context) error {
	sections, err := c.service.GetAllSections(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not retrieve sections: " + err.Error()})
	}
	return ctx.JSON(http.StatusOK, sections)
}

func (c *SectionController) GetSectionByID(ctx echo.Context) error {
	id := ctx.Param("id")
	section, err := c.service.GetSectionByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": "Section not found"})
	}
	return ctx.JSON(http.StatusOK, section)
}

func (c *SectionController) UpdateSection(ctx echo.Context) error {
	id := ctx.Param("id")
	var section models.Section
	if err := ctx.Bind(&section); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	if section.Title == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Title is required"})
	}

	updatedSection, err := c.service.UpdateSection(ctx.Request().Context(), id, &section)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": "Could not update section: " + err.Error()})
	}
	return ctx.JSON(http.StatusOK, updatedSection)
}

func (c *SectionController) DeleteSection(ctx echo.Context) error {
	id := ctx.Param("id")
	err := c.service.DeleteSection(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": "Section not found"})
	}
	return ctx.NoContent(http.StatusNoContent)
}

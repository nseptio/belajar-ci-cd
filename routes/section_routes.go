package routes

import (
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/controllers"
	"github.com/labstack/echo/v4"
)

func SectionRoutes(e *echo.Echo, sectionController controllers.SectionControllerInterface) {
	sections := e.Group("/sections")

	sections.POST("", sectionController.CreateSection)
	sections.GET("", sectionController.GetAllSections)
	sections.GET("/:id", sectionController.GetSectionByID)
	sections.PUT("/:id", sectionController.UpdateSection)
	sections.DELETE("/:id", sectionController.DeleteSection)
}

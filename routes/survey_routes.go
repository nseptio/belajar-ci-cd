package routes

import (
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/controllers"
	"github.com/labstack/echo/v4"
)

func SurveyRoutes(e *echo.Echo, surveyController controllers.SurveyControllerInterface) {
	surveys := e.Group("/surveys")

	surveys.POST("", surveyController.CreateSurvey)
	surveys.GET("", surveyController.GetAllSurveys)
	surveys.GET("/:id", surveyController.GetSurveyByID)
	surveys.PUT("/:id", surveyController.UpdateSurvey)
	surveys.DELETE("/:id", surveyController.DeleteSurvey)
}

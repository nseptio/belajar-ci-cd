package app

import (
	"fmt"
	"net/http"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/controllers"
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/repositories"
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (a *App) loadRoutes() {
	app := echo.New()

	// general middleware
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
			return nil
		},
	}))
	app.Use(middleware.Recover())

	router := app.Group("/api/v1")
	router.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	userGroup := router.Group("/users")
	a.loadUserRoutes(userGroup)

	studentGroup := router.Group("/students")
	userRepository := repositories.NewUserRepository(a.mongoDB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	a.loadStudentRoutes(studentGroup, userController)

	app.Pre(middleware.RemoveTrailingSlash())
	a.router = app
}

func (a *App) loadUserRoutes(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Get all users")
	})
}

func (a *App) loadStudentRoutes(g *echo.Group, c *controllers.UserController) {

	g.POST("/", c.CreateStudent)
	g.GET("/", c.GetAllStudents)
	g.GET("/:id", c.GetStudentByID)
	g.PUT("/:id", c.UpdateStudent)
	g.DELETE("/:id", c.DeleteStudent)
}

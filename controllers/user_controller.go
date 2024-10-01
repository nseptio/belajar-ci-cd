package controllers

import (
	"errors"
	"net/http"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/services"

	"github.com/labstack/echo/v4"
)

type UserControllerInterface interface {
	CreateStudent(ctx echo.Context) error
	GetAllStudents(ctx echo.Context) error
	GetStudentByID(ctx echo.Context) error
	UpdateStudent(ctx echo.Context) error
	DeleteStudent(ctx echo.Context) error
}

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

// CreateStudent handles POST /students
func (c *UserController) CreateStudent(ctx echo.Context) error {
	var student models.Student
	if err := ctx.Bind(&student); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Validation for required fields
	if student.Name == "" || student.Email == "" || student.PhoneNumber == "" || student.UniversityName == "" || student.StartYear == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Name, Email, Phone Number, University Name, and Start Year are required"})
	}

	createdStudent, err := c.service.CreateStudent(ctx.Request().Context(), &student)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create student: " + err.Error()})
	}

	return ctx.JSON(http.StatusCreated, createdStudent)
}

// GetAllStudents handles GET /students
func (c *UserController) GetAllStudents(ctx echo.Context) error {
	students, err := c.service.GetAllStudents(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not retrieve students: " + err.Error()})
	}
	return ctx.JSON(http.StatusOK, students)
}

// GetStudentByID handles GET /students/:id
func (c *UserController) GetStudentByID(ctx echo.Context) error {
	id := ctx.Param("id")
	student, err := c.service.GetStudentByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": "Student not found"})
	}
	return ctx.JSON(http.StatusOK, student)
}

// UpdateStudent handles PUT /students/:id
func (c *UserController) UpdateStudent(ctx echo.Context) error {
	id := ctx.Param("id")
	var student models.Student
	if err := ctx.Bind(&student); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	updatedStudent, err := c.service.UpdateStudent(ctx.Request().Context(), id, &student)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not update student: " + err.Error()})
	}

	return ctx.JSON(http.StatusOK, updatedStudent)
}

// DeleteStudent handles DELETE /students/:id
func (c *UserController) DeleteStudent(ctx echo.Context) error {
	id := ctx.Param("id")
	err := c.service.DeleteStudent(ctx.Request().Context(), id)

	if err != nil {
		if errors.Is(err, services.ErrStudentNotFound) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"message": "Student not found"})
		}

		return ctx.JSON(http.StatusNotFound, map[string]string{"message": "Could not delete student: " + err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}

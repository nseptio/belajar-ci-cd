package services

import (
	"context"
	"errors"
	"time"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/repositories"
)

var ErrStudentNotFound = errors.New("student not found")

type UserService interface {
	CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
	GetAllStudents(ctx context.Context) ([]*models.Student, error)
	GetStudentByID(ctx context.Context, id string) (*models.Student, error)
	UpdateStudent(ctx context.Context, id string, student *models.Student) (*models.Student, error)
	DeleteStudent(ctx context.Context, id string) error
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repository: repo,
	}
}

func (s *userService) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	// check if email or phone number already exists
	students, err := s.repository.GetAllStudents(ctx)
	if err != nil {
		return nil, err
	}

	for _, s := range students {
		if s.Email == student.Email || s.PhoneNumber == student.PhoneNumber {
			return nil, errors.New("email or phone number already exists")
		}
	}

	student.CreatedAt = time.Now().Unix()
	student.UpdatedAt = time.Now().Unix()

	return s.repository.CreateStudent(ctx, student)
}

func (s *userService) GetAllStudents(ctx context.Context) ([]*models.Student, error) {
	return s.repository.GetAllStudents(ctx)
}

func (s *userService) GetStudentByID(ctx context.Context, id string) (*models.Student, error) {
	// check if student exists
	_, err := s.repository.GetStudentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.repository.GetStudentByID(ctx, id)
}

func (s *userService) UpdateStudent(ctx context.Context, id string, student *models.Student) (*models.Student, error) {
	// check if student exists
	_, err := s.repository.GetStudentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	student.UpdatedAt = time.Now().Unix()

	return s.repository.UpdateStudent(ctx, id, student)
}

func (s *userService) DeleteStudent(ctx context.Context, id string) error {
	// check if student exists
	_, err := s.repository.GetStudentByID(ctx, id)
	if err != nil {
		return ErrStudentNotFound
	}

	return s.repository.DeleteStudent(ctx, id)
}

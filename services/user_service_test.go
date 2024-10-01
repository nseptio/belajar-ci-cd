package services

import (
	"context"
	"errors"
	"testing"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserRepository struct {
	mock.Mock
}

// Mock for UserRepository
func (m *MockUserRepository) CreateStudent(ctx context.Context, user *models.Student) (*models.Student, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*models.Student), args.Error(1)
}

func (m *MockUserRepository) GetAllStudents(ctx context.Context) ([]*models.Student, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Student), args.Error(1)
}

func (m *MockUserRepository) GetStudentByID(ctx context.Context, id string) (*models.Student, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Student), args.Error(1)
}

func (m *MockUserRepository) UpdateStudent(ctx context.Context, id string, user *models.Student) (*models.Student, error) {
	args := m.Called(ctx, id, user)
	return args.Get(0).(*models.Student), args.Error(1)
}

func (m *MockUserRepository) DeleteStudent(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Tests
func TestCreateStudent_Success(t *testing.T) {
	// Initialize the mock repository
	mockUserRepo := new(MockUserRepository)

	// Initialize the service with the mock repository
	service := NewUserService(mockUserRepo)

	// Define the student to be created
	user := &models.Student{
		Name:           "John Doe",
		Email:          "user1@example.com",
		PhoneNumber:    "081234567890",
		UniversityName: "University 1",
		StartYear:      2021,
		IsActive:       true,
	}

	// Set up the expected calls and return values

	// Expectation for GetAllStudents called within CreateStudent
	mockUserRepo.On("GetAllStudents", mock.Anything).Return([]*models.Student{}, nil)

	// Expectation for CreateStudent
	// Using mock.AnythingOfType("*models.Student") to allow flexibility in the student struct
	mockUserRepo.On("CreateStudent", mock.Anything, mock.AnythingOfType("*models.Student")).Return(user, nil)

	// Invoke the service method
	createdUser, err := service.CreateStudent(context.TODO(), user)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, user, createdUser)

	// Assert that all expectations were met
	mockUserRepo.AssertExpectations(t)
}

func TestGetAllStudents(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	service := NewUserService(mockUserRepo)

	users := []*models.Student{
		{
			Name:           "Harry",
			Email:          "user1@example.com",
			PhoneNumber:    "081234567890",
			UniversityName: "University 1",
			StartYear:      2021,
			IsActive:       true,
		},
		{
			Name:           "Ahmad Khasim",
			Email:          "user2@example.com",
			PhoneNumber:    "081234567891",
			UniversityName: "University 2",
			StartYear:      2021,
			IsActive:       true,
		},
	}

	mockUserRepo.On("GetAllStudents", mock.Anything).Return(users, nil)

	result, err := service.GetAllStudents(context.TODO())

	assert.NoError(t, err)
	assert.Equal(t, users, result)
	assert.Len(t, result, 2)
}

func TestGetAllStudents_Error(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	service := NewUserService(mockUserRepo)

	mockUserRepo.On("GetAllStudents", mock.Anything).Return(([]*models.Student)(nil), errors.New("Failed to fetch data"))

	users, err := service.GetAllStudents(context.TODO())

	assert.Error(t, err)
	assert.Nil(t, users)
}

func TestGetStudentByID_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	service := NewUserService(mockUserRepo)

	userID := primitive.NewObjectID().Hex()
	user := &models.Student{
		ID:             primitive.NewObjectID(),
		Name:           "John Doe",
		Email:          "user1@example.com",
		PhoneNumber:    "081234567890",
		UniversityName: "University 1",
		StartYear:      2021,
		IsActive:       true,
	}

	mockUserRepo.On("GetStudentByID", mock.Anything, userID).Return(user, nil)

	result, err := service.GetStudentByID(context.Background(), userID)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

func TestGetStudentByID_InvalidID(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	service := NewUserService(mockUserRepo)

	userID := primitive.NewObjectID().Hex()

	mockUserRepo.On("GetStudentByID", mock.Anything, userID).Return((*models.Student)(nil), errors.New("User not found"))

	result, err := service.GetStudentByID(context.Background(), userID)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateStudent_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	service := NewUserService(mockUserRepo)

	userID := primitive.NewObjectID().Hex()
	user := &models.Student{
		ID:             primitive.NewObjectID(),
		Name:           "John Doe",
		Email:          "user1@example.com",
		PhoneNumber:    "081234567890",
		UniversityName: "University 1",
		StartYear:      2021,
		IsActive:       true,
	}

	mockUserRepo.On("GetStudentByID", mock.Anything, userID).Return(user, nil)

	user.Name = "John Doe Updated"
	mockUserRepo.On("UpdateStudent", mock.Anything, userID, user).Return(user, nil)

	updatedUser, err := service.UpdateStudent(context.Background(), userID, user)

	assert.NoError(t, err)
	assert.Equal(t, user, updatedUser)
	assert.Equal(t, "John Doe Updated", updatedUser.Name)
}

func TestUpdateStudent_NotFound(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	service := NewUserService(mockUserRepo)

	userID := primitive.NewObjectID().Hex()
	user := &models.Student{
		ID:             primitive.NewObjectID(),
		Name:           "John Doe",
		Email:          "user1@example.com",
		PhoneNumber:    "081234567890",
		UniversityName: "University 1",
		StartYear:      2021,
		IsActive:       true,
	}

	mockUserRepo.On("GetStudentByID", mock.Anything, userID).Return((*models.Student)(nil), errors.New("User not found"))

	updatedUser, err := service.UpdateStudent(context.Background(), userID, user)

	assert.Error(t, err)
	assert.Nil(t, updatedUser)
}

func TestDeleteStudent_Success(t *testing.T) {
	// Initialize the mock repository
	mockUserRepo := new(MockUserRepository)

	// Initialize the service with the mock repository
	service := NewUserService(mockUserRepo)

	// Generate a new user ID
	userID := primitive.NewObjectID().Hex()

	// Define a mock student that exists
	existingStudent := &models.Student{
		ID:             primitive.NewObjectID(),
		Name:           "Jane Doe",
		Email:          "jane.doe@example.com",
		PhoneNumber:    "081234567891",
		UniversityName: "University 2",
		StartYear:      2020,
		IsActive:       true,
		CreatedAt:      primitive.NilObjectID.Timestamp().Unix(),
		UpdatedAt:      primitive.NilObjectID.Timestamp().Unix(),
	}

	// Set up the expectation for GetStudentByID
	mockUserRepo.On("GetStudentByID", mock.Anything, userID).Return(existingStudent, nil)

	// Set up the expectation for DeleteStudent
	mockUserRepo.On("DeleteStudent", mock.Anything, userID).Return(nil)

	// Invoke the service method
	err := service.DeleteStudent(context.Background(), userID)

	// Assertions
	assert.NoError(t, err)

	// Assert that all expectations were met
	mockUserRepo.AssertExpectations(t)
}

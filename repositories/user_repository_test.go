package repositories

import (
	"context"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/utils"
	"github.com/stretchr/testify/suite"
)

type UserRepositorySuite struct {
	suite.Suite
	repository   UserRepository
	testDatabase *utils.TestDatabase
}

func (suite *UserRepositorySuite) SetupSuite() {
	suite.testDatabase = utils.SetupTestDatabase()
	suite.repository = NewUserRepository(suite.testDatabase.DbInstance)
}

func (suite *UserRepositorySuite) TearDownSuite() {
	suite.testDatabase.TearDown()
}

func (suite *UserRepositorySuite) TestCreateStudent() {
	user := &models.Student{
		Name:           "User 1",
		Email:          "user1@example.com",
		PhoneNumber:    "08123456789",
		UniversityName: "University 1",
		StartYear:      2021,
		IsActive:       true,
	}

	createdUser, err := suite.repository.CreateStudent(context.Background(), user)
	suite.Require().NoError(err)
	suite.Require().NotNil(createdUser)
}

func (suite *UserRepositorySuite) TestGetAllStudents() {
	user := &models.Student{
		Name:           "User 1",
		Email:          "user1@example.com",
		PhoneNumber:    "08123456789",
		UniversityName: "University 1",
		StartYear:      2021,
		IsActive:       true,
	}

	user2 := &models.Student{
		Name:           "User 2",
		Email:          "user2@example.com",
		PhoneNumber:    "08122456789",
		UniversityName: "University 2",
		StartYear:      2021,
		IsActive:       true,
	}

	_, err := suite.repository.CreateStudent(context.Background(), user)
	suite.Require().NoError(err)
	_, err = suite.repository.CreateStudent(context.Background(), user2)
	suite.Require().NoError(err)

	users, err := suite.repository.GetAllStudents(context.Background())
	suite.Require().NoError(err)
	suite.Require().NotEmpty(users)
}

func (suite *UserRepositorySuite) TestGetStudentByID_Success() {
	user := &models.Student{
		Name:           "User 1",
		Email:          "user1@example.com",
		PhoneNumber:    "08123456789",
		UniversityName: "University 1",
		StartYear:      2021,
		IsActive:       true,
	}

	createdUser, err := suite.repository.CreateStudent(context.Background(), user)
	suite.Require().NoError(err)

	foundUser, err := suite.repository.GetStudentByID(context.Background(), createdUser.ID.Hex())

	suite.Require().NoError(err)
	suite.Require().NotNil(foundUser)
}

func (suite *UserRepositorySuite) TestGetStudentByID_InvalidID() {
	invalidID := "invalidHexID"

	user, err := suite.repository.GetStudentByID(context.Background(), invalidID)

	suite.Require().Error(err)
	suite.Require().Nil(user)
	suite.Require().EqualError(err, "invalid user ID")
}

func (suite *UserRepositorySuite) TestUpdateStudent_Success() {
	user := &models.Student{
		Name:           "User 1",
		Email:          "user1@example.com",
		PhoneNumber:    "08123456789",
		UniversityName: "University 1",
		StartYear:      2021,
		IsActive:       true,
	}

	createdUser, err := suite.repository.CreateStudent(context.Background(), user)
	suite.Require().NoError(err)

	createdUser.Name = "User 2"

	updatedUser, err := suite.repository.UpdateStudent(context.Background(), createdUser.ID.String(), &models.Student{
		Name:  "User 2",
		Email: "user2@example.com",
	})

	suite.Require().NoError(err)
	suite.Require().NotNil(updatedUser)
}

func (suite *UserRepositorySuite) TestUpdateStudent_InvalidID() {
	invalidID := "invalidHexID"

	user, err := suite.repository.UpdateStudent(context.Background(), invalidID, &models.Student{
		Name:  "User 2",
		Email: "user1@example.com",
	})

	suite.Require().Error(err)
	suite.Require().Nil(user)
	suite.Require().EqualError(err, "user with ID invalidHexID not found")
}

func (suite *UserRepositorySuite) TestDeleteStudent_Success() {
	user := &models.Student{
		Name:           "User 1",
		Email:          "user1@example.com",
		PhoneNumber:    "08123456789",
		UniversityName: "University 1",
		StartYear:      2021,
		IsActive:       true,
	}

	createdUser, err := suite.repository.CreateStudent(context.Background(), user)
	suite.Require().NoError(err)

	err = suite.repository.DeleteStudent(context.Background(), createdUser.ID.Hex())

	suite.Require().NoError(err)
}

func (suite *UserRepositorySuite) TestDeleteStudent_InvalidID() {
	invalidID := "invalidHexID"

	err := suite.repository.DeleteStudent(context.Background(), invalidID)

	suite.Require().Error(err)
	suite.Require().EqualError(err, "user with ID invalidHexID not found")
}

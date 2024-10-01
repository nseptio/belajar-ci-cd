package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/utils"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SurveyRepositorySuite struct {
	suite.Suite
	repository   SurveyRepository
	testDatabase *utils.TestDatabase
}

func (suite *SurveyRepositorySuite) SetupSuite() {
	suite.testDatabase = utils.SetupTestDatabase()
	suite.repository = NewSurveyRepository(suite.testDatabase.DbInstance)
}

func (suite *SurveyRepositorySuite) TearDownSuite() {
	suite.testDatabase.TearDown()
}

func (suite *SurveyRepositorySuite) TestCreateSurvey() {
	survey := &models.Survey{
		Title:       "Survey 1",
		Description: "Description 1",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		IsPublished: false,
		Year:        2024,
	}

	createdSurvey, err := suite.repository.CreateSurvey(context.Background(), survey)
	suite.Require().NoError(err)
	suite.Require().NotNil(createdSurvey)
	suite.Require().NotEqual(primitive.NilObjectID, createdSurvey.ID)
}

func (suite *SurveyRepositorySuite) TestGetAllSurveys() {
	survey := &models.Survey{
		Title:       "Survey 1",
		Description: "Description 1",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		IsPublished: false,
		Year:        2024,
	}

	_, err := suite.repository.CreateSurvey(context.Background(), survey)
	suite.Require().NoError(err)

	surveys, err := suite.repository.GetAllSurveys(context.Background())
	suite.Require().NoError(err)
	suite.Require().NotEmpty(surveys)
}

func (suite *SurveyRepositorySuite) TestGetSurveyByID_Success() {
	survey := &models.Survey{
		Title:       "Survey 1",
		Description: "Description 1",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		IsPublished: false,
		Year:        2024,
	}

	createdSurvey, err := suite.repository.CreateSurvey(context.Background(), survey)
	suite.Require().NoError(err)

	foundSurvey, err := suite.repository.GetSurveyByID(context.Background(), createdSurvey.ID.Hex())
	suite.Require().NoError(err)
	suite.Require().NotNil(foundSurvey)
}

func (suite *SurveyRepositorySuite) TestGetSurveyByID_InvalidID() {
	invalidID := "invalidHexID"

	survey, err := suite.repository.GetSurveyByID(context.Background(), invalidID)
	suite.Require().Error(err)
	suite.Require().Nil(survey)
	suite.Require().EqualError(err, "invalid survey ID")
}

func (suite *SurveyRepositorySuite) TestUpdateSurvey_Success() {
	survey := &models.Survey{
		Title:       "Survey 1",
		Description: "Description 1",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		IsPublished: false,
		Year:        2024,
	}

	createdSurvey, err := suite.repository.CreateSurvey(context.Background(), survey)
	suite.Require().NoError(err)

	createdSurvey.Title = "Survey 2"
	updatedSurvey, err := suite.repository.UpdateSurvey(context.Background(), createdSurvey.ID.Hex(), createdSurvey)
	suite.Require().NoError(err)
	suite.Require().NotNil(updatedSurvey)
	suite.Require().Equal("Survey 2", updatedSurvey.Title)
}

func (suite *SurveyRepositorySuite) TestUpdateSurvey_InvalidID() {
	invalidID := "invalidHexID"
	survey := &models.Survey{
		Title: "Updated Title",
	}

	updatedSurvey, err := suite.repository.UpdateSurvey(context.Background(), invalidID, survey)
	suite.Require().Error(err)
	suite.Require().Nil(updatedSurvey)
	suite.Require().EqualError(err, "invalid survey ID")
}

func (suite *SurveyRepositorySuite) TestUpdateSurvey_SurveyNotFound() {
	survey := &models.Survey{
		Title:       "Survey 1",
		Description: "Description 1",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		IsPublished: false,
		Year:        2024,
	}

	createdSurvey, err := suite.repository.CreateSurvey(context.Background(), survey)
	suite.Require().NoError(err)

	createdSurvey.ID = primitive.NewObjectID()
	updatedSurvey, err := suite.repository.UpdateSurvey(context.Background(), createdSurvey.ID.Hex(), createdSurvey)
	suite.Require().Error(err)
	suite.Require().Nil(updatedSurvey)
}

func (suite *SurveyRepositorySuite) TestDeleteSurvey_Success() {
	survey := &models.Survey{
		Title:       "Survey 1",
		Description: "Description 1",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		IsPublished: false,
		Year:        2024,
	}

	createdSurvey, err := suite.repository.CreateSurvey(context.Background(), survey)
	suite.Require().NoError(err)

	err = suite.repository.DeleteSurvey(context.Background(), createdSurvey.ID.Hex())
	suite.Require().NoError(err)

	_, err = suite.repository.GetSurveyByID(context.Background(), createdSurvey.ID.Hex())
	suite.Require().Error(err)
}

func (suite *SurveyRepositorySuite) TestDeleteSurvey_InvalidID() {
	invalidID := "invalidHexID"

	err := suite.repository.DeleteSurvey(context.Background(), invalidID)
	suite.Require().Error(err)
	suite.Require().EqualError(err, "invalid survey ID")
}

func (suite *SurveyRepositorySuite) TestDeleteSurvey_SurveyNotFound() {
	err := suite.repository.DeleteSurvey(context.Background(), primitive.NewObjectID().Hex())
	suite.Require().Error(err)
}

func TestSurveyRepositorySuite(t *testing.T) {
	suite.Run(t, new(SurveyRepositorySuite))
}

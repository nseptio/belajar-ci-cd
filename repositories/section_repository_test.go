package repositories

import (
	"context"
	"testing"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/utils"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SectionRepositorySuite struct {
	suite.Suite
	repository   SectionRepository
	testDatabase *utils.TestDatabase
}

func (suite *SectionRepositorySuite) SetupSuite() {
	suite.testDatabase = utils.SetupTestDatabase()
	suite.repository = NewSectionRepository(suite.testDatabase.DbInstance)
}

func (suite *SectionRepositorySuite) TearDownSuite() {
	suite.testDatabase.TearDown()
}

func (suite *SectionRepositorySuite) TestCreateSection() {
	section := &models.Section{
		Title: "Section 1",
	}

	createdSection, err := suite.repository.CreateSection(context.Background(), section)
	suite.Require().NoError(err)
	suite.Require().NotNil(createdSection)
	suite.Require().NotEqual(primitive.NilObjectID, createdSection.ID)
}

func (suite *SectionRepositorySuite) TestGetAllSections() {
	section := &models.Section{
		Title: "Section 1",
	}

	_, err := suite.repository.CreateSection(context.Background(), section)
	suite.Require().NoError(err)

	sections, err := suite.repository.GetAllSections(context.Background())
	suite.Require().NoError(err)
	suite.Require().NotEmpty(sections)
}

func (suite *SectionRepositorySuite) TestGetSectionByID_Success() {
	section := &models.Section{
		Title: "Section 1",
	}

	createdSection, err := suite.repository.CreateSection(context.Background(), section)
	suite.Require().NoError(err)

	foundSection, err := suite.repository.GetSectionByID(context.Background(), createdSection.ID.Hex())
	suite.Require().NoError(err)
	suite.Require().NotNil(foundSection)
}

func (suite *SectionRepositorySuite) TestGetSectionByID_InvalidID() {
	invalidID := "invalidHexID"

	section, err := suite.repository.GetSectionByID(context.Background(), invalidID)
	suite.Require().Error(err)
	suite.Require().Nil(section)
	suite.Require().EqualError(err, "invalid section ID")
}

func (suite *SectionRepositorySuite) TestUpdateSection_Success() {
	section := &models.Section{
		Title: "Section 1",
	}

	createdSection, err := suite.repository.CreateSection(context.Background(), section)
	suite.Require().NoError(err)

	createdSection.Title = "Section 2"
	updatedSection, err := suite.repository.UpdateSection(context.Background(), createdSection.ID.Hex(), createdSection)
	suite.Require().NoError(err)
	suite.Require().NotNil(updatedSection)
	suite.Require().Equal("Section 2", updatedSection.Title)
}

func (suite *SectionRepositorySuite) TestUpdateSection_InvalidID() {
	invalidID := "invalidHexID"
	section := &models.Section{
		Title: "Updated Title",
	}

	updatedSection, err := suite.repository.UpdateSection(context.Background(), invalidID, section)
	suite.Require().Error(err)
	suite.Require().Nil(updatedSection)
	suite.Require().EqualError(err, "invalid section ID")
}

func (suite *SectionRepositorySuite) TestUpdateSection_SectionNotFound() {
	section := &models.Section{
		Title: "Section 1",
	}

	createdSection, err := suite.repository.CreateSection(context.Background(), section)
	suite.Require().NoError(err)

	createdSection.ID = primitive.NewObjectID()
	updatedSection, err := suite.repository.UpdateSection(context.Background(), createdSection.ID.Hex(), createdSection)
	suite.Require().Error(err)
	suite.Require().Nil(updatedSection)
}

func (suite *SectionRepositorySuite) TestDeleteSection_Success() {
	section := &models.Section{
		Title: "Section 1",
	}

	createdSection, err := suite.repository.CreateSection(context.Background(), section)
	suite.Require().NoError(err)

	err = suite.repository.DeleteSection(context.Background(), createdSection.ID.Hex())
	suite.Require().NoError(err)

	_, err = suite.repository.GetSectionByID(context.Background(), createdSection.ID.Hex())
	suite.Require().Error(err)
}

func (suite *SectionRepositorySuite) TestDeleteSection_InvalidID() {
	invalidID := "invalidHexID"

	err := suite.repository.DeleteSection(context.Background(), invalidID)
	suite.Require().Error(err)
	suite.Require().EqualError(err, "invalid section ID")
}

func (suite *SectionRepositorySuite) TestDeleteSection_SectionNotFound() {
	err := suite.repository.DeleteSection(context.Background(), primitive.NewObjectID().Hex())
	suite.Require().Error(err)
}

func TestSectionRepositorySuite(t *testing.T) {
	suite.Run(t, new(SectionRepositorySuite))
}

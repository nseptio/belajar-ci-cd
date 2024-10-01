package repositories

import (
	"context"
	"errors"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SurveyRepository interface {
	CreateSurvey(ctx context.Context, survey *models.Survey) (*models.Survey, error)
	GetAllSurveys(ctx context.Context) ([]*models.Survey, error)
	GetSurveyByID(ctx context.Context, id string) (*models.Survey, error)
	UpdateSurvey(ctx context.Context, id string, survey *models.Survey) (*models.Survey, error)
	DeleteSurvey(ctx context.Context, id string) error
}

type surveyRepository struct {
	collection *mongo.Collection
}

func NewSurveyRepository(db *mongo.Database) SurveyRepository {
	return &surveyRepository{
		collection: db.Collection("surveys"),
	}
}

func (r *surveyRepository) CreateSurvey(ctx context.Context, survey *models.Survey) (*models.Survey, error) {
	survey.ID = primitive.NewObjectID()

	_, err := r.collection.InsertOne(ctx, survey)
	if err != nil {
		return nil, err
	}

	return survey, nil
}

func (r *surveyRepository) GetAllSurveys(ctx context.Context) ([]*models.Survey, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var surveys []*models.Survey
	err = cursor.All(ctx, &surveys)
	if err != nil {
		return nil, err
	}

	return surveys, nil
}

func (r *surveyRepository) GetSurveyByID(ctx context.Context, id string) (*models.Survey, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid survey ID")
	}

	var survey models.Survey
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&survey)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("survey not found")
	}

	return &survey, nil
}

func (r *surveyRepository) UpdateSurvey(ctx context.Context, id string, survey *models.Survey) (*models.Survey, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid survey ID")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": survey}

	result, _ := r.collection.UpdateOne(ctx, filter, update)

	if result.MatchedCount == 0 {
		return nil, errors.New("survey not found")
	}

	return survey, nil
}

func (r *surveyRepository) DeleteSurvey(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid survey ID")
	}

	result, _ := r.collection.DeleteOne(ctx, bson.M{"_id": objID})

	if result.DeletedCount == 0 {
		return errors.New("survey not found")
	}

	return nil
}

package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SectionRepository interface {
	CreateSection(ctx context.Context, section *models.Section) (*models.Section, error)
	GetAllSections(ctx context.Context) ([]*models.Section, error)
	GetSectionByID(ctx context.Context, id string) (*models.Section, error)
	UpdateSection(ctx context.Context, id string, section *models.Section) (*models.Section, error)
	DeleteSection(ctx context.Context, id string) error
}

type sectionRepository struct {
	collection *mongo.Collection
}

func NewSectionRepository(db *mongo.Database) SectionRepository {
	return &sectionRepository{
		collection: db.Collection("sections"),
	}
}

func (r *sectionRepository) CreateSection(ctx context.Context, section *models.Section) (*models.Section, error) {
	section.ID = primitive.NewObjectID()

	_, err := r.collection.InsertOne(ctx, section)
	if err != nil {
		return nil, fmt.Errorf("failed to insert section: %w", err)
	}

	return section, nil
}

func (r *sectionRepository) GetAllSections(ctx context.Context) ([]*models.Section, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find sections: %w", err)
	}
	defer cursor.Close(ctx)

	var sections []*models.Section
	if err := cursor.All(ctx, &sections); err != nil {
		return nil, fmt.Errorf("failed to decode sections: %w", err)
	}

	return sections, nil
}

func (r *sectionRepository) GetSectionByID(ctx context.Context, id string) (*models.Section, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid section ID")
	}

	var section models.Section
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&section)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("section not found")
	}

	return &section, nil
}

func (r *sectionRepository) UpdateSection(ctx context.Context, id string, section *models.Section) (*models.Section, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid section ID")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": section}

	result, _ := r.collection.UpdateOne(ctx, filter, update)

	if result.MatchedCount == 0 {
		return nil, errors.New("section not found")
	}

	return section, nil
}

func (r *sectionRepository) DeleteSection(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid section ID")
	}

	result, _ := r.collection.DeleteOne(ctx, bson.M{"_id": objID})

	if result.DeletedCount == 0 {
		return errors.New("section not found")
	}

	return nil
}

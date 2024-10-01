package repositories

import (
	"context"
	"fmt"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateStudent(ctx context.Context, user *models.Student) (*models.Student, error)
	GetAllStudents(ctx context.Context) ([]*models.Student, error)
	GetStudentByID(ctx context.Context, id string) (*models.Student, error)
	UpdateStudent(ctx context.Context, id string, user *models.Student) (*models.Student, error)
	DeleteStudent(ctx context.Context, id string) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (r *userRepository) CreateStudent(ctx context.Context, user *models.Student) (*models.Student, error) {
	user.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	return user, nil
}

func (r *userRepository) DeleteStudent(ctx context.Context, id string) error {
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": userID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("user with ID %s not found", id)
	}

	return nil
}

func (r *userRepository) GetAllStudents(ctx context.Context) ([]*models.Student, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}
	defer cursor.Close(ctx)

	var users []*models.Student
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to decode users: %v", err)
	}

	return users, nil
}

func (r *userRepository) GetStudentByID(ctx context.Context, id string) (*models.Student, error) {
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.Student
	err = r.collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}

	return &user, nil
}

func (r *userRepository) UpdateStudent(ctx context.Context, id string, user *models.Student) (*models.Student, error) {
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": userID}
	update := bson.M{"$set": user}

	result, err := r.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}

	return user, nil
}

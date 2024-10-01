package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Survey struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	StartDate   time.Time          `bson:"start_date" json:"start_date"`
	EndDate     time.Time          `bson:"end_date" json:"end_date"`
	IsPublished bool               `bson:"is_published" json:"is_published"`
	Year        int                `bson:"year" json:"year"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SectionID primitive.ObjectID `bson:"section_id" json:"section_id"`
	Title     string             `bson:"title" json:"title"`
	Type      string             `bson:"type" json:"type"`
}

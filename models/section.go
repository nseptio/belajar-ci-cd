package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Section struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	SurveyID    primitive.ObjectID `bson:"survey_id" json:"survey_id"`
}

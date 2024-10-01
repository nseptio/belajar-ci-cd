package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TextAnswer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	QuestionID primitive.ObjectID `bson:"question_id" json:"question_id"`
}

type MultipleAnswer struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	QuestionID    primitive.ObjectID `bson:"question_id" json:"question_id"`
	ListOfOptions []string           `bson:"list_of_options" json:"list_of_options"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuestionAnswer struct {
	QuestionID primitive.ObjectID `bson:"question_id" json:"question_id"`
	Answer     string             `bson:"answer" json:"answer"`
}

type Response struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SurveyID           primitive.ObjectID `bson:"survey_id" json:"survey_id"`
	StudentID          primitive.ObjectID `bson:"student_id" json:"student_id"`
	ListQuestionAnswer []QuestionAnswer   `bson:"list_question_answer" json:"list_question_answer"`
}

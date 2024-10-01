package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
}

type Student struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           string             `bson:"name" json:"name"`
	Email          string             `bson:"email" json:"email"`
	PhoneNumber    string             `bson:"phone_number" json:"phone_number"`
	UniversityName string             `bson:"university_name" json:"university_name"`
	StartYear      int                `bson:"start_year" json:"start_year"`
	IsActive       bool               `bson:"is_active" json:"is_active"`
	CreatedAt      int64              `bson:"created_at" json:"created_at"`
	UpdatedAt      int64              `bson:"updated_at" json:"updated_at"`
}

type UniversityStaff struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           string             `bson:"name" json:"name"`
	Email          string             `bson:"email" json:"email"`
	UniversityName string             `bson:"university_name" json:"university_name"`
}

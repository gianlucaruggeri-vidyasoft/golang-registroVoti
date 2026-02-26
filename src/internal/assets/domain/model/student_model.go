package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type StudentModel struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty"`
	Name    string               `bson:"name"`
	Surname string               `bson:"surname"`
}
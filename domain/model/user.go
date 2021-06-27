package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	name     string             `json:"name" bson:"name,omitempty"`
	email    string             `json:"email" bson:"email,omitempty"`
	password string             `json:"password,omitempty" bson:"password,omitempty"`
}

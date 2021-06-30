package model

import (
	"time"
)

type User struct {
	// TODO: Change to uuid
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string    `json:"username" bson:"username,omitempty"`
	Email     string    `json:"email" bson:"email,omitempty"`
	Password  string    `json:"password" bson:"password,omitempty"`
	Salt      string    `json:"-" bson:"salt"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

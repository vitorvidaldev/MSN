package model

import "time"

type Message struct {
	ID         string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Body       string    `json:"body,omitempty" bson:"body"`
	SenderId   string    `json:"senderId,omitempty" bson:"senderId,omitempty"`
	ReceiverId string    `json:"receiverId,omitempty" bson:"receiverId,omitempty"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

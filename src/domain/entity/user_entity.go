package entity

import "github.com/google/uuid"

type User struct {
	User_id  uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

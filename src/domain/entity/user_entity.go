package entity

import "github.com/google/uuid"

type User struct {
	User_id  uuid.UUID // Primary Key
	Email    string
	Username string
	Password string
}

package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// TODO: Change to uuid
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string    `json:"username" bson:"username,omitempty"`
	Email     string    `json:"email" bson:"email,omitempty"`
	Hash      string    `json:"hash" bson:"hash,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

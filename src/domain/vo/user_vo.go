package vo

import "github.com/google/uuid"

type UserVO struct {
	User_id  uuid.UUID // Primary Key
	Email    string
	Username string
}

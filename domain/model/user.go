package model

import (
	"log"
	"time"

	"github.com/vitorvidaldev/MSN/application/vo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string    `json:"username" bson:"username"`
	Email     string    `json:"email" bson:"email"`
	Hash      string    `json:"hash" bson:"hash,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal("There was a problem with your password")
	}
	return string(bytes)
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func FromCreateVO(vo vo.CreateUserVO) User {
	var user User
	user.Email = vo.Email
	user.Username = vo.Username
	user.Hash = HashPassword(vo.Password)
	return user
}

func FromUpdateVO(vo vo.UpdateUserVO) User {
	var user User
	user.Email = vo.Email
	user.Username = vo.Username
	return user
}

func ToReturnVO(user User) vo.ReturnUserVO {
	var returnUserVO vo.ReturnUserVO
	returnUserVO.ID = user.ID
	returnUserVO.Username = user.Username
	returnUserVO.Email = user.Email
	return returnUserVO
}

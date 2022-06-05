package repository

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/vitorvidaldev/msn/src/config"
	"github.com/vitorvidaldev/msn/src/domain/entity"
	"github.com/vitorvidaldev/msn/src/domain/vo"
)

func InsertUser(user entity.User) entity.User {
	db := config.InitPGConnection()

	sqlStatement := `insert into "user"("user_id", "email", "username", "password") values ($1, $2, $3, $4);`

	_, err := db.Exec(sqlStatement, user.User_id, user.Email, user.Username, user.Password)

	// TODO: return http 409 when user with email already exists
	if err != nil {
		log.Fatal("[UserRepository] Error inserting user. ", err)
	}

	defer db.Close()
	return user
}

func LoginUser(loginVO vo.LoginUserVO) (entity.User, error) {
	db := config.InitPGConnection()

	sqlStatement := `select * from "user" where email = $1;`

	var user entity.User
	err := db.QueryRow(sqlStatement, loginVO.Email).Scan(&user)

	if err != nil {
		log.Fatal("[UserRepository] Error in user login. ", err)
	}

	defer db.Close()
	if loginVO.Password == user.Password {
		return user, nil
	}
	log.Fatal("[UserRepository] Password didn't match ")

	return entity.User{}, errors.New("[UserRepository] Password didn't match ")
}

func GetUser(userId uuid.UUID) entity.User {
	db := config.InitPGConnection()

	sqlStatement := `select * from "user" where user_id = $1;`

	var user entity.User
	err := db.QueryRow(sqlStatement, userId).Scan(&user)

	if err != nil {
		log.Fatal("[UserRepository] Error getting user. ", err)
	}

	defer db.Close()

	return user
}

func DeleteUser(userId uuid.UUID) {
	db := config.InitPGConnection()

	sqlStatement := `delete from "user" where user_id = $1;`
	_, err := db.Exec(sqlStatement, userId)
	if err != nil {
		log.Fatal("[UserRepository] Error deleting user. ", err)
	}
}

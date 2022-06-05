package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/vitorvidaldev/msn/src/config"
	"github.com/vitorvidaldev/msn/src/domain/entity"
)

func InsertUser(user entity.User) entity.User {
	db := config.InitPGConnection()

	sqlStatement := `insert into "user"("user_id", "email", "username", "password") values ($1, $2, $3, $4);`

	_, err := db.Exec(sqlStatement, user.User_id, user.Email, user.Username, user.Password)

	// TODO: return http 409 when user with email already exists
	if err != nil {
		log.Fatal("[UserRepository] Error inserting user in database. ", err)
	}

	defer db.Close()
	return user
}

func DeleteUser(userId uuid.UUID) {

}

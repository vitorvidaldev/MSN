package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/vitorvidaldev/msn/src/domain/entity"
	"github.com/vitorvidaldev/msn/src/domain/vo"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user entity.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode create user request body. %v", err)
	}

	// userId := createUser
	userId := uuid.New()

	res := vo.UserVO{UserId: userId}

	json.NewEncoder(w).Encode(res)

}

func LoginUser(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

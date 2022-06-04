package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vitorvidaldev/msn/src/application/service"
	"github.com/vitorvidaldev/msn/src/domain/vo"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var createUserVO vo.CreateUserVO

	err := json.NewDecoder(r.Body).Decode(&createUserVO)
	if err != nil {
		log.Fatalf("Unable to decode create user request body. %v", err)
	}

	userVO := service.CreateUser(createUserVO)

	json.NewEncoder(w).Encode(userVO)

}

func LoginUser(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

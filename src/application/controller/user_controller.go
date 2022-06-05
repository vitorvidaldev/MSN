package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vitorvidaldev/msn/src/application/service"
	"github.com/vitorvidaldev/msn/src/domain/vo"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var createUserVO vo.CreateUserVO

	err := json.NewDecoder(r.Body).Decode(&createUserVO)
	if err != nil {
		log.Fatalf("[UserController] Unable to decode create user request body. %v", err)
	}

	userVO := service.CreateUser(createUserVO)

	w.WriteHeader(201)

	json.NewEncoder(w).Encode(userVO)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "POST")

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "GET")

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

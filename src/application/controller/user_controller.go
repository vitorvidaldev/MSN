package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userVO)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var loginUserVO vo.LoginUserVO

	err := json.NewDecoder(r.Body).Decode(&loginUserVO)
	if err != nil {
		log.Fatalf("[UserController] Unable to decode login user request body. %v", err)
	}

	userVO := service.LoginUser(loginUserVO)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userVO)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	userId := stringToUUID(extractPathVariable(r, "userId"))

	userVO := service.GetUser(userId)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userVO)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	userId := stringToUUID(extractPathVariable(r, "userId"))

	service.DeleteUser(userId)

	w.WriteHeader(http.StatusNoContent)
}

func stringToUUID(stringId string) uuid.UUID {
	userId, err := uuid.Parse(stringId)

	if err != nil {
		log.Fatal("[userController] Could not convert id to uuid. ", err)
	}
	return userId
}

func extractPathVariable(r *http.Request, name string) string {
	vars := mux.Vars(r)

	stringId, doesExists := vars[name]
	if !doesExists {
		log.Fatal("[UserController] User id not provided.")
	}
	return stringId
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

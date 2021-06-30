package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitorvidaldev/MSN/application/vo"
	"github.com/vitorvidaldev/MSN/domain/model"
	"github.com/vitorvidaldev/MSN/domain/repository"
	"github.com/vitorvidaldev/MSN/infra/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []model.User

	cur := repository.FindAll()
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var user model.User

		err := cur.Decode(&user)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	id := extractID(r)
	err := repository.FindById(id, &user)

	if err != nil {
		util.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userVO vo.UserVO
	_ = json.NewDecoder(r.Body).Decode(&userVO)
	result, err := repository.CreateUser(userVO)

	if err != nil {
		util.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	id := extractID(r)

	_ = json.NewDecoder(r.Body).Decode(&user)

	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "username", Value: user.Username},
				primitive.E{Key: "email", Value: user.Email},
				primitive.E{Key: "password", Value: user.Password},
			},
		}}

	err := repository.UpdateUser(id, update, &user)

	if err != nil {
		util.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(id)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := extractID(r)
	deleteResult, err := repository.DeleteUser(id)

	if err != nil {
		util.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

func extractID(r *http.Request) primitive.ObjectID {
	var params = mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	return id
}

package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitorvidaldev/MSN/domain/model"
	"github.com/vitorvidaldev/MSN/infra/config"
	"github.com/vitorvidaldev/MSN/infra/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = config.MongoConfig()

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []model.User

	// bson.M{} -> We passed an empty filter, so we want to get all data
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		util.GetError(err, w)
		return
	}

	// Close the cursor once finished
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

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		util.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		util.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// TODO: Fix save to collection bug
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var user model.User

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&user)

	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "name", Value: user.Name},
				primitive.E{Key: "email", Value: user.Email},
				primitive.E{Key: "password", Value: user.Password},
			},
		}}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)

	if err != nil {
		util.GetError(err, w)
		return
	}

	user.ID = id

	json.NewEncoder(w).Encode(user)
}

// TODO: Fix warning
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		util.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

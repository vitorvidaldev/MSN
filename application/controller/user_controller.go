package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vitorvidaldev/MSN/domain/model"
	"github.com/vitorvidaldev/MSN/infra/config"
	"github.com/vitorvidaldev/MSN/infra/util"
	"go.mongodb.org/mongo-driver/bson"
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
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

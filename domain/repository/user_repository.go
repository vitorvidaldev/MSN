package repository

import (
	"context"
	"log"

	"github.com/vitorvidaldev/MSN/application/vo"
	"github.com/vitorvidaldev/MSN/domain/model"
	"github.com/vitorvidaldev/MSN/infra/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = config.MongoConfig()

func FindAll() *mongo.Cursor {
	// TODO: Never return password
	curr, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	return curr
}

func FindById(id primitive.ObjectID, user *model.User) error {
	filter := bson.M{"_id": id}
	// TODO: Never return password
	return collection.FindOne(context.TODO(), filter).Decode(user)

}

func CreateUser(userVO vo.UserVO) (*mongo.InsertOneResult, error) {
	var user model.User
	// TODO: Implement explicit conversion between userVO and user
	return collection.InsertOne(context.TODO(), user)
}

func UpdateUser(id primitive.ObjectID, update primitive.D, user *model.User) error {
	// TODO: Never return password
	filter := bson.M{"_id": id}
	return collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(user)
}

func DeleteUser(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	// TODO: Never return password
	filter := bson.M{"_id": id}
	return collection.DeleteOne(context.TODO(), filter)
}

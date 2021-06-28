package repository

import (
	"context"

	"github.com/vitorvidaldev/MSN/domain/model"
	"github.com/vitorvidaldev/MSN/infra/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = config.MongoConfig()

// bson.M{} -> We passed an empty filter, so we want to get all data
func FindAll() (*mongo.Cursor, error) {
	return collection.Find(context.TODO(), bson.M{})
}

func FindById(filter primitive.M, user *model.User) error {
	return collection.FindOne(context.TODO(), filter).Decode(user)
}

func CreateUser(user model.User) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(context.TODO(), user)
}

func UpdateUser(filter primitive.M, update primitive.D, user *model.User) error {
	return collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(user)
}

func DeleteUser(filter primitive.M) (*mongo.DeleteResult, error) {
	return collection.DeleteOne(context.TODO(), filter)
}

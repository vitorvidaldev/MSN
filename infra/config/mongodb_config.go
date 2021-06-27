package config

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConfig() *mongo.Collection {
	credential := options.Credential{
		Username: "msn_",
		Password: "mongo_password",
	}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	collection := client.Database("msn_database").Collection("users")

	return collection
}

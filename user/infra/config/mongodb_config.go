package config

import (
	"context"
	"fmt"
	"log"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("msn_database").Collection("users")

	return collection
}

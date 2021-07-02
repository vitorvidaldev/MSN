package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vitorvidaldev/MSN/infra/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConfig() *mongo.Collection {
	util.LoadEnv()
	credential := options.Credential{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")).SetAuth(credential))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return userCollection(client)
}

func userCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database(os.Getenv("MSN_DATABASE")).Collection(os.Getenv("USER_COLLECTION"))

	return collection
}

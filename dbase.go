package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connection string for MongoDB
var  mongoConnect = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000"
// Handler for all MongoDB operations
var mgH *mongo.Client = initMongo()

// initializes and returns a Go MongoDB client
func initMongo() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoConnect))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
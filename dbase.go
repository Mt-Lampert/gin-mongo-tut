package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Episode struct {
	Title     string `bson:"title"     json:"title"`
	Desc      string `bson:"desc"      json:"desc"`
	Duration  string `bson:"duration"  json:"duration"`
	CreatedAt string `bson:"createdAt" json:"createdAt"`
}

type Podcast struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"       json:"id"`
	Title    string             `bson:"title,omitempty"     json:"title"`
	Author   string             `bson:"author,omitempty"    json:"author"`
	Episodes []Episode          `bson:"episodes"  json:"episodes"`
}

// connection string for MongoDB
var mongoConnect = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000"

// Handler for all MongoDB operations
var mgH *mongo.Client = initMongo()

var dbase string = "GinMongoTut"
var coll string = "podcasts"

// initializes and returns a Go MongoDB client
func initMongo() *mongo.Client {
	ctx, delCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer delCtx()
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoConnect))
	if err != nil {
		log.Fatal("Could not connect to MongoDB: \n", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Could not set up a database handler: \n", err)
	}

	return client
}

func dbGetAllPodcasts() ([]Podcast, error) {
	ctx, delCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer delCtx()
	pcColl := mgH.Database(dbase).Collection(coll)
	// container for the podcasts found, initialized as empty slice
	podcasts := make([]Podcast, 0)

	// find all podcasts; bson.D{} means "find everything"
	cursor, err := pcColl.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal("Something went wrong reading the podcasts:\n", err)
		return nil, err
	}

	// fetch the podcasts being found
	if err = cursor.All(ctx, &podcasts); err != nil {
		log.Fatal("Could not convert found podcasts to Go type:\n", err)
		return nil, err
	}
	return podcasts, nil
}

func dbAddPodcast(pc *Podcast) error {
	ctx, delCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer delCtx()
	podcasts := mgH.Database(dbase).Collection(coll)

	_, err := podcasts.InsertOne(ctx, pc)
	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
//const uri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"
const uri = "mongodb+srv://mulop001:Blessing@cluster0.uejox.mongodb.net/test"

func DatabaseConnector() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	coll := client.Database("myDB").Collection("favorite_books")
	doc := bson.D{{"title", "Invisible Cities"}, {"author", "Italo Calvino"}, {"year_published", 1974}}

	result, err := coll.InsertOne(context.TODO(), doc)

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}

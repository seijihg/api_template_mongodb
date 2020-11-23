package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectDB is setting up database connection.
func ConnectDB() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://lehoang:WegTCYVpPINPmG3y@golangtemplate.uusoo.mongodb.net/<dbname>?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second) // Set Timeout of 3s
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connected. Lists:", databases)

	return client
}

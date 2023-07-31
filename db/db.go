package db

import (
	"context"
	"log"
	"time"

	"github.com/charlygame/CatGameService/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const (
	// DBName is the name of the database
	connectTimeout = 5
)

var (
	client   *mongo.Client
	database *mongo.Database
	ctx      context.Context
)

func Init() {
	config := config.GetConfig()
	connectionString := config.MongoURI

	parsedConnectionString, err := connstring.ParseAndValidate(connectionString)
	if err != nil {
		log.Fatalf("Error parsing connection string: %v", err)
	}

	client, err = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatalf("Error creating mongo client: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting to mongo: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Error pinging mongo: %v", err)
	}
	database = client.Database(parsedConnectionString.Database)
}

func Disconnect() {
	client.Disconnect(ctx)
	log.Print("Disconnected from MongoDB")
}

func ClearDB() {
	database.Drop(ctx)
}

func GetDB() *mongo.Database {
	return database
}

// db/mongo.go
package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

// InitMongo initializes MongoDB connection
func InitMongo(uri string) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	mongoClient = client
}

// GetMongoClient returns MongoDB client
func GetMongoClient() *mongo.Client {
	return mongoClient
}

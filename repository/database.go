package repository

import (
	"context"
	"github.com/gofiber/fiber/v3/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var LinksCollection *mongo.Collection

func ConnectDB() {
	client, err := mongo.Connect(nil, options.Client().ApplyURI("mongodb://admin:password@127.0.0.1:27017/"))
	if err != nil {
		log.Fatal("Error connect to mongoDB", err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal("Error connect to mongoDB", err)
	}
	log.Info("Connected to MongoDB")
	LinksCollection = client.Database("storage").Collection("storage-links")
}

package repository

import (
	"context"
	"github.com/gofiber/fiber/v3/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var LinksCollection *mongo.Collection

const url = "MONGO_URL"

func ConnectDB() {
	client, err := mongo.Connect(nil, options.Client().ApplyURI(os.Getenv(url)))
	if err != nil {
		log.Fatal("Error connect to mongoDB", err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal("Error connect to mongoDB", err)
	}
	log.Info("Connected to MongoDB")
	LinksCollection = client.Database("storage").Collection("storage-links")
}

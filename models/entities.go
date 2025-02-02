package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	LinksEntity struct {
		ID     primitive.ObjectID `bson:"_id,omitempty"`
		ChatId string             `bson:"chatId"`
		Links  []string           `bson:"links"`
	}
)

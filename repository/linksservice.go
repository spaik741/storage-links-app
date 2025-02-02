package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"storage-links-app/models"
)

var ctx = context.Background()

func FindByChatId(chatId string) (models.LinksEntity, error) {
	res := LinksCollection.FindOne(ctx, bson.M{"chatId": chatId})
	var linksEntity models.LinksEntity
	err := res.Decode(&linksEntity)
	if err != nil {
		return models.LinksEntity{}, err
	}
	return linksEntity, nil
}

func Save(linksEntity models.LinksEntity) error {
	if _, err := LinksCollection.InsertOne(ctx, linksEntity); err != nil {
		return err
	}
	return nil
}

func Update(linksEntity models.LinksEntity) error {
	filter := bson.M{"_id": linksEntity.ID}
	update := bson.M{"$set": linksEntity}
	if _, err := LinksCollection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil
}

package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Card struct {
	ID      string   `bson:"_id,omitempty"`
	Content []string `bson:"content"`
	Level   int      `bson:"level"`
}

type CardModel struct {
	DB     *mongo.Client
	DBName string
}

func (cm *CardModel) InsertCardIntoSet(c Card, setName string) error {
	coll := cm.DB.Database(cm.DBName).Collection(setName)

	insertResult, err := coll.InsertOne(
		context.TODO(),
		c,
	)

	if err != nil {
		return fmt.Errorf("cannot insert document into collection: %w", err)
	}

	fmt.Println(insertResult)

	return nil
}

func (cm *CardModel) GetAllCardsFromSet(setName string) ([]Card, error) {
	coll := cm.DB.Database(cm.DBName).Collection(setName)
	var cards []Card

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &cards); err != nil {
		return nil, err
	}

	return cards, nil
}

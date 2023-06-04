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
	DB  *mongo.Database
	Ctx context.Context
}

func (cm *CardModel) InsertCardIntoSet(c Card, setName string) error {
	coll := cm.DB.Collection(setName)

	insertResult, err := coll.InsertOne(
		cm.Ctx,
		c,
	)

	if err != nil {
		return fmt.Errorf("cannot insert document into collection: %w", err)
	}

	fmt.Println(insertResult)

	return nil
}

func (cm *CardModel) GetAllCardsFromSet(setName string) ([]Card, error) {
	coll := cm.DB.Collection(setName)
	var cards []Card

	cursor, err := coll.Find(cm.Ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(cm.Ctx, &cards); err != nil {
		return nil, err
	}

	return cards, nil
}

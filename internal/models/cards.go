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

	_, err := coll.InsertOne(
		cm.Ctx,
		c,
	)

	if err != nil {
		return fmt.Errorf("cannot insert document into collection: %w", err)
	}

	return nil
}

func (cm *CardModel) GetAllCardsFromSet(setName string) ([]Card, error) {
	coll := cm.DB.Collection(setName)
	var cards []Card

	cursor, err := coll.Find(cm.Ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("cannot get documents: %w", err)
	}

	if err := cursor.All(cm.Ctx, &cards); err != nil {
		return nil, fmt.Errorf("cannot unmarshal values: %w", err)
	}

	return cards, nil
}

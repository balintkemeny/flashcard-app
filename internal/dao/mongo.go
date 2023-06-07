package dao

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocumentDBAdapter interface {
	InsertOne(string, interface{}) error
	GetAll(string, interface{}) error
}

type Mongo struct {
	DB  *mongo.Database
	Ctx context.Context
}

func (m *Mongo) InsertOne(collectionName string, v interface{}) error {
	coll := m.DB.Collection(collectionName)

	_, err := coll.InsertOne(
		m.Ctx,
		v,
	)

	if err != nil {
		return fmt.Errorf("cannot insert document into collection: %w", err)
	}

	return nil
}

func (cm *Mongo) GetAll(collectionName string, dest interface{}) error {
	coll := cm.DB.Collection(collectionName)

	cursor, err := coll.Find(cm.Ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("cannot get documents: %w", err)
	}

	if err := cursor.All(cm.Ctx, dest); err != nil {
		return fmt.Errorf("cannot unmarshal values: %w", err)
	}

	return nil
}

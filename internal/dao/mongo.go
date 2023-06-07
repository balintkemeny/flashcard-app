package dao

import (
	"context"
	"errors"
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
		return fmt.Errorf("%w %s: %w", ErrInsertDocument, collectionName, err)
	}

	return nil
}

func (m *Mongo) GetAll(collectionName string, dest interface{}) error {
	coll := m.DB.Collection(collectionName)

	cursor, err := coll.Find(m.Ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("%w %s: %w", ErrGetDocuments, collectionName, err)
	}

	if err := cursor.All(m.Ctx, dest); err != nil {
		return errors.Join(ErrUnmarshal, err)
	}

	return nil
}

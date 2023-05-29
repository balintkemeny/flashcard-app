package models

import "go.mongodb.org/mongo-driver/mongo"

type CardModel struct {
	DB *mongo.Client
}

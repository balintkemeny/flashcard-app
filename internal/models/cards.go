package models

import (
	"github.com/balintkemeny/flashcard-app/internal/dao"
)

type Card struct {
	ID      string   `bson:"_id,omitempty"`
	Content []string `bson:"content"`
	Level   int      `bson:"level"`
}

type CardRepository struct {
	Dao dao.DocumentDBAdapter
}

func (cm *CardRepository) InsertCardIntoSet(c Card, setName string) error {
	return cm.Dao.InsertOne(setName, c)
}

func (cm *CardRepository) GetAllCardsFromSet(setName string) ([]Card, error) {
	cards := []Card{}

	err := cm.Dao.GetAll(setName, &cards)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

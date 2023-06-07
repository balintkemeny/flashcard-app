package models

import (
	"strings"

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

func (r *CardRepository) InsertCard(c Card, userName, cardSetName string) error {
	return r.Dao.InsertOne(getCollectionName(userName, cardSetName), c)
}

func (r *CardRepository) GetAllCards(userName, cardSetName string) ([]Card, error) {
	cards := []Card{}

	err := r.Dao.GetAll(getCollectionName(userName, cardSetName), &cards)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func getCollectionName(userName, cardSetName string) string {
	return strings.Join([]string{userName, cardSetName}, "-")
}

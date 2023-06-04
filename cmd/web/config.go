package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type config struct {
	Port        int    `env:"PORT,default=3000"`
	MongoDBURI  string `env:"MONGODB_URI,required"`
	CardsDBName string `env:"CARDS_DB_NAME,default=cards"`
}

func NewConfig() (*config, error) {
	c := &config{}
	_ = godotenv.Load(".env")

	err := envconfig.Process(context.Background(), c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

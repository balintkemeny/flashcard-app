package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/balintkemeny/flashcard-app/internal/dao"
	"github.com/balintkemeny/flashcard-app/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	cardRepository *models.CardRepository
}

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	c, err := NewConfig()
	if err != nil {
		errorLog.Fatalf("cannot create config: %s", err)
	}

	addr := ":" + strconv.Itoa(c.Port)
	ctx := context.Background()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(c.MongoDBURI))
	if err != nil {
		errorLog.Fatalf("cannot connect to database: %s", err)
	}

	cardDao := &dao.Mongo{DB: mongoClient.Database(c.CardsDBName), Ctx: ctx}

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		cardRepository: &models.CardRepository{Dao: cardDao},
	}

	srv := http.Server{
		Addr:     addr,
		Handler:  app.routes(),
		ErrorLog: app.errorLog,
	}

	infoLog.Printf("starting server at: %s", addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/balintkemeny/flashcard-app/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	cardModel *models.CardModel
}

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	c, err := NewConfig()
	if err != nil {
		errorLog.Fatalf("cannot create config: %s", err)
	}

	addr := ":" + strconv.Itoa(c.Port)

	db, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.MongoDbURI))
	if err != nil {
		errorLog.Fatalf("cannot connect to database: %s", err)
	}

	app := &application{
		errorLog:  errorLog,
		infoLog:   infoLog,
		cardModel: &models.CardModel{DB: db},
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

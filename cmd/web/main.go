package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	c, err := NewConfig()
	if err != nil {
		errorLog.Fatalf("cannot create config: %s", err)
	}

	addr := ":" + strconv.Itoa(c.Port)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
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

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"simple-microservice/homepage"
	"simple-microservice/server"
)

var (
	serverAddress = os.Getenv("SERV_ADDRESS")
)

func main() {
	logger := log.New(os.Stdout, "test", log.LstdFlags|log.Lshortfile)

	db, err := sqlx.Open("postgres", "postgres://postgres:postgres@")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	h := homepage.NewHandlers(logger, db)
	mux := http.NewServeMux()
	h.SetupRoutes(mux)
	srv := server.New(mux, serverAddress)

	logger.Println("server starting")
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}

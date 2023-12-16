package main

import (
	"go-migrations/config"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	var db = config.ConnectToDB()
	config.ApplyMigrations(db.DB)

	var router = chi.NewRouter()

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"go-migrations/config"
	"go-migrations/todos"
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

	router.Post("/", todos.CreateTodo(db))
	router.Get("/", todos.GetTodos(db))
	router.Get("/{id}", todos.GetTodo(db))
	router.Put("/{id}", todos.GetTodo(db))
	router.Delete("/{id}", todos.GetTodo(db))

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		log.Fatal(err)
	}
}

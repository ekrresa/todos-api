package todos

import (
	"errors"
	"go-migrations/helpers"
	"go-migrations/model"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func CreateTodo(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody model.CreateTodoInput

		var decodeErr = helpers.DecodeJSONBody(w, r.Body, &requestBody)
		if decodeErr != nil {
			var requestError *helpers.RequestError
			if errors.As(decodeErr, &requestError) {
				helpers.ErrorResponse(w, requestError.Error(), requestError.StatusCode)
			} else {
				helpers.ErrorResponse(w, decodeErr.Error(), http.StatusBadRequest)
			}
			return
		}

		var validate = validator.New()
		var validateErr = validate.Struct(&requestBody)
		if validateErr != nil {
			helpers.ErrorResponse(w, validateErr.Error(), http.StatusBadRequest)
			return
		}

		var todo = model.Todo{
			Title:       requestBody.Title,
			Description: requestBody.Description,
			Status:      requestBody.Status,
		}

		if todo.Status == "" {
			todo.Status = "new"
		}

		var insertErr = db.QueryRowx(`
			INSERT INTO todos (title, description, status) 
			VALUES ($1, $2, $3) RETURNING *`, todo.Title, todo.Description, todo.Status).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)

		if insertErr != nil {
			log.Println("Error inserting todo:", insertErr.Error())
			helpers.ErrorResponse(w, "Error inserting todo", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponse(w, todo, "Todo created successfully", http.StatusCreated)
	}
}

func GetTodos(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todos []model.Todo

		var err = db.Select(&todos, "SELECT * FROM todos")
		if err != nil {
			log.Println("Error getting todos:", err.Error())
			helpers.ErrorResponse(w, "Error getting todos", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponse(w, todos, "Todos retrieved successfully", http.StatusOK)
	}
}

func UpdateTodo(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteTodo(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func GetTodo(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

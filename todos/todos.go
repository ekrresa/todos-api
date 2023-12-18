package todos

import (
	"database/sql"
	"errors"
	"go-migrations/helpers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func CreateTodo(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody CreateTodoInput

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

		var todo Todo

		var insertErr = db.Get(&todo, `
			INSERT INTO todos (title, description, status) 
			VALUES ($1, $2, $3) RETURNING *`, requestBody.Title, requestBody.Description, requestBody.Status)

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
		var todos = []Todo{}

		var err = db.Select(&todos, "SELECT * FROM todos")
		if err != nil {
			log.Println("Error getting todos:", err.Error())
			helpers.ErrorResponse(w, "Error getting todos", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponse(w, todos, "Todos retrieved successfully", http.StatusOK)
	}
}

func GetTodo(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todoId = chi.URLParam(r, "id")
		var todo Todo

		var err = db.Get(&todo, "SELECT * FROM todos WHERE id = $1", todoId)
		if err != nil {
			if err == sql.ErrNoRows {
				helpers.ErrorResponse(w, "Todo not found", http.StatusNotFound)
			} else {
				log.Println("Error getting todo:", err.Error())
				helpers.ErrorResponse(w, "Error getting todo", http.StatusInternalServerError)
			}
			return
		}

		helpers.SuccessResponse(w, todo, "Todo retrieved successfully", http.StatusOK)
	}
}

func UpdateTodo(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todoId = chi.URLParam(r, "id")
		var requestBody UpdateTodoInput

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

		var todo Todo

		var err = db.Get(&todo, `UPDATE todos
		SET title = $2, description = $3, status = $4
		WHERE id = $1
		RETURNING *`, todoId, requestBody.Title, requestBody.Description, requestBody.Status)

		if err != nil {
			if err == sql.ErrNoRows {
				helpers.ErrorResponse(w, "Todo not found", http.StatusNotFound)
			} else {
				log.Println("Error updating todo:", err.Error())
				helpers.ErrorResponse(w, "Error updating todo", http.StatusInternalServerError)
			}
			return
		}

		helpers.SuccessResponse(w, todo, "Todo updated successfully", http.StatusOK)
	}
}

func DeleteTodo(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todoId = chi.URLParam(r, "id")

		var _, err = db.Exec("DELETE FROM todos WHERE id = $1", todoId)
		if err != nil {
			log.Println("Error deleting todo:", err.Error())
			helpers.ErrorResponse(w, "Error deleting todo", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponse(w, nil, "Todo deleted successfully", http.StatusOK)
	}
}

package model

import "time"

type TodoStatus string

const (
	New        TodoStatus = "new"
	InProgress TodoStatus = "in_progress"
	Completed  TodoStatus = "completed"
	Cancelled  TodoStatus = "cancelled"
)

type Todo struct {
	ID          int        `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Status      TodoStatus `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

type CreateTodoInput struct {
	Title       string     `json:"title" validate:"required,min=8"`
	Description string     `json:"description,omitempty"`
	Status      TodoStatus `json:"status,omitempty"`
}

type UpdateTodoInput struct {
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      TodoStatus `json:"status,omitempty"`
}

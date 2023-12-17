package todos

import (
	"encoding/json"
	"time"
)

type TodoStatus string

const (
	New        TodoStatus = "new"
	InProgress TodoStatus = "in_progress"
	Completed  TodoStatus = "completed"
	Cancelled  TodoStatus = "cancelled"
)

type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TodoStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

type CreateTodoInput struct {
	Title       string     `json:"title" validate:"required,min=8"`
	Description string     `json:"description,omitempty"`
	Status      TodoStatus `json:"status,omitempty"`
}

func (i *CreateTodoInput) UnmarshalJSON(data []byte) error {
	type Alias CreateTodoInput
	var temp = Alias{
		Status: New,
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*i = CreateTodoInput(temp)

	return nil
}

type UpdateTodoInput struct {
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      TodoStatus `json:"status,omitempty"`
}

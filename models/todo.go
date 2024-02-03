package models

import "github.com/google/uuid"

type Todo struct {
	Guid        string `json:"guid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoStatusUpdate struct {
	Guid      string `json:"guid"`
	Completed bool   `json:"completed"`
}

func NewTodo(title string, description string, completed bool) *Todo {
	return &Todo{
		Guid:        uuid.NewString(),
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

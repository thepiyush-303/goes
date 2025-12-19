package cmd

import (
	"time"
)

type Todo struct {
	// id int
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time // can be nil
}

type Todos []Todo

// SetTodos replaces the package-level todos with the provided value.
func SetTodos(t Todos) {
	todos = t
}

// GetTodos returns the current package-level todos value.
func GetTodos() Todos {
	return todos
}



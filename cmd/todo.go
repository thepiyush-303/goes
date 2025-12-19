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

// func (todos *Todos) add(title string) {
// 	todo := Todo{
// 		Title:       title,
// 		Completed:   false,
// 		CreatedAt:   time.Now(),
// 		CompletedAt: nil,
// 	}
// 	*todos = append(*todos, todo)
// }


// func (todos *Todos) delete(index int) error {
// 	if err := todos.validateIndex(index); err != nil {
// 		return err
// 	}

// 	t := *todos
// 	*todos = append(t[:index], t[index+1:]...)
// 	return nil
// }

func (todos *Todos) toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	t := *todos
	todo := &t[index]

	if todo.Completed {
		todo.CompletedAt = nil
		todo.Completed = false
	} else {
		completedTime := time.Now()
		todo.CompletedAt = &completedTime
		todo.Completed = true
	}

	return nil
}

// SetTodos replaces the package-level todos with the provided value.
func SetTodos(t Todos) {
	todos = t
}

// GetTodos returns the current package-level todos value.
func GetTodos() Todos {
	return todos
}



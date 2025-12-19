/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"todo-cli/cmd"
)


func main() {
	// load todos from storage into a local variable, then set into cmd package
	var loaded cmd.Todos
	storage := NewStorage[cmd.Todos]("todos.json")

	err := storage.Load(&loaded)
	if err != nil {
		fmt.Println("could not able to load")
	}

	// give the loaded todos to the cmd package so subcommands operate on the same data
	cmd.SetTodos(loaded)

	cmd.Execute()

	// save the todos from the cmd package (commands may have mutated them)
	err = storage.Save(cmd.GetTodos())
	if err != nil {
		fmt.Printf("Error saving todos in storage: %v\n", err)
	}
}

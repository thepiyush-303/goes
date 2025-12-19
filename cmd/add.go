/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"
	"github.com/spf13/cobra"
)

// Todos represents a slice of Todo items

var todos Todos
// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task for your day",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todos.add(args[0])
		fmt.Println("task added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func (todos *Todos) add(title string){
	todo := Todo{
		Title : title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

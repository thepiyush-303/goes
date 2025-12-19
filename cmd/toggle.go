/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil{
			fmt.Println("Not a valid Index")
		}
		err = todos.toggle(index)
		if err != nil{
			fmt.Println("Toggle can't work")
		}
		fmt.Println("Task Toggled")
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toggleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toggleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func(todos *Todos) toggle(index int) error{
	if err := todos.validateIndex(index); err != nil{
		return err
	}
	t := *todos
	todo := &t[index]

	if todo.Completed{
		todo.CompletedAt = nil
	} else {
		completed := time.Now()
		todo.CompletedAt = &completed
	}
	todo.Completed = !todo.Completed
	return nil
}
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/som-matrix/todo-cli/todos"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing command")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "list":
		todos.ListTodos()
	case "add":
		if len(os.Args) == 2 {
			fmt.Println("Missing content")
			os.Exit(1)
		}

		content := os.Args[2]
		todos.AddTodo(content)
	case "complete":
		if len(os.Args) == 2 {
			fmt.Println("Missing id")
			os.Exit(1)
		}

		id := os.Args[2]
		idInt, _ := strconv.Atoi(id)
		todos.CompleteTodo(idInt)
	case "delete":
		if len(os.Args) == 2 {
			fmt.Println("Missing id")
			os.Exit(1)
		}

		index := os.Args[2]
		idInt, _ := strconv.Atoi(index)
		todos.DeleteTodo(idInt)
	default:
		fmt.Println("Invalid command")
	}
}

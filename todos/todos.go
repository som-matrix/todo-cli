package todos

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Id          int
	Content     string
	IsCompleted bool
}

var fileName = "todos.json"

func loadTodos() ([]Todo, error) {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return []Todo{}, err
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		return []Todo{}, err
	}

	if len(content) == 0 {
		return []Todo{}, err
	}

	var todos []Todo

	err = json.Unmarshal(content, &todos)
	if err != nil {
		return []Todo{}, err
	}
	return todos, nil
}

func saveTodos(todos []Todo) error {
	content, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, content, 0644)
}

// Read
func ListTodos() {
	todos, err := loadTodos()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, value := range todos {
		fmt.Println(value.Content)
	}

}

// Create
func AddTodo(content string) {
	todos, err := loadTodos()
	if err != nil {
		fmt.Println(err)
		return
	}

	newTodo := Todo{
		Id:          len(todos) + 1,
		Content:     content,
		IsCompleted: false,
	}

	todos = append(todos, newTodo)

	error := saveTodos(todos)
	if error != nil {
		fmt.Println(error)
		return
	}
}

// Update
func CompleteTodo(id int) {
	todos, err := loadTodos()
	if err != nil {
		fmt.Println(err)
		return
	}

	found := false
	for i, todo := range todos {
		if todo.Id == id {
			todos[i].IsCompleted = true
			found = true
		}
	}

	if !found {
		fmt.Println("Todo not found")
		return
	}

	fmt.Println("Todo is marked as completed")

	error := saveTodos(todos)
	if error != nil {
		fmt.Println(error)
		return
	}
}

// Delete
func DeleteTodo(id int) {
	todos, err := loadTodos()
	if err != nil {
		fmt.Println(err)
		return
	}

	found := false
	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			found = true
		}
	}

	if !found {
		fmt.Println("Todo not found")
		return
	}

	fmt.Println("Todo is deleted")

	error := saveTodos(todos)
	if error != nil {
		fmt.Println(error)
		return
	}
}

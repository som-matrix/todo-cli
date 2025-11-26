Creating a go project

1. Create a go module using  

```bash
go mod init github.com/som-matrix/todo-cli
```

Note: Replace `github.com/som-matrix/todo-cli` with your own module name (repository name)

2. Create a main.go file

```bash
touch main.go
```

3. Run the project

```bash
go run main.go
```

4.  Learning how imports work

Create a new folder called todos in the root directory of the project

```bash
mkdir todos
```

Create a new file called todos.go in the todos folder

```bash
touch todos/todos.go
```

Add the following code to todos.go

```go
package todos

import "fmt"

func LoadTodo() {
	fmt.Println("Loading todos...")
}
```

Go to main.go and add the following code

```go
package main

import "github.com/som-matrix/todo-cli/todos"

func main() {
	todos.LoadTodo()
}
```

5. In go a function can return multiple values (amazing!)

6. In go array has a fixed size and it is always defined with the type

```go
var todos [3]string
```

7. In go if we don't specify the size of the array it is treated as a slice and slice is dynamic. 

8. In go we can use append function to add elements to the slice

9. In go we only have `for` loop

10. unordered collection of key value pairs in go is called map

```go
myMap := map[string]string{}

// key is of string
// and value is of string
```

11. strings in go are bytes not characters

12. In go a rune represents a single character and it is an alias for int32 , so rune is basically a int32 generally used to reprsent unicode characters

13. strings uses double quotes and runes uses single quotes

14. Capitalization in GO means they can be exported or accessible from other packages but if it is in lowercase it is not accessible from other packages , hence considered as private

15. Structs as value type and reference(pointer) type

16. Assigning a new struct to a already initialized struct will create a new copy with a new memory address

17. Referencing a struct to a already initialized struct will create a new reference to the same memory address (useful when you want to modify the struct in multiple places, and it saves memory)

18. Functions vs methods in go, methods are functions with a receiver which means it is associated with a type where as functions are not associated with a type type is defined in the function signature
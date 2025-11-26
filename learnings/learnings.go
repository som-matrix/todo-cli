package learnings

import (
	"errors"
	"fmt"
)

// Single return value

func LoadTodo(command string) string {
	return "Loaded all todos"
}

// the above function takes a argument as string and returns a string

// Multi return values

func AddTodo(command string) (string, error) {
	if command == "" {
		return "", errors.New("content is missing")
	}

	return "Todo1", nil
}

//  the above function takes a argument as string and returns a string and an error

// Slice vs Array

func ListTodos() []string {
	return []string{"Todo1", "Todo2", "Todo3", "todo4"}
}

// the above function returns a slice of strings it is not treated as an array , in go
// slice can shrink or grow in size

func ListTodos1() [3]string {
	return [3]string{"Todo1", "Todo2", "Todo3"}
}

// the above function returns an array of strings it is not treated as a slice , in go
// array has a fixed size and it is always defined with the type

// Modify a given slice

func ModifyTodo(content string, todos []string) []string {
	return append(todos, content)
}

// Loops in go

func ClassicForLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func RangeForLoop() {
	for i := range 10 {
		fmt.Println(i)
	}
}

// with index and value
func RangeForLoopWithSlice() {
	todos := []string{"TODO1", "TODO2", "TODO3", "TODO4"}
	for i, todo := range todos {
		fmt.Println(i, todo)
	}
}

// with just value
func RangeForLoopWithSlice2() {
	todos := []string{"TODO1", "TODO2", "TODO3", "TODO4"}
	for _, todo := range todos {
		fmt.Println(todo)
	}
}

// with just index
func RangeForLoopWithSlice3() {
	todos := []string{"TODO1", "TODO2", "TODO3", "TODO4"}
	for i := range todos {
		fmt.Println(i)
	}
}

// Infinite loop
// func InfiniteLoop() {
// 	for {
// 		fmt.Println("Infinite loop")
// 	}
// }

// Maps in go
func MapsInGo() {
	myMap := map[string]string{
		"Todo1": "Buy Groceries",
		"Todo2": "Buy Clothes",
		"Todo3": "Buy Plants",
	}
	fmt.Println(myMap)
	myMap["Todo4"] = "Buy Books"
	fmt.Println(myMap)
}

// delete a key from map
func DeleteKeyFromMap() {
	myMap := map[string]string{
		"Todo1": "Buy Groceries",
		"Todo2": "Buy Clothes",
		"Todo3": "Buy Plants",
	}
	delete(myMap, "Todo1")
	fmt.Println(myMap)
}

// using map in for loop with range

func RangeForLoopWithMap() {
	myMap := map[string]string{
		"Todo1": "Buy Groceries",
		"Todo2": "Buy Clothes",
		"Todo3": "Buy Plants",
	}
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}

// Runes in go

func RunesInGo() {
	r := '🐱'               // --> can only contain a single unicode character
	fmt.Println(r)         // --> will print the unicode value of the character
	fmt.Println(string(r)) // --> will print the character
}

// runes in a for loop

func RunesInForLoop() {
	name := "satya"
	for _, r := range name {
		fmt.Println(r) // --> will print the unicode value of each character
	}
}

// structs in go

type Person struct {
	Name    string
	Age     int
	Email   string
	IsAdult bool
}

func StructsInGo() {
	p := Person{
		Name:    "Satya",
		Age:     22,
		Email:   "satya@example.com",
		IsAdult: true,
	}

	// Access the struct filed using dot operator
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	fmt.Println(p.Email)
	fmt.Println(p.IsAdult)
}

func StructsAsValueTypeInGo() {
	p1 := Person{
		Name:    "Adam",
		Age:     22,
		Email:   "adam@example.com",
		IsAdult: true,
	}

	p2 := p1 // --> p2 is a copy of p1 a whole new memory address

	p2.Name = "John" // only p2 is modified and p1 is unaffected

	fmt.Println(p1.Name)
	fmt.Println(p2.Name)
}

func StructsAsReferenceTypeInGo() {
	p1 := &Person{
		Name:    "Chris",
		Age:     22,
		Email:   "chris@example.com",
		IsAdult: true,
	}

	p2 := p1 // --> p2 points to the same memory address as p1

	p2.Name = "John" // both p1 and p2 are modified
	p2.Age = 23
	p2.Email = "john@example.com"
	p2.IsAdult = false

	fmt.Println(p1.Name)
	fmt.Println(p2.Name)
}

// Methods in go

type Currency int

func (c Currency) FormatUSD() string { // --> receiver is c
	return fmt.Sprintf("$%d.00", c)
}

func MethodsInGo() {
	c := Currency(1000)

	fmt.Println(c.FormatUSD())
}

type Counter struct {
	Count int
}

func (c Counter) IncrementAsValueType() {
	c.Count++
	fmt.Println("Inside Value Method:", c.Count)
}

func (c *Counter) IncrementAsReferenceType() {
	c.Count++
	fmt.Println("Inside Reference Method:", c.Count)
}

func MethodsInStructs() {
	myCounter := Counter{Count: 5}

	myCounter.IncrementAsValueType()
	fmt.Println(myCounter.Count)
	myCounter.IncrementAsReferenceType()
	fmt.Println(myCounter.Count)
}

package terminal_msgs

import "fmt"

func Welcome() string {
	fmt.Println("Welcome to Todo App")
}

func todoMenu() {
	fmt.Println("1. Create Todo")
	fmt.Println("2. Get Todo")
	fmt.Println("3. Get Todos")
	fmt.Println("4. Update Todo")
	fmt.Println("5. Delete Todo")
	fmt.Println("6. Exit")
}

func CreateTodo() {
	fmt.Println("Create Todo")
}

func GetTodo() {
	fmt.Println("Get Todo")
}

func GetTodos() {
	fmt.Println("Get Todos")
}

func UpdateTodo() {
	fmt.Println("Update Todo")
}

func DeleteTodo() {
	fmt.Println("Delete Todo")
}

func Exit() {
	fmt.Println("Exit")
}

func clearScreen() {
	fmt.Println("\033[H\033[2J")
}

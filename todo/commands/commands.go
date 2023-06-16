package commands

import (
	"fmt"
	"log"
	"os"
)

func CreateTodo() {
	log.Println("Create Todo")
	// database.CreateTodo()

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
	log.Println("Are you sure you want to exit? (y/n)")
	var option string

	fmt.Scanln(&option)
	if option == "y" {
		log.Println("Exiting...")
		os.Exit(0)
	}
}

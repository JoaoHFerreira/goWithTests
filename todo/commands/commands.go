package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"todo/database"

	"gorm.io/gorm"
)

func CreateTodo(db *gorm.DB) {
	var todo database.TodoFields
	reader := bufio.NewReader(os.Stdin)
	tempStatus := ""

	fmt.Println("Enter title")
	todo.Title, _ = reader.ReadString('\n')

	fmt.Println("Enter description")
	todo.Description, _ = reader.ReadString('\n')

	for {
		fmt.Println("Enter status. Should be 'TODO', 'WIP' or 'DONE'")
		fmt.Scanln(&tempStatus)
		if tempStatus == "TODO" || tempStatus == "WIP" || tempStatus == "DONE" {
			todo.Status = tempStatus
			break
		}
	}
	database.Insert(db, todo)
}

func GetTodo(db *gorm.DB) {
	var id int

	fmt.Println("Enter ID")
	fmt.Scanln(&id)
	todo := database.Get(db, id)

	fmt.Println("ID\tTitle\tDescription\tStatus\tDue Date")
	fmt.Printf("%d\t%s\t%s\t%s\t%s\n", todo.ID, todo.Title, todo.Description, todo.Status, todo.DueDate)
}

func GetTodos(db *gorm.DB) {
	todo := database.GetAll(db)
	fmt.Println("ID\tTitle\tDescription\tStatus\tDue Date")
	for _, t := range todo {
		fmt.Printf("%d\t%s\t%s\t%s\t%s\n", t.ID, t.Title, t.Description, t.Status, t.DueDate)
	}
}

func UpdateTodo(db *gorm.DB) {
	var id int
	var todo database.TodoFields

	fmt.Println("Enter ID")
	fmt.Scanln(&id)

	reader := bufio.NewReader(os.Stdin)
	tempStatus := ""

	fmt.Println("Enter title")
	todo.Title, _ = reader.ReadString('\n')

	fmt.Println("Enter description")
	todo.Description, _ = reader.ReadString('\n')

	for {
		fmt.Println("Enter status. Should be 'TODO', 'WIP' or 'DONE'")
		fmt.Scanln(&tempStatus)
		if tempStatus == "TODO" || tempStatus == "WIP" || tempStatus == "DONE" {
			todo.Status = tempStatus
			break
		}
	}
	database.Update(db, id, todo)
}

func DeleteTodo(db *gorm.DB) {
	var id int

	fmt.Println("Enter ID to delete")
	fmt.Scanln(&id)
	database.Delete(db, id)
}

func Exit(db *gorm.DB) {
	log.Println("Are you sure you want to exit? (y/n)")
	var option string

	fmt.Scanln(&option)
	if option == "y" {
		log.Println("Exiting...")
		os.Exit(0)
	}
}

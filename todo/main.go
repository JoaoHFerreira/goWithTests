package main

import (
	"fmt"
	"todo/commands"
	"todo/database"

	"gorm.io/gorm"
)

func menu() {
	fmt.Println("1. Create Todo")
	fmt.Println("2. Get Todo")
	fmt.Println("3. Get Todos")
	fmt.Println("4. Update Todo")
	fmt.Println("5. Delete Todo")
	fmt.Println("6. Exit")
	fmt.Println("\n\n\n\n")
}

func terminalMsgs(option int) bool {
	terminalOptions := map[int]string{
		1: "Create Todo",
		2: "Get Todo",
		3: "Get Todos",
		4: "Update Todo",
		5: "Delete Todo",
		6: "Exit",
	}

	todoAction, ok := terminalOptions[option]
	if !ok {
		return false
	}

	fmt.Printf("You chose %s\n", todoAction)
	return true
}

func getTerminalActions() map[int]func(*gorm.DB) {
	return map[int]func(*gorm.DB){
		1: func(db *gorm.DB) { commands.CreateTodo(db) },
		2: func(db *gorm.DB) { commands.GetTodo(db) },
		3: func(db *gorm.DB) { commands.GetTodos(db) },
		4: func(db *gorm.DB) { commands.UpdateTodo(db) },
		5: func(db *gorm.DB) { commands.DeleteTodo(db) },
		6: func(db *gorm.DB) { commands.Exit(db) },
	}
}

func main() {
	database.Migrate()
	db := database.GetDBInstance()
	fmt.Println(db)

	terminalActionMap := getTerminalActions()
	for {
		menu()
		var option int
		fmt.Scanln(&option)

		if !terminalMsgs(option) {
			fmt.Println("Invalid option")
			continue
		}

		action, ok := terminalActionMap[option]
		if !ok {
			fmt.Println("Invalid option")
			continue
		}
		// pass database db variable reference to action method
		action(db)
	}
}

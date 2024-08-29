package main

import (
	"fmt"
	"os"
	"todo/cmd"
	"todo/repository"
	"todo/utils"
)



func Runner() {
	var option int

	fmt.Println("Choose one of these options: ")
	// Category Options
	fmt.Println("1. Create Category")
	fmt.Println("2. List all Categories")
	fmt.Println("3. Get Category")
	fmt.Println("4. Edit Category")
	fmt.Println("5. Delete Category")
	// Task Options
	fmt.Println("6. Create Task")
	fmt.Println("7. Edit Task")
	fmt.Println("8. Delete Task")
	fmt.Println("10. List all Tasks")
	fmt.Println("9. List User Tasks")
	fmt.Println("10. Change Task Status")
	// User Options
	fmt.Println("11. Register User")
	fmt.Println("12. Login User")
	// Exit option
	fmt.Println("0. Exit")

	fmt.Scan(&option)

	handleOption(option)
}

func auth(next func()) {
	auth := utils.Auth{}
	isAuthenticated := auth.IsValidToken()
	if !isAuthenticated {
		fmt.Println("You need to be logged in to use this option")
		return
	}
	next()
}

func handleOption(option int) {
	categoriesPath := "./data/categories.json"
	categoryRepo := repository.NewCategoryRepository(categoriesPath)
	
	cmd := cmd.NewCategoryCommand(categoryRepo)
	switch(option) {
	case 0:
		fmt.Println("Goodbye")
		os.Exit(0)
	case 1:
		auth(cmd.CreateCategory)
	case 2:
		auth(cmd.ListAllCategories)
	case 3:
		auth(cmd.GetCategory)
	case 4:
		auth(cmd.EditCategory)
	case 5:
		auth(cmd.DeleteCategory)
	// case 6:
	// 	auth(cmd.CreateTask)
	// case 7:
	// 	auth(cmd.EditTask)
	// case 8:
	// 	auth(cmd.DeleteTask)
	// case 9:
	// 	auth(cmd.ListUserTasks)
	// case 10:
	// 	auth(cmd.ChangeTaskStatus)
	// case 11:
	// 	auth(cmd.RegisterUser)
	// case 12:
	// 	auth(cmd.LoginUser)
	default:
		fmt.Println("Invalid Option")
	}
}
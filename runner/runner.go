package runner

import (
	"fmt"
	"os"
	"todo/cmd"
	"todo/repository"
)


type TodoApp struct {
	CategoryCommand cmd.CategoryCommand
	TaskCommand cmd.TaskCommand
	UserCommand cmd.UserCommand
}

func NewTodoApp(
	categoryCommand cmd.CategoryCommand,
	taskCommand cmd.TaskCommand,
	userCommand cmd.UserCommand,
) TodoApp {
	return TodoApp{
		categoryCommand,
		taskCommand,
		userCommand,
	}
}

var todoApp TodoApp = NewTodoApp(
	cmd.NewCategoryCommand(repository.NewCategoryRepository("./data/categories.json")),
	cmd.NewTaskCommand(repository.NewTaskRepository("./data/tasks.json")),
	cmd.NewUserCommand(repository.NewUserRepository("./data/users.json")),
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
	fmt.Println("9. List User Tasks")
	fmt.Println("10. List all Tasks")
	fmt.Println("11. Change Task Status")
	// User Options
	fmt.Println("12. Register User")
	fmt.Println("13. Login User")
	// Exit option
	fmt.Println("0. Exit")

	fmt.Scan(&option)

	handleOption(option)
}

func auth(next func()) {
	isAuthenticated := todoApp.UserCommand.IsAuthenticated()
	if !isAuthenticated {
		fmt.Println("You need to be logged in to use this option")
		return
	}
	next()
}

func handleOption(option int) {
	switch(option) {
	case 0:
		fmt.Println("Goodbye")
		os.Exit(0)
	case 1:
		auth(todoApp.CategoryCommand.CreateCategory)
	case 2:
		auth(todoApp.CategoryCommand.ListAllCategories)
	case 3:
		auth(todoApp.CategoryCommand.GetCategory)
	case 4:
		auth(todoApp.CategoryCommand.EditCategory)
	case 5:
		auth(todoApp.CategoryCommand.DeleteCategory)
	case 6:
		auth(todoApp.TaskCommand.CreateTask)
	case 7:
		auth(todoApp.TaskCommand.EditTask)
	case 8:
		auth(todoApp.TaskCommand.DeleteTask)
	case 9:
		auth(todoApp.TaskCommand.ListUserTasks)
	case 10:
		auth(todoApp.TaskCommand.ListAllTasks)
	case 11:
		auth(todoApp.TaskCommand.ChangeTaskStatus)
	case 12:
		todoApp.UserCommand.RegisterUser()
	case 13:
		todoApp.UserCommand.LoginUser()
	default:
		fmt.Println("Invalid Option")
	}
}
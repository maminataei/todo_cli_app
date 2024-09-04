package cmd

import (
	"fmt"
	"math/rand/v2"
	"todo/interfaces"
	"todo/model"
	"todo/repository"
)

type UserCommand struct {
	repo interfaces.Repo[model.User]
}

func NewUserCommand(repo interfaces.Repo[model.User]) UserCommand {
	return UserCommand{repo}
}

func (cmd *UserCommand) RegisterUser() {
	fmt.Println("Register User Command ...")
	var Email, Password string

	fmt.Println("Please enter the email of the user : ")
	fmt.Scanln(&Email)

	fmt.Println("Please enter the user password : ")
	fmt.Scanln(&Password)

	if err := cmd.repo.Create(model.User{Id: rand.IntN(100), Email: Email, Password: Password}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User registered successfully")
}
func (cmd *UserCommand) LoginUser() {
	fmt.Println("Login User Command ...")
	var Email, Password string

	fmt.Println("Please enter the email of the user : ")
	fmt.Scanln(&Email)

	fmt.Println("Please enter the user password : ")
	fmt.Scanln(&Password)
	userRepo, ok := cmd.repo.(*repository.UserRepository)
    if !ok {
        fmt.Println("cmd.repo is not a *repository.UserRepository")
        return
    }
	if err := userRepo.LoginUser(Email, Password); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User logged in successfully")
}
func (cmd UserCommand) IsAuthenticated() bool {
	userRepo, ok := cmd.repo.(*repository.UserRepository)
    if !ok {
        fmt.Println("cmd.repo is not a *repository.UserRepository")
        return false
    }
	return userRepo.IsLoggedIn()
}
package cmd

import (
	"fmt"
	"math/rand/v2"
	"todo/model"
	"todo/repository"
)

type UserCommand struct {
	repo repository.UserRepository
}

func NewUserCommand(repo repository.UserRepository) UserCommand {
	return UserCommand{repo}
}

func (cmd *UserCommand) RegisterUser() {
	fmt.Println("Register User Command ...")
	var Email, Password string

	fmt.Println("Please enter the email of the user : ")
	fmt.Scanln(&Email)

	fmt.Println("Please enter the user password : ")
	fmt.Scanln(&Password)

	err := cmd.repo.CreateUser(model.User{Id: rand.IntN(100), Email: Email, Password: Password})
	if err != nil {
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

	err := cmd.repo.LoginUser(Email, Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User logged in successfully")
}
func (cmd UserCommand) IsAuthenticated() bool {
	return cmd.repo.IsLoggedIn()
}
package cmd

import (
	"fmt"
	"todo/model"
	"todo/repository"
)

type UserCommand struct {
	repo repository.UserRepository
}

func (cmd UserCommand) RegisterUser() {
	fmt.Println("Register User Command ...")
	var Email, Password string

	fmt.Println("Please enter the email of the user : ")
	fmt.Scanln(&Email)

	fmt.Println("Please enter the user password : ")
	fmt.Scanln(&Password)

	err := cmd.repo.CreateUser(model.User{Email: Email, Password: Password})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User registered successfully")
}
func (cmd UserCommand) LoginUser() {
	fmt.Println("Login User Command ...")
	var Email, Password string

	fmt.Println("Please enter the email of the user : ")
	fmt.Scanln(&Email)

	fmt.Println("Please enter the user password : ")
	fmt.Scanln(&Password)

	fmt.Println("User logged in successfully")
}
func (cmd UserCommand) IsAuthenticated() bool {
	return true
}
package main

import (
	"fmt"
	"todo/runner"
)

func main() {
	fmt.Println("Welcome to the Todo App")
	for {
		runner.Runner()
	}
}
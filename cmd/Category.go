package cmd

import (
	"fmt"
	"math/rand/v2"
	"todo/model"
	"todo/repository"
)

type CategoryCommand struct {
	repo repository.CategoryRepository
}


func NewCategoryCommand(repo repository.CategoryRepository) CategoryCommand {
	return CategoryCommand{repo}
}

func (cmd CategoryCommand) CreateCategory() {
	var category model.Category = model.Category{}
	fmt.Println("Create Category Command ...")
	
	fmt.Println("Please enter the title of the category : ")
	fmt.Scanln(&category.Title)

	fmt.Println("Please enter the color of the category :")
	fmt.Scanln(&category.Color)
	category.Id = rand.IntN(100)

	err := cmd.repo.CreateCategory(category)
	if err != nil {
		fmt.Println(err)
	} 
	fmt.Println("Category created successfully")
}
func (cmd CategoryCommand) ListAllCategories() {
	fmt.Println("List All Categories Command ...")
	categories, err := cmd.repo.ListAllCategories()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("List of all categories : ")
	fmt.Println(categories)
}
func (cmd CategoryCommand) GetCategory() {
	fmt.Println("Get Category Command ...")
	var categoryId int
	fmt.Println("Please enter the id of the category : ")
	fmt.Scanf("%d", &categoryId)
	category, err := cmd.repo.GetCategory(categoryId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Category details : ")
	fmt.Println(category)
}
func (cmd CategoryCommand) EditCategory() {
	fmt.Println("Edit Category Command ...")
	category := model.Category{}
	fmt.Println("Please enter the id of the category : ")
	fmt.Scanf("%d", &category.Id)

	fmt.Println("Now enter the title of category : ")
	fmt.Scanln(&category.Title)

	fmt.Println("Now enter the color of category :")
	fmt.Scanln(&category.Color)

	err := cmd.repo.EditCategory(category)	
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Category edited successfully")
}
func (cmd CategoryCommand) DeleteCategory() {
	fmt.Println("Delete Category Command ...")
	fmt.Println("Please enter the id of the category : ")
	var categoryId int
	fmt.Scanf("%d", &categoryId)

	err := cmd.repo.DeleteCategory(categoryId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Category deleted successfully")
}
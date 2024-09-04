package cmd

import (
	"fmt"
	"math/rand/v2"
	"todo/interfaces"
	"todo/model"
	"todo/utilities"
)

type CategoryCommand struct {
	repo interfaces.Repo[model.Category]
	ioUtil utilities.IO
}
func NewCategoryCommand(repo interfaces.Repo[model.Category]) CategoryCommand {
	return CategoryCommand{repo: repo, ioUtil: utilities.NewIO()}
}
func (cmd *CategoryCommand) Create() {
	var category model.Category = model.Category{}
	fmt.Println("Create Category Command ...")
	
	fmt.Println("Please enter the title of the category : ")
	categoryTitle, readCatTitleErr := cmd.ioUtil.ReadStr()
	if(readCatTitleErr != nil) {
		fmt.Println(readCatTitleErr)
		return
	}
	category.Title = categoryTitle

	fmt.Println("Please enter the color of the category :")
	categoryColor, readCatColorErr := cmd.ioUtil.ReadStr()
	if(readCatColorErr != nil) {
		fmt.Println(readCatColorErr)
		return
	}
	category.Color = categoryColor
	category.Id = rand.IntN(100)

	err := cmd.repo.Create(category)
	if err != nil {
		fmt.Println(err)
	} 
	fmt.Println("Category created successfully")
}
func (cmd *CategoryCommand) ListAll() {
	fmt.Println("List All Categories Command ...")
	categories, err := cmd.repo.ListAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("List of all categories : ")
	fmt.Println(categories)
}
func (cmd *CategoryCommand) Get() {
	fmt.Println("Get Category Command ...")
	fmt.Println("Please enter the id of the category : ")
	categoryId, readCatIdErr := cmd.ioUtil.ReadNumber()
	if(readCatIdErr != nil) {
		fmt.Println(readCatIdErr)
		return
	}
	category, getCatErr := cmd.repo.Get(categoryId)
	if getCatErr != nil {
		fmt.Println(getCatErr)
		return
	}
	fmt.Println("Category details : ")
	fmt.Println(category)
}
func (cmd *CategoryCommand) Edit() {
	fmt.Println("Edit Category Command ...")
	category := model.Category{}
	fmt.Println("Please enter the id of the category : ")
	categoryId, readCatIdErr := cmd.ioUtil.ReadNumber()
	if readCatIdErr != nil {
		fmt.Println(readCatIdErr)
		return
	}
	category.Id = categoryId
	fmt.Println("Now enter the title of category : ")
	categoryTitle, readCatTitleErr := cmd.ioUtil.ReadStr()
	if readCatTitleErr != nil {
		fmt.Println(readCatTitleErr)
		return
	}
	category.Title = categoryTitle
	fmt.Println("Now enter the color of category :")
	categoryColor, readCatColorErr := cmd.ioUtil.ReadStr()
	if readCatColorErr != nil {
		fmt.Println(readCatColorErr)
		return
	}
	category.Color = categoryColor
	err := cmd.repo.Edit(category)	
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Category edited successfully")
}
func (cmd *CategoryCommand) Delete() {
	fmt.Println("Delete Category Command ...")
	fmt.Println("Please enter the id of the category : ")
	categoryId, readCatIdErr := cmd.ioUtil.ReadNumber()
	if readCatIdErr != nil {
		fmt.Println(readCatIdErr)
		return
	}
	err := cmd.repo.Delete(categoryId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Category deleted successfully")
}
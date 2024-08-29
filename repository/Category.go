package repository

import (
	"encoding/json"
	"errors"
	"todo/model"
	"todo/utils"
)

type CategoryRepository struct {
	filePath string
}

var file utils.File = utils.File{}

func NewCategoryRepository(path string) CategoryRepository {
	return CategoryRepository{filePath: path}
}
func (repo CategoryRepository) CreateCategory(cat model.Category) error {
	categoriesStr, readJSONFileErr := file.Read(repo.filePath)
	if(readJSONFileErr != nil) {
		return errors.New("Error reading file")
	}
	var categories []model.Category
	jsonUnmarshalErr := json.Unmarshal([]byte(categoriesStr), &categories)
	if(jsonUnmarshalErr != nil) {
		return errors.New("Error unmarshalling")
	}
	categories = append(categories, cat)
	categoriesJson, categoriesJSONErr := json.Marshal(categories)
	if(categoriesJSONErr != nil) {
		return errors.New("Error marshalling")
	}
	saveJSONErr := file.Save(repo.filePath, string(categoriesJson))
	if(saveJSONErr != nil) {
		return errors.New("Error saving file")
	}
	return nil
}
func (repo CategoryRepository) ListAllCategories() ([]model.Category, error) {
	categoriesStr, readJSONFileErr := file.Read(repo.filePath)
	if readJSONFileErr != nil {
		return []model.Category{}, errors.New("Error reading file")
	}
	var categories []model.Category
	jsonUnmarshalErr := json.Unmarshal([]byte(categoriesStr), &categories)
	if jsonUnmarshalErr != nil {
		return []model.Category{}, errors.New("Error reading file")
	}
	return categories, nil
}
func (repo CategoryRepository) GetCategory(id int) (model.Category, error) {
	categories, err := repo.ListAllCategories()
	if err != nil {
		return model.Category{}, err
	}
	findCategory := model.Category{}
	for _, category := range categories {
		if category.Id == id {
			findCategory = category
			return findCategory, nil
		}
	}
	return model.Category{}, nil
}
func (repo CategoryRepository) EditCategory(cat model.Category) error {
	categories, err := repo.ListAllCategories()
	if err != nil {
		return err
	}
	for i, category := range categories {
		if category.Id == cat.Id {
			categories[i] = model.Category{
				Id: cat.Id,
				Title: cat.Title,
				Color: cat.Color,
			}
			categoriesJson, categoriesJSONErr := json.Marshal(categories)
			if(categoriesJSONErr != nil) {
				return errors.New("Error marshalling")
			}
			saveJSONErr := file.Save(repo.filePath, string(categoriesJson))
			if(saveJSONErr != nil) {
				return errors.New("Error saving file")
			}
			return nil
		}
	}
	return errors.New("Category not found")
}
func (repo CategoryRepository) DeleteCategory(id int) error {
	categories, err := repo.ListAllCategories()
	if err != nil {
		return err
	}
	for i, category := range categories {
		if category.Id == id {
			categories = append(categories[:i], categories[i+1:]...)
			categoriesJson, categoriesJSONErr := json.Marshal(categories)
			if(categoriesJSONErr != nil) {
				return errors.New("Error marshalling")
			}
			saveJSONErr := file.Save(repo.filePath, string(categoriesJson))
			if(saveJSONErr != nil) {
				return errors.New("Error saving file")
			}
			return nil
		}
	}
	return errors.New("Category not found")
}
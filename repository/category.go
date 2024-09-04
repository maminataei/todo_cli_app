package repository

import (
	"encoding/json"
	"errors"
	"todo/interfaces"
	"todo/model"
	"todo/utilities"
)

type CategoryRepository struct {
	filePath string
	fileUtil utilities.File
}

func NewCategoryRepository(path string) interfaces.Repo[model.Category] {
	return &CategoryRepository{filePath: path, fileUtil: utilities.File{}}
}
func (repo *CategoryRepository) Create(cat model.Category) error {
	categoriesStr, readJSONFileErr := repo.fileUtil.Read(repo.filePath)
	if(readJSONFileErr != nil) {
		return errors.New("error reading file")
	}
	var categories []model.Category
	if jsonUnmarshalErr := json.Unmarshal([]byte(categoriesStr), &categories); jsonUnmarshalErr != nil {
		return errors.New("error unmarshalling")
	}
	categories = append(categories, cat)
	categoriesJson, categoriesJSONErr := json.Marshal(categories)
	if(categoriesJSONErr != nil) {
		return errors.New("error marshalling")
	}
	if saveJSONErr := repo.fileUtil.Save(repo.filePath, string(categoriesJson)); saveJSONErr != nil {
		return errors.New("error saving file: " + saveJSONErr.Error())
	}
	return nil
}
func (repo *CategoryRepository) ListAll() ([]model.Category, error) {
	categoriesStr, readJSONFileErr := repo.fileUtil.Read(repo.filePath)
	if readJSONFileErr != nil {
		return []model.Category{}, errors.New("error reading file")
	}
	var categories []model.Category
	if jsonUnmarshalErr := json.Unmarshal([]byte(categoriesStr), &categories); jsonUnmarshalErr != nil {
		return []model.Category{}, errors.New("error reading file: " + jsonUnmarshalErr.Error())
	}
	return categories, nil
}
func (repo *CategoryRepository) Get(id int) (model.Category, error) {
	categories, err := repo.ListAll()
	if err != nil {
		return model.Category{}, err
	}
	for _, category := range categories {
		if category.Id == id {
			return category, nil
		}
	}
	return model.Category{}, nil
}
func (repo *CategoryRepository) Edit(cat model.Category) error {
	categories, err := repo.ListAll()
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
				return errors.New("error marshalling")
			}
			saveJSONErr := repo.fileUtil.Save(repo.filePath, string(categoriesJson))
			if(saveJSONErr != nil) {
				return errors.New("error saving file")
			}
			return nil
		}
	}
	return errors.New("category not found")
}
func (repo *CategoryRepository) Delete(id int) error {
	categories, err := repo.ListAll()
	if err != nil {
		return err
	}
	for i, category := range categories {
		if category.Id == id {
			categories = append(categories[:i], categories[i+1:]...)
			categoriesJson, categoriesJSONErr := json.Marshal(categories)
			if(categoriesJSONErr != nil) {
				return errors.New("error marshalling")
			}
			if saveJSONErr := repo.fileUtil.Save(repo.filePath, string(categoriesJson)); saveJSONErr != nil {
				return errors.New("error saving file: " + saveJSONErr.Error())
			}
			return nil
		}
	}
	return errors.New("category not found")
}
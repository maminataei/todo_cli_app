package repository

import (
	"encoding/json"
	"errors"
	"todo/model"
	"todo/utils"
)

type UserRepository struct {
	filePath string
}

var file utils.File = utils.File{}

func NewUserRepository(path string) UserRepository {
	return UserRepository{filePath: path}
}

func (repo UserRepository) CreateUser(user model.User) error {
	usersStr, readJSONFileErr := file.Read(repo.filePath)
	if(readJSONFileErr != nil) {
		return errors.New("Error reading file")
	}
	var users []model.User
	jsonUnmarshalErr := json.Unmarshal([]byte(usersStr), &users)
	if(jsonUnmarshalErr != nil) {
		return errors.New("Error unmarshalling")
	}
	users = append(users, user)
	usersJson, usersJSONErr := json.Marshal(users)
	if(usersJSONErr != nil) {
		return errors.New("Error marshalling")
	}
	saveJSONErr := file.Save(repo.filePath, string(usersJson))
	if(saveJSONErr != nil) {
		return errors.New("Error saving file")
	}
	return nil
}

func (repo UserRepository) ListAllUsers() ([]model.User, error) {
	usersStr, readJSONFileErr := file.Read(repo.filePath)
	if readJSONFileErr != nil {
		return []model.User{}, errors.New("Error reading file")
	}
	var users []model.User
	jsonUnmarshalErr := json.Unmarshal([]byte(usersStr), &users)
	if jsonUnmarshalErr != nil {
		return []model.User{}, errors.New("Error reading file")
	}
	return users, nil
}

func (repo UserRepository) GetUser(id int) (model.User, error) {
	users, err := repo.ListAllUsers()
	if err != nil {
		return model.User{}, err
	}
	findUser := model.User{}
	for _, user := range users {
		if user.Id == id {
			findUser = user
			return findUser, nil
		}
	}
	return model.User{}, nil
}

func (repo UserRepository) EditUser(cat model.User) error {
	users, err := repo.ListAllUsers()
	if err != nil {
		return err
	}
	for i, user := range users {
		if user.Id == cat.Id {
			users[i] = model.User{
				Id: cat.Id,
				Email: cat.Email,
				Password: cat.Password,
			}
			usersJson, usersJSONErr := json.Marshal(users)
			if(usersJSONErr != nil) {
				return errors.New("Error marshalling")
			}
			saveJSONErr := file.Save(repo.filePath, string(usersJson))
			if(saveJSONErr != nil) {
				return errors.New("Error saving file")
			}
			return nil
		}
	}
	return errors.New("User not found")
}
func (repo UserRepository) DeleteUser(id int) error {
	users, err := repo.ListAllUsers()
	if err != nil {
		return err
	}
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			usersJson, usersJSONErr := json.Marshal(users)
			if(usersJSONErr != nil) {
				return errors.New("Error marshalling")
			}
			saveJSONErr := file.Save(repo.filePath, string(usersJson))
			if(saveJSONErr != nil) {
				return errors.New("Error saving file")
			}
			return nil
		}
	}
	return errors.New("User not found")
}
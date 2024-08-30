package repository

import (
	"encoding/json"
	"errors"
	"todo/model"
	"todo/utilities"
)

type UserRepository struct {
	FilePath string
	FileUtil utilities.File
	loggedIn bool
}
func NewUserRepository(path string) UserRepository {
	return UserRepository{FilePath: path, FileUtil: utilities.File{}, loggedIn: false}
}
func (repo *UserRepository) CreateUser(user model.User) error {
	usersStr, readJSONFileErr := repo.FileUtil.Read(repo.FilePath)
	if(readJSONFileErr != nil) {
		return errors.New("error reading file")
	}
	var users []model.User
	jsonUnmarshalErr := json.Unmarshal([]byte(usersStr), &users)
	if(jsonUnmarshalErr != nil) {
		return errors.New("error unmarshalling")
	}
	users = append(users, user)
	usersJson, usersJSONErr := json.Marshal(users)
	if(usersJSONErr != nil) {
		return errors.New("error marshalling")
	}
	saveJSONErr := repo.FileUtil.Save(repo.FilePath, string(usersJson))
	if(saveJSONErr != nil) {
		return errors.New("error saving file")
	}
	return nil
}
func (repo *UserRepository) ListAllUsers() ([]model.User, error) {
	usersStr, readJSONFileErr := repo.FileUtil.Read(repo.FilePath)
	if readJSONFileErr != nil {
		return []model.User{}, errors.New("error reading file")
	}
	var users []model.User
	jsonUnmarshalErr := json.Unmarshal([]byte(usersStr), &users)
	if jsonUnmarshalErr != nil {
		return []model.User{}, errors.New("error reading file")
	}
	return users, nil
}

func (repo *UserRepository) GetUser(id int) (model.User, error) {
	users, err := repo.ListAllUsers()
	if err != nil {
		return model.User{}, err
	}
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}
	return model.User{}, nil
}
func (repo *UserRepository) GetUserByEmailAndPassword(email string, password string) (model.User, error) {
	users, err := repo.ListAllUsers()
	if err != nil {
		return model.User{}, err
	}
	for _, user := range users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return model.User{}, nil
}
func (repo *UserRepository) EditUser(cat model.User) error {
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
				return errors.New("error marshalling")
			}
			saveJSONErr := repo.FileUtil.Save(repo.FilePath, string(usersJson))
			if(saveJSONErr != nil) {
				return errors.New("error saving file")
			}
			return nil
		}
	}
	return errors.New("user not found")
}
func (repo *UserRepository) DeleteUser(id int) error {
	users, err := repo.ListAllUsers()
	if err != nil {
		return err
	}
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			usersJson, usersJSONErr := json.Marshal(users)
			if(usersJSONErr != nil) {
				return errors.New("error marshalling")
			}
			saveJSONErr := repo.FileUtil.Save(repo.FilePath, string(usersJson))
			if(saveJSONErr != nil) {
				return errors.New("error saving file")
			}
			return nil
		}
	}
	return errors.New("user not found")
}
func (repo *UserRepository) LoginUser(email string, password string) error {
	_, err := repo.GetUserByEmailAndPassword(email, password)
	if err != nil {
		repo.loggedIn = false
		return err
	}
	repo.loggedIn = true
	return nil
}
func (repo *UserRepository) LogoutUser() {
	repo.loggedIn = false
}
func (repo UserRepository) IsLoggedIn() bool {
	return repo.loggedIn
}
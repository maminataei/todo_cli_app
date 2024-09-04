package repository

import (
	"encoding/json"
	"errors"
	"todo/interfaces"
	"todo/model"
	"todo/utilities"
)

type TaskRepository struct {
	filePath string
	fileUtil utilities.File
}


func NewTaskRepository(path string) interfaces.Repo[model.Task] {
	return &TaskRepository{filePath: path, fileUtil: utilities.File{}}
}
func (repo *TaskRepository) Create(task model.Task) error {
	tasksStr, readJSONFileErr := repo.fileUtil.Read(repo.filePath)
	if(readJSONFileErr != nil) {
		return errors.New("error reading file")
	}
	var tasks []model.Task
	jsonUnmarshalErr := json.Unmarshal([]byte(tasksStr), &tasks)
	if(jsonUnmarshalErr != nil) {
		return errors.New("error unmarshalling")
	}
	tasks = append(tasks, task)
	tasksJson, tasksJSONErr := json.Marshal(tasks)
	if(tasksJSONErr != nil) {
		return errors.New("error marshalling")
	}
	saveJSONErr := repo.fileUtil.Save(repo.filePath, string(tasksJson))
	if(saveJSONErr != nil) {
		return errors.New("error saving file")
	}
	return nil
}
func (repo *TaskRepository) ListAll() ([]model.Task, error) {
	tasksStr, readJSONFileErr := repo.fileUtil.Read(repo.filePath)
	if readJSONFileErr != nil {
		return []model.Task{}, errors.New("error reading file")
	}
	var tasks []model.Task
	jsonUnmarshalErr := json.Unmarshal([]byte(tasksStr), &tasks)
	if jsonUnmarshalErr != nil {
		return []model.Task{}, errors.New("error reading file")
	}
	return tasks, nil
}
func (repo *TaskRepository) Get(id int) (model.Task, error) {
	tasks, err := repo.ListAll()
	if err != nil {
		return model.Task{}, err
	}
	for _, task := range tasks {
		if task.Id == id {
			return task, nil
		}
	}
	return model.Task{}, nil
}
func (repo *TaskRepository) Edit(t model.Task) error {
	tasks, err := repo.ListAll()
	if err != nil {
		return err
	}
	for i, task := range tasks {
		if task.Id == t.Id {
			tasks[i] = model.Task{
				Id: t.Id,
				Title: t.Title,
				DueDate: t.DueDate,
				Category: t.Category,
				IsDone: t.IsDone || tasks[i].IsDone,
				UserId: t.UserId | tasks[i].UserId,
			}
			tasksJson, tasksJSONErr := json.Marshal(tasks)
			if(tasksJSONErr != nil) {
				return errors.New("error marshalling")
			}
			saveJSONErr := repo.fileUtil.Save(repo.filePath, string(tasksJson))
			if(saveJSONErr != nil) {
				return errors.New("error saving file")
			}
			return nil
		}
	}
	return errors.New("task not found")
}
func (repo *TaskRepository) Delete(id int) error {
	tasks, err := repo.ListAll()
	if err != nil {
		return err
	}
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			tasksJson, tasksJSONErr := json.Marshal(tasks)
			if(tasksJSONErr != nil) {
				return errors.New("error marshalling")
			}
			saveJSONErr := repo.fileUtil.Save(repo.filePath, string(tasksJson))
			if(saveJSONErr != nil) {
				return errors.New("error saving file")
			}
			return nil
		}
	}
	return errors.New("task not found")
}
func (repo *TaskRepository) ListUserTasks(userId int) ([]model.Task, error) {
	tasks, err := repo.ListAll()
	if err != nil {
		return []model.Task{}, err
	}
	var userTasks []model.Task
	for _, task := range tasks {
		if task.UserId == userId {
			userTasks = append(userTasks, task)
		}
	}
	return userTasks, nil
}
func (repo *TaskRepository) ChangeTaskStatus(taskId int, status bool) error {
	tasks, err := repo.ListAll()
	if err != nil {
		return err
	}
	for i, task := range tasks {
		if task.Id == taskId {
			tasks[i].IsDone = status
			tasksJSON, tasksJSONErr := json.Marshal(tasks)
			if(tasksJSONErr != nil) {
				return errors.New("error marshalling")
			}
			saveJSONErr := repo.fileUtil.Save(repo.filePath, string(tasksJSON))
			if(saveJSONErr != nil) {
				return errors.New("error saving file")
			}
			return nil
		}
	}
	return errors.New("task not found")
}
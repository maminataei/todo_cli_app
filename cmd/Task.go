package cmd

import (
	"fmt"
	"todo/model"
	"todo/repository"
	"todo/utilities"
)

type TaskCommand struct {
	repo repository.TaskRepository
	ioUtil utilities.IO
}

func NewTaskCommand(repo repository.TaskRepository) TaskCommand {
	return TaskCommand{repo: repo, ioUtil: utilities.NewIO()}
}

func (cmd *TaskCommand) CreateTask() {
	fmt.Println("Create Task Command ...")
	var task model.Task
	var err error
	fmt.Print("Please enter the title of the task : ")
	task.Title, err = cmd.ioUtil.ReadStr()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Please enter the due date of the task : ")
	task.DueDate, err = cmd.ioUtil.ReadStr()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Please enter the category of the task : ")
	task.Category, err = cmd.ioUtil.ReadStr()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Please enter the user id of the task : ")
	task.UserId, err = cmd.ioUtil.ReadNumber()
	if err != nil {
		fmt.Println(err)
		return
	}

	task.IsDone = false

	createTaskErr := cmd.repo.CreateTask(task)
	if createTaskErr != nil {
		fmt.Println(createTaskErr)
	}
	fmt.Println("Task created successfully")
}
func (cmd *TaskCommand) ListAllTasks() {
	fmt.Println("List All Tasks Command ...")
	tasks, err := cmd.repo.ListAllTasks()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("List of all tasks : ")
	fmt.Println(tasks)
}
func (cmd *TaskCommand) GetTask() {
	fmt.Println("Get Task Command ...")
	fmt.Println("Please enter the id of the task : ")
	taskId, err := cmd.ioUtil.ReadNumber()
	if err != nil {
		fmt.Println(err)
		return
	}
	task, err := cmd.repo.GetTask(taskId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Task details : ")
	fmt.Println(task)
}
func (cmd *TaskCommand) EditTask() {
	fmt.Println("Edit Task Command ...")
	task := model.Task{}
	fmt.Println("Please enter the id of the task : ")
	var err error
	task.Id, err = cmd.ioUtil.ReadNumber()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Please enter the title of the task : ")
	task.Title, err = cmd.ioUtil.ReadStr()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Please enter the due date of the task : ")
	task.DueDate, err = cmd.ioUtil.ReadStr()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Please enter the category of the task : ")
	task.Category, err = cmd.ioUtil.ReadStr()
	if err != nil {
		fmt.Println(err)
		return
	}

	editTaskError := cmd.repo.EditTask(task)
	if editTaskError != nil {
		fmt.Println(editTaskError)
		return
	}
	fmt.Println("Task edited successfully")
}
func (cmd *TaskCommand) DeleteTask() {
	fmt.Println("Delete Task Command ...")
	fmt.Println("Please enter the id of the task : ")
	taskId, err := cmd.ioUtil.ReadNumber()
	if err != nil {
		fmt.Println(err)
		return
	}

	deleteTaskErr := cmd.repo.DeleteTask(taskId)
	if deleteTaskErr != nil {
		fmt.Println(deleteTaskErr)
		return
	}
	fmt.Println("Task deleted successfully")
}
func (cmd *TaskCommand) ListUserTasks() {
	fmt.Println("List User Tasks Command ...")
	var userId int
	fmt.Println("Please enter the user id : ")
	userId, readUserIdErr := cmd.ioUtil.ReadNumber()
	if readUserIdErr != nil {
		fmt.Println("Please enter a valid user id")
		return
	}
	tasks, err := cmd.repo.ListUserTasks(userId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("List of all tasks : ")
	fmt.Println(tasks)
}
func (cmd *TaskCommand) ChangeTaskStatus() {
	fmt.Println("Change Task Status Command ...")
	var taskId int
	var isDone bool
	fmt.Println("Please enter the id of the task : ")
	taskId, readTaskIdErr := cmd.ioUtil.ReadNumber()
	if readTaskIdErr != nil {
		fmt.Println(readTaskIdErr)
		return
	}
	fmt.Println("Please enter the new status of the task : ")
	isDone, readIsDoneErr := cmd.ioUtil.ReadBool()
	if readIsDoneErr != nil {
		fmt.Println(readIsDoneErr)
		return
	}

	task, findTaskErr := cmd.repo.GetTask(taskId)
	if findTaskErr != nil {
		fmt.Println(findTaskErr)
		return
	}
	task.IsDone = isDone
	editTaskErr := cmd.repo.EditTask(task)
	if editTaskErr != nil {
		fmt.Println(editTaskErr)
		return
	}
	fmt.Println("Task status changed successfully")
}
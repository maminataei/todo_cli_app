package model

type Task struct {
	Id int
	Title string
	DueDate string
	Category string
	IsDone bool
	UserId int
}
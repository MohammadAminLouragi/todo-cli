package dto

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID       int
	Title    string
	DueDate  string
	Category int
	IsDone   bool
	UserId   int
}

type Category struct {
	ID     int
	Title  string
	Color  string
	UserId int
}
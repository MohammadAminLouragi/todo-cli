package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

var userStorage []User
var taskStorage []Task
var categoryStorage []Category
var authenticatedUser *User

func main() {
	fmt.Println("Hello to TODO app.")

	// load users from user.txt file
	loadUsersFromFile("user.txt")

	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	for {

		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command:")
		scanner.Scan()
		*command = scanner.Text()

	}

}

func loadUsersFromFile(s string) {
	file, err := os.Open(s)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var id int
		var name, email, password string
		line := scanner.Text()
	
		parts := strings.Split(line, ",")
		if len(parts) != 4 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		fmt.Sscanf(parts[0], "ID:%d", &id)
		fmt.Sscanf(parts[1], "Name:%s", &name)
		fmt.Sscanf(parts[2], "Email:%s", &email)
		fmt.Sscanf(parts[3], "Password:%s", &password)

		user := User{
			Id:       id,
			Name:     name,
			Email:    email,
			Password: password,
		}
		userStorage = append(userStorage, user)
	}

	for _, user := range userStorage {

		fmt.Printf("Loaded user: ID=%d, Name=%s, Email=%s\n", user.Id, user.Name, user.Email)
	}
}

func runCommand(command string) {
	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()
		if authenticatedUser == nil {
			return
		}
	}
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "list-task":
		listTasks()
	case "login-out":
		authenticatedUser = nil
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}

func listTasks() {
	for _, task := range taskStorage {
		if task.ID == authenticatedUser.Id {
			fmt.Println(task.ID, " | ", task.Title, " | ", task.Category, " | ", task.DueDate)
		}
	}
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, dueDate, category string

	fmt.Println("please enter the task title:")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the task dueDate:")
	scanner.Scan()
	dueDate = scanner.Text()

	fmt.Println("please enter the task category:")
	scanner.Scan()
	category = scanner.Text()

	categoryId, err := strconv.Atoi(category)
	if err != nil {
		fmt.Println("invalid category id")

		return
	}

	isFound := false
	for _, cat := range categoryStorage {
		if cat.ID == categoryId && cat.UserId == authenticatedUser.Id {
			isFound = true
			break
		}
	}
	if !isFound {
		fmt.Println("category not found")

		return
	}

	task := Task{
		ID:       len(taskStorage) + 1,
		Title:    name,
		Category: categoryId,
		DueDate:  dueDate,
		IsDone:   false,
		UserId:   authenticatedUser.Id,
	}

	taskStorage = append(taskStorage, task)
	fmt.Println("task:", name, category, dueDate)

}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the category color:")
	scanner.Scan()
	color = scanner.Text()

	category := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserId: authenticatedUser.Id,
	}

	categoryStorage = append(categoryStorage, category)

	fmt.Println("category:", category.ID, title, color, category.UserId)
}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, email, password string

	fmt.Println("please enter your name:")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password:")
	scanner.Scan()
	password = scanner.Text()

	user := User{
		Id:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)

	// Add new user to user.txt file
	file, err := os.OpenFile("user.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("ID:%d,Name:%s,Email:%s,Password:%s\n", user.Id, user.Name, user.Email, user.Password))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("User registration completed: ", user.Id, user.Name, user.Email)

	//fmt.Printf("userStorage: %v \n", userStorage)
}

func login() {

	scanner := bufio.NewScanner(os.Stdin)
	var name, email, password string
	var id int

	fmt.Println("please enter your name:")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password:")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			id = user.Id
			fmt.Println("You are logged in.")
			authenticatedUser = &user

			break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("User Not Found.")
		return
	}

	fmt.Println("user:", id, name, email, password)

}

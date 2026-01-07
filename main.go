package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	contract "github.com/MohammadAminLouragi/todo-cli/contract"
	entity "github.com/MohammadAminLouragi/todo-cli/entity"
	"github.com/MohammadAminLouragi/todo-cli/storage"
	"golang.org/x/crypto/bcrypt"
)

var (
	UserStorage       []entity.User
	taskStorage       []entity.Task
	categoryStorage   []entity.Category
	authenticatedUser *entity.User
	writeToMyStorage  contract.WriteUserToStorage
	ReadFromMyStorage contract.ReadUserFromStorage
)

const (
	UserFile          = "user.txt"
	SerializationMode = "json"
)

func main() {
	fmt.Println("Hello to TODO app...fake")
	fmt.Println("Please enter a command to continue...")


	// init storage
	dataStore := storage.NewFileStorage(&UserStorage, SerializationMode, UserFile)

	writeToMyStorage = dataStore
	ReadFromMyStorage = dataStore
	ReadFromMyStorage.Load()

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

func runCommand(command string) {
	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()
		if authenticatedUser == nil {
			return
		}
	}

	if command == "login" && authenticatedUser != nil {
		fmt.Println("You are already logged in.")
		return
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
	case "login":
		login()
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

	task := entity.Task{
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

	category := entity.Category{
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

	// Hash the password before storing
	hashedPassword, err := hashPassword(password)
	if err != nil {
		fmt.Println("error while hashing password")

		return
	}

	user := entity.User{
		Id:       len(UserStorage) + 1,
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	UserStorage = append(UserStorage, user)

	// Add new user to user.txt file
	writeToMyStorage.Save(user)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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

	for _, user := range UserStorage {

		if user.Email == email && CheckPasswordHash(password, user.Password) {
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

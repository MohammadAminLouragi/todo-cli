package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	Id       int
	Email    string
	Password string
}

var userStorage []User

func main() {
	fmt.Println("Hello to TODO app.")

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
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	default:
		fmt.Println("command is not valid", command)
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

	fmt.Println("task:", name, category, dueDate)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color, userid string

	fmt.Println("please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the category color:")
	scanner.Scan()
	color = scanner.Text()

	fmt.Println("please enter the user id:")
	scanner.Scan()
	userid = scanner.Text()

	fmt.Println("category:", title, color, userid)
}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, email, password string

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password:")
	scanner.Scan()
	password = scanner.Text()

	id = email
	fmt.Println("user:", id, email, password)

	user := User{
		Id:       len(userStorage) + 1,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)

	fmt.Printf("userStorage: %v \n", userStorage)
}

func login() {

	scanner := bufio.NewScanner(os.Stdin)
	var id, email, password string

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password:")
	scanner.Scan()
	password = scanner.Text()

	id = email

	fmt.Println("user:", id, email, password)

}

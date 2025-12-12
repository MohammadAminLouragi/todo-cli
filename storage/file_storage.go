package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/MohammadAminLouragi/todo-cli/dto"
)



type FileStorage struct {
	SerializationMode string
	UserPath          string
	UserStorage *[]dto.User
}



func NewFileStorage(userTorage *[]dto.User, mode string, userPath string) *FileStorage {
	return &FileStorage{
		SerializationMode: mode,
		UserPath:          userPath,
		UserStorage: userTorage,
	}
}


func (f *FileStorage) Save(user dto.User) {

	file, err := os.OpenFile("user.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	var output string

	switch f.SerializationMode {
	case "json":
		data, err := json.Marshal(user)
		if err != nil {
			fmt.Println("Error serializing to JSON:", err)
			return
		}
		output = string(data) + "\n"

	default: // fallback to text format
		output = fmt.Sprintf(
			"ID:%d,Name:%s,Email:%s,Password:%s\n",
			user.Id, user.Name, user.Email, user.Password,
		)
	}

	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("User registration completed: ", user.Id, user.Name, user.Email)
}

func (f *FileStorage) Load() {
	file, err := os.Open(f.UserPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		user, err := parseUserLine(line, f.SerializationMode)
		if err != nil {
			fmt.Println("Skipping invalid line:", err)
			continue
		}

		*f.UserStorage = append(*f.UserStorage, *user)
		fmt.Printf("Loaded user\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func parseUserLine(line string, mode string) (*dto.User, error) {

	// JSON Mode
	if mode == "json" {
		var u dto.User
		if err := json.Unmarshal([]byte(line), &u); err != nil {
			return nil, fmt.Errorf("json parse failed: %w", err)
		}
		return &u, nil
	}
	return nil, fmt.Errorf("unsupported serialization mode")
}

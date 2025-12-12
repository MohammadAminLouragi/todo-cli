package storage

import "github.com/MohammadAminLouragi/todo-cli/dto"

type MyStorage interface {
	Save(user dto.User)
	Load() 
}
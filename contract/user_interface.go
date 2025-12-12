package storage

import entity "github.com/MohammadAminLouragi/todo-cli/entity"

type WriteUserToStorage interface {
	Save(user entity.User)
}

type ReadUserFromStorage interface {
	Load()
}

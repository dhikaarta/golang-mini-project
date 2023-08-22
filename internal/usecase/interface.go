package usecase

import (
	"task-management-app/internal/entity"
)

type TaskUseCase interface {
	CreateTask(title, description, content string) (entity.Task, error)
	ListTask() ([]entity.Task, error)
	UpdateTask(id, title, description, content string) error
}

type UserUseCase interface {
	CreateUser(username, email, password string) (string, error)
	LoginUser(username, password string) (entity.User, error)
	ListUser() ([]entity.User, error)
}

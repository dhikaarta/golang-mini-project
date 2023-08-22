package persistence

import "task-management-app/internal/entity"

type TaskRepository interface {
	CreateTask(task entity.Task) (string, error)
	ListTasks() ([]entity.Task, error)
	UpdateTask(task entity.Task) error
}

type UserRepository interface {
	CreateUser(user entity.User) (string, error)
	GetUserByUsername(username string) (entity.User, error)
	ListUsers() ([]entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(userID string) error
	CheckUserExist(username string) (bool, error)
	CheckEmailExist(email string) (bool, error)
}
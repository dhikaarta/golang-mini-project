package usecase

import (
	"task-management-app/interfaces/persistence"
	"task-management-app/internal/entity"
)

type UserInteractor struct {
	userRepository persistence.UserRepository
}

func NewUserUseCase(userRepository persistence.UserRepository) *UserInteractor {
	return &UserInteractor{
		userRepository: userRepository,
	}
}

func (u *UserInteractor) CreateUser(username, email, password string) (string, error) {

	user := entity.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	createdUserID, err := u.userRepository.CreateUser(user)
	if err != nil {
		return "", err
	}

	return createdUserID, nil

}

func (u *UserInteractor) LoginUser(username, password string) (entity.User, error) {

	user, err := u.userRepository.GetUserByUsername(username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserInteractor) ListUser() ([]entity.User, error) {

	users, err := u.userRepository.ListUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

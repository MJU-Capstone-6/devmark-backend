package user

import (
	"context"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type UserService struct {
	Repository interfaces.IRepository
}

func (u *UserService) FindUserByUserName(username *string) (*repository.User, error) {
	user, err := u.Repository.FindUserByUsername(context.Background(), username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) CreateUser(arg repository.CreateUserParams) (*repository.User, error) {
	user, err := u.Repository.CreateUser(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) UpdateUser(arg repository.UpdateUserParams) (*repository.User, error) {
	user, err := u.Repository.UpdateUser(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func InitUserService(repo interfaces.IRepository) *UserService {
	return &UserService{Repository: repo}
}

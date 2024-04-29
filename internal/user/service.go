package user

import (
	"context"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type UserService struct {
	UserRepository interfaces.IUserRepository
}

func (u *UserService) FindUserByUserName(username *string) (*repository.User, error) {
	user, err := u.UserRepository.FindUserByUsername(context.Background(), username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func InitUserService(repo interfaces.IUserRepository) *UserService {
	return &UserService{UserRepository: repo}
}

package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

//go:generate mockery --name IUserService
type IUserService interface {
	FindUserByUserName(*string) (*repository.User, error)
	FindUserById(int) (*repository.User, error)
	FindJoinedWorkspace(int) (*repository.UserWorkspaceView, error)
	CreateUser(repository.CreateUserParams) (*repository.User, error)
	UpdateUser(repository.UpdateUserParams) (*repository.User, error)
}

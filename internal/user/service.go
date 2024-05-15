package user

import (
	"context"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type UserService struct {
	Repository interfaces.IRepository
}

func (u *UserService) FindUserByUserName(username *string) (*repository.User, error) {
	user, err := u.Repository.FindUserByUsername(context.Background(), username)
	if err != nil {
		return nil, customerror.UserNotFound(err)
	}
	return &user, nil
}

func (u *UserService) FindUserById(id int) (*repository.FindUserByIdRow, error) {
	user, err := u.Repository.FindUserById(context.Background(), int64(id))
	if err != nil {
		return nil, customerror.UserNotFound(err)
	}
	return &user, nil
}

func (u *UserService) FindJoinedWorkspace(id int) (*repository.UserWorkspaceView, error) {
	userId := int64(id)
	workspaceRow, err := u.Repository.FindUserWorkspace(context.Background(), &userId)
	if err != nil {
		return &repository.UserWorkspaceView{ID: &userId, Workspaces: []repository.Workspace{}}, nil
	}
	return &workspaceRow, nil
}

func (u *UserService) CreateUser(arg repository.CreateUserParams) (*repository.User, error) {
	user, err := u.Repository.CreateUser(context.Background(), arg)
	if err != nil {
		return nil, customerror.UserCreationFail(err)
	}
	return &user, nil
}

func (u *UserService) UpdateUser(arg repository.UpdateUserParams) (*repository.User, error) {
	user, err := u.Repository.UpdateUser(context.Background(), arg)
	if err != nil {
		return nil, customerror.UserUpdateFail(err)
	}
	return &user, nil
}

func InitUserService(repo interfaces.IRepository) *UserService {
	return &UserService{Repository: repo}
}

package interfaces

import (
	"context"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
)

//go:generate mockery --name IRepository
type IRepository interface {
	FindUserByUsername(context.Context, *string) (repository.User, error)
	FindUserById(context.Context, int64) (repository.User, error)
	CreateRefreshToken(context.Context, repository.CreateRefreshTokenParams) (repository.RefreshToken, error)
	FindRefreshTokenByUserID(context.Context, *int32) (repository.RefreshToken, error)
	UpdateRefreshToken(context.Context, repository.UpdateRefreshTokenParams) (repository.RefreshToken, error)
	CreateUser(context.Context, repository.CreateUserParams) (repository.User, error)
	UpdateUser(context.Context, repository.UpdateUserParams) (repository.User, error)
	CreateWorkspace(context.Context, *string) (repository.Workspace, error)
	FindWorkspace(context.Context, int64) (repository.WorkspaceUserCategory, error)
	UpdateWorkspace(context.Context, repository.UpdateWorkspaceParams) (repository.Workspace, error)
	DeleteWorkspace(context.Context, int64) error
	FindInviteCodeByWorkspaceID(context.Context, *int32) (repository.InviteCode, error)
	CreateInviteCode(context.Context, repository.CreateInviteCodeParams) (repository.InviteCode, error)
	JoinWorkspace(context.Context, repository.JoinWorkspaceParams) error
	CreateCategory(context.Context, *string) (repository.Category, error)
	RegisterCategoryToWorkspace(context.Context, repository.RegisterCategoryToWorkspaceParams) error
	UpdateCategory(context.Context, repository.UpdateCategoryParams) (repository.Category, error)
	DeleteCategory(context.Context, int64) error
}

package interfaces

import (
	"context"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
)

//go:generate mockery --name IRepository
type IRepository interface {
	FindUserByUsername(context.Context, *string) (repository.User, error)
	FindUserById(context.Context, int64) (repository.FindUserByIdRow, error)
	CreateRefreshToken(context.Context, repository.CreateRefreshTokenParams) (repository.RefreshToken, error)
	FindRefreshTokenByUserID(context.Context, *int32) (repository.RefreshToken, error)
	UpdateRefreshToken(context.Context, repository.UpdateRefreshTokenParams) (repository.RefreshToken, error)
	FindUserWorkspace(context.Context, *int64) (repository.UserWorkspaceView, error)
	CreateUser(context.Context, repository.CreateUserParams) (repository.User, error)
	UpdateUser(context.Context, repository.UpdateUserParams) (repository.User, error)
	CreateWorkspace(context.Context, repository.CreateWorkspaceParams) (repository.Workspace, error)
	FindWorkspace(context.Context, int64) (repository.FindWorkspaceRow, error)
	UpdateWorkspace(context.Context, repository.UpdateWorkspaceParams) (repository.Workspace, error)
	DeleteWorkspace(context.Context, int64) error
	FindInviteCodeByCode(context.Context, *string) (repository.InviteCode, error)
	FindInviteCodeByWorkspaceID(context.Context, *int32) (repository.InviteCode, error)
	CreateInviteCode(context.Context, repository.CreateInviteCodeParams) (repository.InviteCode, error)
	JoinWorkspace(context.Context, repository.JoinWorkspaceParams) error
	JoinWorkspaceWithoutCode(context.Context, repository.JoinWorkspaceWithoutCodeParams) error
	FindCategoryById(context.Context, int64) (repository.Category, error)
	CreateCategory(context.Context, *string) (repository.Category, error)
	RegisterCategoryToWorkspace(context.Context, repository.RegisterCategoryToWorkspaceParams) error
	UpdateCategory(context.Context, repository.UpdateCategoryParams) (repository.Category, error)
	DeleteCategory(context.Context, int64) error
	FindBookmark(context.Context, int64) (repository.FindBookmarkRow, error)
	CreateBookmark(context.Context, repository.CreateBookmarkParams) (repository.Bookmark, error)
	DeleteBookmark(context.Context, int64) error
	UpdateBookmark(context.Context, repository.UpdateBookmarkParams) (repository.Bookmark, error)
	FindBookmarkComment(context.Context, int64) ([]*repository.Comment, error)
	FindComment(context.Context, int64) (repository.Comment, error)
	CreateComment(context.Context, repository.CreateCommentParams) (repository.Comment, error)
	DeleteComment(context.Context, int64) error
	UpdateComment(context.Context, repository.UpdateCommentParams) (repository.Comment, error)
	FindWorkspaceCategory(context.Context, int64) ([]*repository.Category, error)
	FindWorkspaceCategoryBookmark(context.Context, repository.FindWorkspaceCategoryBookmarkParams) ([]repository.Bookmark, error)
	CheckWorkspaceExists(context.Context, int64) (repository.Workspace, error)
}

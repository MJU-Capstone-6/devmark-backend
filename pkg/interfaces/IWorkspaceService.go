package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

//go:generate mockery --name IWorkspaceService

type IWorkspaceService interface {
	Create(int, repository.CreateWorkspaceParams) (*repository.Workspace, error)
	FindById(int) (*repository.FindWorkspaceRow, error)
	FindCategoriesById(int) (*[]*repository.Category, error)
	FindCategoryBookmark(repository.FindWorkspaceCategoryBookmarkParams) (*[]repository.Bookmark, error)
	SearchBookmark(repository.SearchWorkspaceBookmarkParams) (*[]repository.Bookmark, error)
	Update(repository.UpdateWorkspaceParams) (*repository.Workspace, error)
	Delete(int) error
	Join(string, repository.JoinWorkspaceParams) error
	RegisterCategory(repository.RegisterCategoryToWorkspaceParams) error
}

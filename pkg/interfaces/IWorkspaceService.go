package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

//go:generate mockery --name IWorkspaceService

type IWorkspaceService interface {
	Create(int, repository.CreateWorkspaceParams) (*repository.Workspace, error)
	FindById(int) (*repository.WorkspaceUserCategory, error)
	Update(repository.UpdateWorkspaceParams) (*repository.Workspace, error)
	Delete(int) error
	Join(string, repository.JoinWorkspaceParams) error
}

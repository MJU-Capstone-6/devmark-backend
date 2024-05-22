package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

//go:generate mockery --name IWorkspaceCodeService
type IWorkspaceCodeService interface {
	FindByCode(string) (*repository.WorkspaceCode, error)
	Update(repository.UpdateWorkspaceCodeParams) (*repository.WorkspaceCode, error)
	Create(repository.CreateWorkspaceCodeParams) (*repository.WorkspaceCode, error)
}

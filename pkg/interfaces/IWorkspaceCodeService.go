package interfaces

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/request"
)

//go:generate mockery --name IWorkspaceCodeService
type IWorkspaceCodeService interface {
	CreateCode(int) *string
	FindByCode(string) (*repository.FindWorkspaceCodeRow, error)
	PredictCategory(request.PredictCategoryParam) (*repository.Bookmark, error)
	Update(repository.UpdateWorkspaceCodeParams) (*repository.WorkspaceCode, error)
	Create(repository.CreateWorkspaceCodeParams) (*repository.WorkspaceCode, error)
}

package workspacecode

import "github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"

type WorkspaceCodeService struct {
	Repository interfaces.IRepository
}

func InitWorkspaceCodeService() *WorkspaceCodeService {
	return &WorkspaceCodeService{}
}

func (w WorkspaceCodeService) WithRepository(repo interfaces.IRepository) WorkspaceCodeService {
	return w
}

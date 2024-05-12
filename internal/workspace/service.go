package workspace

import (
	"context"
	"log"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type WorkspaceService struct {
	Repository interfaces.IRepository
}

func (w *WorkspaceService) Create(name string) (*repository.Workspace, error) {
	workspace, err := w.Repository.CreateWorkspace(context.Background(), &name)
	if err != nil {
		return nil, err
	}
	return &workspace, nil
}

func (w *WorkspaceService) FindById(id int) (*repository.WorkspaceUserCategory, error) {
	workspaceId := int64(id)
	workspace, err := w.Repository.FindWorkspace(context.Background(), workspaceId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &workspace, nil
}

func (w *WorkspaceService) Update(param repository.UpdateWorkspaceParams) (*repository.Workspace, error) {
	workspace, err := w.Repository.UpdateWorkspace(context.Background(), param)
	if err != nil {
		return nil, err
	}
	return &workspace, nil
}

func (w *WorkspaceService) Delete(id int) error {
	err := w.Repository.DeleteWorkspace(context.Background(), int64(id))
	if err != nil {
		return err
	}
	return nil
}

func InitWorkspaceService(repo interfaces.IRepository) *WorkspaceService {
	return &WorkspaceService{Repository: repo}
}

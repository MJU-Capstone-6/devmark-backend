package workspace

import (
	"context"
	"log"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type WorkspaceService struct {
	Repository        interfaces.IRepository
	InviteCodeService interfaces.IInviteCodeService
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

func (w *WorkspaceService) Join(code string, param repository.JoinWorkspaceParams) error {
	verifyParam := utils.VerifyCodeParam{
		Code:        code,
		WorkspaceId: int(param.WorkspaceID),
	}
	if ok, err := w.InviteCodeService.VerifyCode(verifyParam); !ok {
		if err != nil {
			return err
		}
		return customerror.CodeVerifyFailedErr(nil)
	}

	err := w.Repository.JoinWorkspace(context.Background(), param)
	if err != nil {
		return err
	}
	return nil
}

func InitWorkspaceService(repo interfaces.IRepository) *WorkspaceService {
	return &WorkspaceService{Repository: repo}
}

func (w WorkspaceService) WithInviteCodeService(service interfaces.IInviteCodeService) WorkspaceService {
	w.InviteCodeService = service
	return w
}

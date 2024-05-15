package workspace

import (
	"context"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type WorkspaceService struct {
	Repository        interfaces.IRepository
	InviteCodeService interfaces.IInviteCodeService
}

func (w *WorkspaceService) Create(userId int, createWorkspaceParam repository.CreateWorkspaceParams) (*repository.Workspace, error) {
	workspace, err := w.Repository.CreateWorkspace(context.Background(), createWorkspaceParam)
	if err != nil {
		return nil, customerror.WorkspaceCreateFail(err)
	}
	param := repository.JoinWorkspaceWithoutCodeParams{
		WorkspaceID: workspace.ID,
		UserID:      int64(userId),
	}
	err = w.Repository.JoinWorkspaceWithoutCode(context.Background(), param)
	if err != nil {
		return nil, customerror.WorkspaceJoinFailErr(err)
	}
	return &workspace, nil
}

func (w *WorkspaceService) FindById(id int) (*repository.WorkspaceUserCategory, error) {
	workspaceId := int64(id)
	workspace, err := w.Repository.FindWorkspace(context.Background(), workspaceId)
	if err != nil {
		return nil, customerror.WorkspaceNotFoundErr(err)
	}
	return &workspace, nil
}

func (w *WorkspaceService) Update(param repository.UpdateWorkspaceParams) (*repository.Workspace, error) {
	workspace, err := w.Repository.UpdateWorkspace(context.Background(), param)
	if err != nil {
		return nil, customerror.WorkspaceUpdateFail(err)
	}
	return &workspace, nil
}

func (w *WorkspaceService) Delete(id int) error {
	err := w.Repository.DeleteWorkspace(context.Background(), int64(id))
	if err != nil {
		return customerror.WorkspaceDeleteFail(err)
	}
	return nil
}

func (w *WorkspaceService) Join(code string, param repository.JoinWorkspaceParams) error {
	inviteCode, err := w.InviteCodeService.FindByCode(code)
	if err != nil {
		return err
	}

	verifyParam := utils.VerifyCodeParam{
		Code:        code,
		WorkspaceId: int(*inviteCode.WorkspaceID),
	}
	if ok, err := w.InviteCodeService.VerifyCode(verifyParam); !ok {
		if err != nil {
			return err
		}
		return customerror.CodeVerifyFailedErr(nil)
	}

	param.WorkspaceID = int64(*inviteCode.WorkspaceID)

	err = w.Repository.JoinWorkspace(context.Background(), param)
	if err != nil {
		return customerror.WorkspaceJoinFailErr(err)
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

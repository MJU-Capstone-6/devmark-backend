package invitecode

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/request"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type InviteCodeService struct {
	Repository       interfaces.IRepository
	WorkspaceService interfaces.IWorkspaceService
}

const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (i *InviteCodeService) CreateCode(length int) *string {
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = CHARSET[rand.Intn(len(CHARSET))]
	}
	generatedString := string(randomString)
	return &generatedString
}

func (i *InviteCodeService) CreateInviteCode(param request.CreateInviteCodeParam) (*repository.InviteCode, error) {
	code := i.CreateCode(constants.CODE_LENGTH)
	if inviteCode, err := i.FindByWorkspaceID(param.WorkspaceID); err == nil {
		updateInviteCodeParam := repository.UpdateInviteCodeParams{
			Code: code,
			ID:   inviteCode.ID,
		}
		inviteCode, err := i.Repository.UpdateInviteCode(context.Background(), updateInviteCodeParam)
		if err != nil {
			return nil, customerror.CodeUpdateFail(err)
		}
		return &inviteCode, nil
	}
	_, err := i.WorkspaceService.FindById(param.WorkspaceID)
	if err != nil {
		return nil, customerror.WorkspaceNotFoundErr(err)
	}
	workspaceId := int32(param.WorkspaceID)
	inviteCodeParam := repository.CreateInviteCodeParams{
		WorkspaceID: &workspaceId,
		Code:        code,
	}
	inviteCode, err := i.Repository.CreateInviteCode(context.Background(), inviteCodeParam)
	if err != nil {
		return nil, customerror.CodeCreationFail(err)
	}
	return &inviteCode, nil
}

func (i *InviteCodeService) FindByCode(code string) (*repository.InviteCode, error) {
	inviteCode, err := i.Repository.FindInviteCodeByCode(context.Background(), &code)
	if err != nil {
		return nil, customerror.CodeNotFound(err)
	}
	return &inviteCode, nil
}

func (i *InviteCodeService) FindByWorkspaceID(id int) (*repository.InviteCode, error) {
	parsedId := int32(id)
	_, err := i.WorkspaceService.FindById(id)
	if err != nil {
		return nil, err
	}
	inviteCode, err := i.Repository.FindInviteCodeByWorkspaceID(context.Background(), &parsedId)
	if err != nil {
		return nil, customerror.CodeNotFound(err)
	}
	return &inviteCode, nil
}

func (i *InviteCodeService) VerifyCode(param utils.VerifyCodeParam) (bool, error) {
	inviteCode, err := i.FindByWorkspaceID(param.WorkspaceId)
	if err != nil {
		return false, customerror.CodeNotFound(err)
	}
	if param.Code != *inviteCode.Code {
		return false, customerror.CodeVerifyFailedErr(errors.New(""))
	}
	if inviteCode.ExpiredAt.Time.Before(time.Now()) {
		return false, customerror.CodeVerifyFailedErr(errors.New(""))
	}
	return true, nil
}

func InitInviteCodeService() *InviteCodeService {
	return &InviteCodeService{}
}

func (i InviteCodeService) WithRepository(repo interfaces.IRepository) InviteCodeService {
	i.Repository = repo
	return i
}

func (i InviteCodeService) WithWorkspaceService(service interfaces.IWorkspaceService) InviteCodeService {
	i.WorkspaceService = service
	return i
}

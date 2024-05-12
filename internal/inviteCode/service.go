package invitecode

import (
	"context"
	"math/rand"
	"time"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
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

func (i *InviteCodeService) CreateInviteCode(param repository.CreateInviteCodeParams) (*repository.InviteCode, error) {
	code := i.CreateCode(constants.CODE_LENGTH)
	param.Code = code
	_, err := i.WorkspaceService.FindById(int(*param.WorkspaceID))
	if err != nil {
		return nil, customerror.WorkspaceNotFoundErr(err)
	}
	inviteCode, err := i.Repository.CreateInviteCode(context.Background(), param)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return &inviteCode, nil
}

func (i *InviteCodeService) VerifyCode(param utils.VerifyCodeParam) (bool, error) {
	inviteCode, err := i.FindByWorkspaceID(param.WorkspaceId)
	if err != nil {
		return false, err
	}
	if param.Code != *inviteCode.Code {
		return false, nil
	}
	if inviteCode.ExpiredAt.Time.Before(time.Now()) {
		return false, nil
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

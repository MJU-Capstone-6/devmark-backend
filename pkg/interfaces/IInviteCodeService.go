package interfaces

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
)

//go:generate mockery --name IInviteCodeService
type IInviteCodeService interface {
	CreateCode(int) *string
	CreateInviteCode(repository.CreateInviteCodeParams) (*repository.InviteCode, error)
	FindByWorkspaceID(int) (*repository.InviteCode, error)
	FindByCode(string) (*repository.InviteCode, error)
	VerifyCode(utils.VerifyCodeParam) (bool, error)
}

package invitecode

import (
	"net/http"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type InviteCodeController struct {
	InviteCodeService interfaces.IInviteCodeService
}

// GenerateInviteCodeController godoc
//
//	@summary	Generate Invite Code from workspace
//	@schemes
//	@description	워크스페이스의 초대코드를 생성합니다.
//	@tags			invitecode
//	@accept			json
//	@produce		json
//	@param			body	body		repository.CreateInviteCodeParams	true	"body to Generate Invite code"
//	@success		200		{object}	repository.InviteCode
//	@failure		401		{object}	customerror.CustomError
//	@failure		404		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/invitecode [POST]
func (i *InviteCodeController) GenerateInviteCodeController(ctx echo.Context) error {
	var param repository.CreateInviteCodeParams
	err := ctx.Bind(&param)
	if err != nil {
		return customerror.InternalServerError(err)
	}
	inviteCode, err := i.InviteCodeService.CreateInviteCode(param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, inviteCode)
}

func InitInviteCodeController() *InviteCodeController {
	return &InviteCodeController{}
}

func (i InviteCodeController) WithInviteCodeService(service interfaces.IInviteCodeService) InviteCodeController {
	i.InviteCodeService = service
	return i
}

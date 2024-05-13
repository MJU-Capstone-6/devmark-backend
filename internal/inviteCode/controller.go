package invitecode

import (
	"net/http"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type InviteCodeController struct {
	InviteCodeService interfaces.IInviteCodeService
}

// GenerateInviteCodeController godoc
//
//	@summary	retrive user info from kakao oauth
//	@schemes
//	@description	retrive user info from kakao oauth. if user exists in our service, then return access token.
//	@tags			users
//	@accept			json
//	@produce		json
//	@success		200	{object}	getkakaoinforesponse
//	@failure		401	{object}	customerror.customerror
//	@failure		500 {object}	customerror.customerror
//	@router			/invitecode [POST]
func (i *InviteCodeController) GenerateInviteCodeController(ctx echo.Context) error {
	var param repository.CreateInviteCodeParams
	err := ctx.Bind(&param)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}
	inviteCode, err := i.InviteCodeService.CreateInviteCode(param)
	if err != nil {
		if _, ok := err.(*customerror.CustomError); ok {
			return responses.NotFound(ctx, err)
		} else {
			return responses.InternalServer(ctx, customerror.InternalServerError(err))
		}
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

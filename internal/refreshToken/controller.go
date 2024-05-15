package refreshtoken

import (
	"net/http"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/request"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type RefreshTokenController struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

// RefreshAccessTokenController godoc
//
//	@summary	RefreshToken으로 AccessToken을 재발급 합니다.
//	@schemes
//	@description RefreshToken으로 AccessToken을 재발급 합니다.
//	@tags			refreshToken
//	@accept			json
//	@produce		json
//	@success		200	{object}	responses.RefreshAccessTokenResponse
//	@failure		401	{object}	customerror.CustomError
//	@failure		500 {object}	customerror.CustomError
//	@router			/api/v1/refresh [POST]
func (r *RefreshTokenController) RefreshAccessTokenController(ctx echo.Context) error {
	var param request.RefreshAccessTokenParam
	err := ctx.Bind(&param)
	if err != nil {
		return customerror.InvalidParamError(err)
	}
	accessToken, err := r.RefreshTokenService.RefreshAccesstoken(param.RefreshToken)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, accessToken)
}

func InitRefreshTokenController() *RefreshTokenController {
	return &RefreshTokenController{}
}

func (r RefreshTokenController) WithRefreshTokenService(service interfaces.IRefreshTokenService) RefreshTokenController {
	r.RefreshTokenService = service
	return r
}

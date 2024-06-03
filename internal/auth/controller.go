package auth

import (
	"errors"
	"net/http"

	"github.com/MJU-Capstone-6/devmark-backend/internal/config"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/request"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

const BASE_URL = "https://kapi.kakao.com/v2/user/me?secure_resource=true"

type AuthController struct {
	AuthService interfaces.IAuthService
	KakaoInfo   config.Kakao
}

// getkakaouserinfo godoc
//
//	@summary	retrive user info from kakao oauth
//	@schemes
//	@description	retrive user info from kakao oauth. if user exists in our service, then return access token.
//	@tags			auth
//	@accept			json
//	@produce		json
//	@param			body	body		request.AuthParam	true	"body to signup/signin"
//	@success		200		{object}	responses.GetKakaoInfoResponse
//	@failure		401		{object}	customerror.CustomError
//	@failure		422		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/auth/kakao [POST]
func (a *AuthController) GetKakaoUserInfo(ctx echo.Context) error {
	var param request.AuthParam
	err := ctx.Bind(&param)
	if err != nil {
		return customerror.InternalServerError(err)
	}
	key := ctx.Get("key")
	provider := ctx.Param("provider")
	if parsedKey, ok := key.(string); ok {
		userInfo, err := requestKakaoUserInfo(parsedKey)
		if err != nil {
			return err
		}
		if userInfo.ID == 0 {
			return customerror.TokenNotValidError(errors.New(""))
		}
		request, err := a.AuthService.KakaoSignUp(userInfo.Properties.Nickname, provider, param.RegistrationToken)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, request)
	} else {
		return customerror.TokenNotProvidedError(errors.New(""))
	}
}

func InitAuthController() *AuthController {
	return &AuthController{}
}

func (a AuthController) WithAuthService(service interfaces.IAuthService) AuthController {
	a.AuthService = service
	return a
}

func (a AuthController) WithKakaoInfo(kakaoInfo config.Kakao) AuthController {
	a.KakaoInfo = kakaoInfo
	return a
}

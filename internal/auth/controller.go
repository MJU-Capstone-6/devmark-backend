package auth

import (
	"errors"

	"github.com/MJU-Capstone-6/devmark-backend/internal/config"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
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
//	@tags			users
//	@accept			json
//	@produce		json
//	@success		200	{object}	getkakaoinforesponse
//	@failure		401	{object}	customerror.customerror
//	@failure		500 {object}	customerror.customerror
//	@router			/auth/kakao [POST]
func (a *AuthController) GetKakaoUserInfo(ctx echo.Context) error {
	key := ctx.Get("key")
	provider := ctx.Param("provider")
	if parsedKey, ok := key.(string); ok {
		userInfo, err := requestKakaoUserInfo(parsedKey)
		if err != nil {
			return err
		}
		if userInfo.ID == 0 {
			return responses.Unauthorized(ctx, customerror.TokenNotValidError(errors.New("")))
		}
		err = a.AuthService.KakaoSignUp(userInfo.Properties.Nickname, provider, ctx)
		if err != nil {
			return err
		}
	} else {
		return responses.Unauthorized(ctx, customerror.TokenNotProvidedError(errors.New("")))
	}
	return nil
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

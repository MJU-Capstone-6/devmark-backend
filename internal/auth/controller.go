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

// GetKakaoUserInfo godoc
//
//	@Summary	Retrive user info from kakao oauth
//	@Schemes
//	@Description	Retrive user info from kakao oauth. If user exists in our service, then return access token.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	GetKakaoInfoResponse
//	@Failure		401	{object}	customerror.CustomError
//	@Failure		500 {object}	customerror.CustomError
//	@Router			/auth/kakao [GET]
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

package auth

import (
	"crypto/ed25519"
	"errors"

	"github.com/MJU-Capstone-6/devmark-backend/internal/config"
	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/jwtToken"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
	"github.com/MJU-Capstone-6/devmark-backend/internal/user"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

const BASE_URL = "https://kapi.kakao.com/v2/user/me?secure_resource=true"

type AuthController struct {
	AuthService AuthService
	UserService *user.UserService
	JWTService  *jwtToken.JWTService
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
	if parsedKey, ok := key.(string); ok {
		userInfo, err := requestKakaoUserInfo(parsedKey)
		if err != nil {
			return err
		}
		if userInfo.ID == 0 {
			return responses.Unauthorized(ctx, customerror.TokenNotValidError(errors.New("")))
		}
		existedUser, err := a.UserService.FindUserByUserName(&userInfo.Properties.Nickname)
		if err != nil {
			return responses.OK(ctx, userInfo)
		}

		accessToken, err := a.JWTService.GenerateToken(int(existedUser.ID), constants.ACCESSTOKEN_EXPIRED_TIME)
		if err != nil {
			return responses.InternalServer(ctx, customerror.InternalServerError(err))
		}

		refreshToken, err := a.JWTService.GenerateToken(int(existedUser.ID), constants.REFRESH_TOKEN_EXPIRED_TIME)
		if err != nil {
			return responses.InternalServer(ctx, customerror.InternalServerError(err))
		}
		return responses.OK(ctx, GetKakaoInfoResponse{AccessToken: *accessToken, RefreshToken: *refreshToken})
	}
	return responses.Unauthorized(ctx, customerror.TokenNotProvidedError(errors.New("")))
}

func InitAuthController() *AuthController {
	return &AuthController{}
}

func (a AuthController) WithAuthService(conn *pgx.Conn) AuthController {
	authService := AuthService{Conn: conn}
	a.AuthService = authService
	return a
}

func (a AuthController) WithUserService(conn *pgx.Conn) AuthController {
	userService := user.InitUserService(repository.New(conn))
	a.UserService = userService
	return a
}

func (a AuthController) WithKakaoInfo(kakaoInfo config.Kakao) AuthController {
	a.KakaoInfo = kakaoInfo
	return a
}

func (a AuthController) WithJWTService(publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey, footer string) AuthController {
	jwtService := jwtToken.InitJWTService(publicKey, privateKey, footer)
	a.JWTService = jwtService
	return a
}

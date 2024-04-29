package auth

import (
	"errors"
	"net/http"

	"github.com/MJU-Capstone-6/devmark-backend/internal/config"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/user"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

const BASE_URL = "https://kapi.kakao.com/v2/user/me?secure_resource=true"

type AuthController struct {
	AuthService AuthService
	UserService *user.UserService
	KakaoInfo   config.Kakao
}

func (a *AuthController) GetKakaoUserInfo(ctx echo.Context) error {
	key := ctx.Get("key")
	if parsedKey, ok := key.(string); ok {
		userInfo, err := requestKakaoUserInfo(parsedKey)
		if err != nil {
			return err
		}
		if userInfo.ID == 0 {
			return ctx.JSON(http.StatusUnauthorized, customerror.TokenNotValid(errors.New("")))
		}
		_, err = a.UserService.FindUserByUserName(&userInfo.Properties.Nickname)
		if err != nil {
			return ctx.JSON(http.StatusOK, userInfo)
		}
		return ctx.JSON(http.StatusOK, GetKakaoInfoResponse{AccessKey: parsedKey})
	}
	return ctx.JSON(http.StatusUnauthorized, customerror.TokenNotProvidedError(errors.New("")))
}

func InitAuthController(conn *pgx.Conn, kakaoInfo config.Kakao) *AuthController {
	authService := AuthService{Conn: conn}
	userService := user.InitUserService(repository.New(conn))
	return &AuthController{AuthService: authService, UserService: userService, KakaoInfo: kakaoInfo}
}

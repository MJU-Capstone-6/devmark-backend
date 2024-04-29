package auth

import (
	"errors"
	"net/http"

	"github.com/MJU-Capstone-6/devmark-backend/internal/config"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

const BASE_URL = "https://kapi.kakao.com/v2/user/me?secure_resource=true"

type AuthController struct {
	AuthService AuthService
	KakaoInfo   config.Kakao
}

func (a *AuthController) GetKakaoUserInfo(ctx echo.Context) error {
	key := ctx.Get("key")
	if parsedKey, ok := key.(string); ok {
		user, err := requestKakaoUserInfo(parsedKey)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, user)
	}
	return ctx.JSON(http.StatusUnauthorized, customerror.UnauthorizedError(errors.New("")))
}

func InitAuthController(conn *pgx.Conn, kakaoInfo config.Kakao) *AuthController {
	authService := AuthService{Conn: conn}
	return &AuthController{AuthService: authService, KakaoInfo: kakaoInfo}
}

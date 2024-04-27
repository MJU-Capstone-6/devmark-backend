package auth

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService AuthService
}

func (a *AuthController) GetKakaoUserInfo(ctx echo.Context) error {
	return nil
}

func InitAuthController(conn *pgx.Conn) *AuthController {
	authService := AuthService{Conn: conn}
	return &AuthController{AuthService: authService}
}

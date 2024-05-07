package interfaces

import (
	"github.com/labstack/echo/v4"
)

//go:generate mockery --name IAuthService
type IAuthService interface {
	KakaoSignUp(string, string, echo.Context) error
}

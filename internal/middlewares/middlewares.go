package middlewares

import (
	"log"
	"strings"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type CustomMiddleware struct {
	userService     interfaces.IUserService
	jwtTokenService interfaces.IJWTService
}

func (cm *CustomMiddleware) ParseHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if header != "" {
			key := strings.Split(header, "Bearer")[1]
			trimmedKey := strings.TrimSpace(key)
			c.Set("key", trimmedKey)
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
		return next(c)
	}
}

func (cm *CustomMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if header != "" {
			key := strings.Split(header, "Bearer")[1]
			trimmedKey := strings.TrimSpace(key)
			token, err := cm.jwtTokenService.VerifyToken(trimmedKey)
			if err != nil {
				c.Error(err)
			}
			userId := token.Get(constants.TOKEN_DATA_KEY)
			log.Println(userId)
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
		return next(c)
	}
}

func InitMiddleware() *CustomMiddleware {
	return &CustomMiddleware{}
}

func (c CustomMiddleware) WithUserService(service interfaces.IUserService) CustomMiddleware {
	c.userService = service
	return c
}

func (c CustomMiddleware) WithJwtTokenService(service interfaces.IJWTService) CustomMiddleware {
	c.jwtTokenService = service
	return c
}

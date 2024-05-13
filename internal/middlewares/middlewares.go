package middlewares

import (
	"errors"
	"strconv"
	"strings"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
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
				return responses.Unauthorized(c, customerror.TokenNotValidError(err))
			}
			userId, err := strconv.Atoi(token.Get(constants.TOKEN_DATA_KEY))
			if err != nil {
				return responses.InternalServer(c, customerror.InternalServerError(err))
			}
			user, err := cm.userService.FindUserById(userId)
			if err != nil {
				return responses.Unauthorized(c, customerror.UserNotFound(err))
			}
			c.Set("user", user)
			if err := next(c); err != nil {
				return responses.InternalServer(c, customerror.InternalServerError(err))
			}
			return nil
		} else {
			return responses.Unauthorized(c, customerror.TokenNotProvidedError(errors.New("")))
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

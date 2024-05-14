package utils

import (
	"errors"
	"strconv"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/labstack/echo/v4"
)

func ParseURLParam(ctx echo.Context, paramName string) (*int, error) {
	param, err := strconv.Atoi(ctx.Param(paramName))
	if err != nil {
		return nil, customerror.InvalidParamError(err)
	}
	return &param, nil
}

func GetAuthUser(ctx echo.Context) (*repository.User, error) {
	if user, ok := ctx.Get(constants.USER_CONTEXT_KEY).(*repository.User); ok {
		return user, nil
	}
	return nil, customerror.UserNotFound(errors.New(""))
}

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

func GetAuthUser(ctx echo.Context) (*repository.FindUserByIdRow, error) {
	if user, ok := ctx.Get(constants.USER_CONTEXT_KEY).(*repository.FindUserByIdRow); ok {
		return user, nil
	}
	return nil, customerror.UserNotFound(errors.New(""))
}

func SliceValueIntoNum(arr []string) (*[]int64, error) {
	intSlice := make([]int64, len(arr))
	for index, value := range arr {
		num, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		intSlice[index] = int64(num)
	}
	return &intSlice, nil
}

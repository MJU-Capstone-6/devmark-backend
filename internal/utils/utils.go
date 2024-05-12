package utils

import (
	"strconv"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
	"github.com/labstack/echo/v4"
)

func ParseURLParam(ctx echo.Context, paramName string) (*int, error) {
	param, err := strconv.Atoi(ctx.Param(paramName))
	if err != nil {
		return nil, responses.BadRequest(ctx, customerror.InvalidParamError(err))
	}
	return &param, nil
}

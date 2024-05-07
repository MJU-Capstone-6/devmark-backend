package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Unauthorized(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusUnauthorized, data)
}

func OK(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, data)
}

func InternalServer(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusInternalServerError, data)
}

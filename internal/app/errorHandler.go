package app

import (
	"net/http"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if customError, ok := err.(*customerror.CustomError); ok {
		code = customError.StatusCode
		c.JSON(code, customError)
	}
}

package app

import (
	"log"
	"net/http"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	log.Println(err)
	code := http.StatusInternalServerError
	if customError, ok := err.(*customerror.CustomError); ok {
		code = customError.StatusCode
		c.JSON(code, customError)
	} else if echoError, ok := err.(*echo.HTTPError); ok {
		switch echoError.Code {
		case http.StatusNotFound:
			c.JSON(echoError.Code, customerror.PageNotFound(echoError))
		case http.StatusInternalServerError:
			c.JSON(echoError.Code, customerror.InternalServerError(echoError))
		}
	}
}

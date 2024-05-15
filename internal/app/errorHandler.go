package app

import (
	"errors"
	"log"
	"net/http"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	log.Println(err)
	code := http.StatusInternalServerError
	if customError, ok := err.(*customerror.CustomError); ok {
		log.Println(customError)
		code = customError.StatusCode
		c.JSON(code, customError)
	} else {
		if c.Response().Status == 404 {
			c.JSON(c.Response().Status, customerror.PageNotFound(errors.New("")))
		}
	}
}

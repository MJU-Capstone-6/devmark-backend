package customerror

import "net/http"

var UNAUTHORIZED_ERROR = NewError(http.StatusUnauthorized, "Access key must be provide.", "101")

func UnauthorizedError(err error) error {
	unauthorizedError := UNAUTHORIZED_ERROR
	unauthorizedError.Data = err.Error()
	return unauthorizedError
}

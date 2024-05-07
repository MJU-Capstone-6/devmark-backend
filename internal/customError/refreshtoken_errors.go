package customerror

import "net/http"

var TOKEN_CREATE_FAILED = NewError(http.StatusInternalServerError, "Token Creation failed", "301")

func TokenCreationFailed(err error) error {
	tokenCreationFailedError := TOKEN_CREATE_FAILED
	tokenCreationFailedError.Data = err.Error()
	return tokenCreationFailedError
}

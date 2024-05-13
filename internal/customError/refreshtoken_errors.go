package customerror

import "net/http"

var TOKEN_CREATE_FAILED = NewError(http.StatusUnprocessableEntity, "Token Creation failed", "301")

var TOKEN_UPDATE_FAILED = NewError(http.StatusUnprocessableEntity, "Token Update failed", "302")

var TOKEN_NOT_FOUND = NewError(http.StatusNotFound, "Token Not Found", "303")

func TokenCreationFailed(err error) error {
	tokenCreationFailedError := TOKEN_CREATE_FAILED
	tokenCreationFailedError.Data = err.Error()
	return tokenCreationFailedError
}

func TokenUpdateFailed(err error) error {
	tokenUpdateFailed := TOKEN_UPDATE_FAILED
	tokenUpdateFailed.Data = err.Error()
	return tokenUpdateFailed
}

func TokenNotFound(err error) error {
	tokenNotFound := TOKEN_UPDATE_FAILED
	tokenNotFound.Data = err.Error()
	return tokenNotFound
}

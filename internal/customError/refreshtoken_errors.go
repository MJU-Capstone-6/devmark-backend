package customerror

import "net/http"

var REFRESH_TOKEN_CREATE_FAILED = NewError(http.StatusUnprocessableEntity, "Token Creation failed", "301")

var REFRESH_TOKEN_UPDATE_FAILED = NewError(http.StatusUnprocessableEntity, "Token Update failed", "302")

var REFRESH_TOKEN_NOT_FOUND = NewError(http.StatusNotFound, "Token Not Found", "303")

func RefreshTokenCreationFailed(err error) error {
	tokenCreationFailedError := REFRESH_TOKEN_CREATE_FAILED
	tokenCreationFailedError.Data = err.Error()
	return tokenCreationFailedError
}

func RefreshTokenUpdateFailed(err error) error {
	tokenUpdateFailed := REFRESH_TOKEN_UPDATE_FAILED
	tokenUpdateFailed.Data = err.Error()
	return tokenUpdateFailed
}

func RefreshTokenNotFound(err error) error {
	tokenNotFound := REFRESH_TOKEN_NOT_FOUND
	tokenNotFound.Data = err.Error()
	return tokenNotFound
}

package customerror

import "net/http"

var TOKEN_NOT_PROVIDED = NewError(http.StatusUnauthorized, "Access key must be provide.", "101")

var TOKEN_NOT_VALID = NewError(http.StatusUnauthorized, "Token is not valid.", "102")

func TokenNotProvidedError(err error) error {
	tokenNotProvidedError := TOKEN_NOT_PROVIDED
	tokenNotProvidedError.Data = err.Error()
	return tokenNotProvidedError
}

func TokenNotValidError(err error) error {
	tokenNotValid := TOKEN_NOT_VALID
	tokenNotValid.Data = err.Error()
	return tokenNotValid
}

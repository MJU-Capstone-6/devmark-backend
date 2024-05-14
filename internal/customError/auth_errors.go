package customerror

import "net/http"

var TOKEN_NOT_PROVIDED = NewError(http.StatusUnauthorized, "Access key must be provide.", "101")

var TOKEN_NOT_VALID = NewError(http.StatusUnauthorized, "Token is not valid.", "102")

var TOKEN_SIGN_FAIL = NewError(http.StatusUnauthorized, "Token sign fail.", "103")

var TOKEN_VERIFY_FAIL = NewError(http.StatusUnauthorized, "Token verify fail.", "104")

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

func TokenSignFail(err error) error {
	tokenSignFail := TOKEN_SIGN_FAIL
	tokenSignFail.Data = err.Error()
	return tokenSignFail
}

func TokenVerifyFail(err error) error {
	tokenVerifyFail := TOKEN_VERIFY_FAIL
	tokenVerifyFail.Data = err.Error()
	return tokenVerifyFail
}

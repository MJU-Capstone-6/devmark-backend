package customerror

import "net/http"

var CODE_VERIFY_FAILED = NewError(http.StatusNotAcceptable, "Code is not valid. please try again", "501")

var CODE_NOT_PROVIDE = NewError(http.StatusBadRequest, "Code is not provide. please try again", "501")

func CodeVerifyFailedErr(err error) error {
	codeVerifyFailErr := CODE_VERIFY_FAILED
	codeVerifyFailErr.Data = err
	return codeVerifyFailErr
}

func CodeNotProvide(err error) error {
	codeNotProvide := CODE_NOT_PROVIDE
	codeNotProvide.Data = err
	return codeNotProvide
}

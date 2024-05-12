package customerror

import "net/http"

var CODE_VERIFY_FAILED = NewError(http.StatusNotAcceptable, "Code is not valid. please try again", "501")

func CodeVerifyFailedErr(err error) error {
	codeVerifyFailErr := CODE_VERIFY_FAILED
	codeVerifyFailErr.Data = err
	return codeVerifyFailErr
}

package customerror

import "net/http"

var CODE_VERIFY_FAILED = NewError(http.StatusNotAcceptable, "Code is not valid. please try again", "501")

var CODE_NOT_PROVIDE = NewError(http.StatusBadRequest, "Code is not provide. please try again", "502")

var CODE_CREATION_FAIL = NewError(http.StatusUnprocessableEntity, "Code creation fail. please try again", "503")

var CODE_NOT_FOUND = NewError(http.StatusNotFound, "Code not found.", "504")

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

func CodeCreationFail(err error) error {
	codeCreationFail := CODE_CREATION_FAIL
	codeCreationFail.Data = err
	return codeCreationFail
}

func CodeNotFound(err error) error {
	codeNotFound := CODE_NOT_FOUND
	codeNotFound.Data = err
	return codeNotFound
}

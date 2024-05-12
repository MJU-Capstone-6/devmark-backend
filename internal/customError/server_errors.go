package customerror

import "net/http"

var INTERNAL_SERVER_ERROR = NewError(http.StatusInternalServerError, "Internal Server Error", "001")

var INVALID_PARAM = NewError(http.StatusBadRequest, "Param is not valid", "002")

func InternalServerError(err error) error {
	internalServerError := INTERNAL_SERVER_ERROR
	internalServerError.Data = err
	return internalServerError
}

func InvalidParamError(err error) error {
	invalidParamError := INVALID_PARAM
	invalidParamError.Data = err
	return invalidParamError
}

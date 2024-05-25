package customerror

import "net/http"

var INTERNAL_SERVER_ERROR = NewError(http.StatusInternalServerError, "Internal Server Error", "001")

var INVALID_PARAM = NewError(http.StatusBadRequest, "Param is not valid", "002")

var (
	PAGE_NOT_FOUND     = NewError(http.StatusNotFound, "Page Not Found", "003")
	METHOD_NOT_ALLOWED = NewError(http.StatusMethodNotAllowed, "Method not allowed", "004")
)

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

func PageNotFound(err error) error {
	pageNotFound := PAGE_NOT_FOUND
	pageNotFound.Data = err
	return pageNotFound
}

func MethodNotAllowed(err error) error {
	methodNotAllowed := METHOD_NOT_ALLOWED
	methodNotAllowed.Data = err
	return methodNotAllowed
}

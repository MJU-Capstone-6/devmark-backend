package customerror

import "net/http"

var INTERNAL_SERVER_ERROR = NewError(http.StatusInternalServerError, "Internal Server Error", "001")

var NOT_FOUND = NewError(http.StatusNotFound, "Resources not found", "002")

func InternalServerError(err error) error {
	internalServerError := INTERNAL_SERVER_ERROR
	internalServerError.Data = err
	return internalServerError
}

package customerror

import "net/http"

var USER_NOT_FOUND = NewError(http.StatusNotFound, "User not found", "601")

func UserNotFound(err error) error {
	userNotFound := USER_NOT_FOUND
	userNotFound.Data = err
	return userNotFound
}

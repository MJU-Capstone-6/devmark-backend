package customerror

import "net/http"

var USER_NOT_FOUND = NewError(http.StatusNotFound, "User not found", "601")

var USER_CREATION_FAILED = NewError(http.StatusUnprocessableEntity, "User creation fail", "602")

var USER_UPDATE_FAILED = NewError(http.StatusUnprocessableEntity, "User update fail", "602")

func UserNotFound(err error) error {
	userNotFound := USER_NOT_FOUND
	userNotFound.Data = err
	return userNotFound
}

func UserCreationFail(err error) error {
	userCreationFail := USER_CREATION_FAILED
	userCreationFail.Data = err
	return userCreationFail
}

func UserUpdateFail(err error) error {
	userUpdateFail := USER_UPDATE_FAILED
	userUpdateFail.Data = err
	return userUpdateFail
}

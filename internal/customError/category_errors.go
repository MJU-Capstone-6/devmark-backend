package customerror

import "net/http"

var (
	CATEGORY_BODY_NOT_PROVIDE = NewError(http.StatusBadRequest, "category body info is not provided.", "801")
	CATEGORY_CREATION_FAIL    = NewError(http.StatusUnprocessableEntity, "category creation fail.", "802")
	CATEGORY_UPDATE_FAIL      = NewError(http.StatusUnprocessableEntity, "category update fail.", "803")
	CATEGORY_DELETE_FAIL      = NewError(http.StatusUnprocessableEntity, "category update fail.", "804")
	CATEGORY_NOT_FOUND        = NewError(http.StatusNotFound, "category not found.", "805")
)

func CategoryBodyNotProvide(err error) error {
	categoryBodyNotProvide := CATEGORY_BODY_NOT_PROVIDE
	categoryBodyNotProvide.Data = err
	return categoryBodyNotProvide
}

func CategoryNotFound(err error) error {
	categoryNotFound := CATEGORY_NOT_FOUND
	categoryNotFound.Data = err
	return categoryNotFound
}

func CategoryCreationFail(err error) error {
	categoryCreationFail := CATEGORY_CREATION_FAIL
	categoryCreationFail.Data = err
	return categoryCreationFail
}

func CategoryUpdateFail(err error) error {
	categoryUpdateFail := CATEGORY_UPDATE_FAIL
	categoryUpdateFail.Data = err
	return categoryUpdateFail
}

func CategoryDeleteFail(err error) error {
	categoryDeleteFail := CATEGORY_DELETE_FAIL
	categoryDeleteFail.Data = err
	return categoryDeleteFail
}

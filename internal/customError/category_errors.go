package customerror

import "net/http"

var CATEGORY_BODY_NOT_PROVIDE = NewError(http.StatusBadRequest, "category body info is not provided.", "801")

func CategoryBodyNotProvide(err error) error {
	categoryBodyNotProvide := CATEGORY_BODY_NOT_PROVIDE
	categoryBodyNotProvide.Data = err
	return categoryBodyNotProvide
}

package customerror

import "net/http"

var (
	COMMENT_CREATION_FAIL = NewError(http.StatusUnprocessableEntity, "comment creation fail.", "1001")
	COMMENT_UPDATE_FAIL   = NewError(http.StatusUnprocessableEntity, "comment update fail.", "1002")
	COMMENT_DELETE_FAIL   = NewError(http.StatusUnprocessableEntity, "comment update fail.", "1003")

	COMMENT_NOT_FOUND = NewError(http.StatusNotFound, "comment not found.", "1004")

	COMMENT_NOT_ALLOWED = NewError(http.StatusUnauthorized, "comment modification not allowed.", "1005")

	COMMENT_PARAM_NOT_VALID = NewError(http.StatusBadRequest, "comment param is not valid.", "1006")
)

func CommentCreationFail(err error) error {
	commentCreationFail := COMMENT_CREATION_FAIL
	commentCreationFail.Data = err
	return commentCreationFail
}

func CommentUpdateFail(err error) error {
	commentUpdateFail := COMMENT_UPDATE_FAIL
	commentUpdateFail.Data = err
	return commentUpdateFail
}

func CommentDeleteFail(err error) error {
	commentDeleteFail := COMMENT_DELETE_FAIL
	commentDeleteFail.Data = err
	return commentDeleteFail
}

func CommentNotFound(err error) error {
	commentNotFound := COMMENT_NOT_FOUND
	commentNotFound.Data = err
	return commentNotFound
}

func CommentNotAllowed(err error) error {
	commentNotAllowed := COMMENT_NOT_ALLOWED
	commentNotAllowed.Data = err
	return commentNotAllowed
}

func CommentParamNotValid(err error) error {
	commentParamNotValid := COMMENT_PARAM_NOT_VALID
	commentParamNotValid.Data = err
	return commentParamNotValid
}

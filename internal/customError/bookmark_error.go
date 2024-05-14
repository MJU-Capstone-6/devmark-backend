package customerror

import "net/http"

var BOOKMARK_CREATION_FAIL = NewError(http.StatusUnprocessableEntity, "Bookmark Creation fail.", "901")

var BOOKMARK_UPDATE_FAIL = NewError(http.StatusUnprocessableEntity, "Bookmark Update fail.", "902")

var BOOKMARK_DELETE_FAIL = NewError(http.StatusUnprocessableEntity, "Bookmark Delete fail.", "903")

var BOOKMARK_NOT_FOUND = NewError(http.StatusNotFound, "Bookmark Update fail.", "904")

func BookmarkCreationFail(err error) error {
	bookmarkCreationFail := BOOKMARK_CREATION_FAIL
	bookmarkCreationFail.Data = err
	return bookmarkCreationFail
}

func BookmarkUpdateFail(err error) error {
	bookmarkUpdateFail := BOOKMARK_UPDATE_FAIL
	bookmarkUpdateFail.Data = err
	return bookmarkUpdateFail
}

func BookmarkDeleteFail(err error) error {
	bookmarkDeleteFail := BOOKMARK_DELETE_FAIL
	bookmarkDeleteFail.Data = err
	return bookmarkDeleteFail
}

func BookmarkNotFound(err error) error {
	bookmarkNotFound := BOOKMARK_NOT_FOUND
	bookmarkNotFound.Data = err
	return bookmarkNotFound
}

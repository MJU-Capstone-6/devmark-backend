package customerror

import "net/http"

var WORKSPACECODE_NOT_FOUND = NewError(http.StatusNotFound, "Workspace Code not found", "1101")

var WORKSPACECODE_CREATION_FAIL = NewError(http.StatusUnprocessableEntity, "Workspace creation fail", "1102")

var WORKSPACECODE_UPDATE_FAIL = NewError(http.StatusUnprocessableEntity, "Workspace update fail", "1103")

func WorkspaceCodeNotFound(err error) error {
	workspaceCodeNotFound := WORKSPACECODE_NOT_FOUND
	workspaceCodeNotFound.Data = err
	return workspaceCodeNotFound
}

func WorkspaceCodeCreationFail(err error) error {
	workspaceCodeCreationFail := WORKSPACECODE_CREATION_FAIL
	workspaceCodeCreationFail.Data = err
	return workspaceCodeCreationFail
}

func WorkspaceCodeUpdateFail(err error) error {
	workspaceCodeUpdateFail := WORKSPACECODE_UPDATE_FAIL
	workspaceCodeUpdateFail.Data = err
	return workspaceCodeUpdateFail
}

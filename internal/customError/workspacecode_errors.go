package customerror

import "net/http"

var WORKSPACECODE_NOT_FOUND = NewError(http.StatusNotFound, "Workspace Code not found", "1101")

func WorkspaceCodeNotFound(err error) error {
	workspaceCodeNotFound := WORKSPACECODE_NOT_FOUND
	workspaceCodeNotFound.Data = err
	return workspaceCodeNotFound
}

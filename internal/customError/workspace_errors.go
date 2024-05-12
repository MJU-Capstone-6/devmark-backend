package customerror

import "net/http"

var WORKSPACE_NOT_FOUND = NewError(http.StatusNotFound, "Workspace not found", "401")

func WorkspaceNotFoundErr(err error) error {
	workspaceNotFound := WORKSPACE_NOT_FOUND
	workspaceNotFound.Data = err
	return workspaceNotFound
}

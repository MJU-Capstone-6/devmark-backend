package customerror

import "net/http"

var WORKSPACE_NOT_FOUND = NewError(http.StatusNotFound, "Workspace not found", "401")

var WORKSPACE_JOIN_FAIL = NewError(http.StatusBadRequest, "Workspace join fail", "402")

var WORKSPACE_UPDATE_FAIL = NewError(http.StatusUnprocessableEntity, "Workspace update fail", "403")

var WORKSPACE_DELETE_FAIL = NewError(http.StatusUnprocessableEntity, "Workspace delete fail", "404")

var WORKSPACE_CREATE_FAIL = NewError(http.StatusUnprocessableEntity, "Workspace delete fail", "404")

func WorkspaceNotFoundErr(err error) error {
	workspaceNotFound := WORKSPACE_NOT_FOUND
	workspaceNotFound.Data = err
	return workspaceNotFound
}

func WorkspaceJoinFailErr(err error) error {
	workspaceJoinFail := WORKSPACE_JOIN_FAIL
	workspaceJoinFail.Data = err
	return workspaceJoinFail
}

func WorkspaceUpdateFail(err error) error {
	workspaceUpdateFail := WORKSPACE_UPDATE_FAIL
	workspaceUpdateFail.Data = err
	return workspaceUpdateFail
}

func WorkspaceCreateFail(err error) error {
	workspaceCreateFail := WORKSPACE_CREATE_FAIL
	workspaceCreateFail.Data = err
	return workspaceCreateFail
}

func WorkspaceDeleteFail(err error) error {
	workspaceDeleteFail := WORKSPACE_DELETE_FAIL
	workspaceDeleteFail.Data = err
	return workspaceDeleteFail
}

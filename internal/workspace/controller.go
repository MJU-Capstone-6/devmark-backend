package workspace

import (
	"net/http"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type WorkspaceController struct {
	WorkspaceService interfaces.IWorkspaceService
}

func (w *WorkspaceController) TestController(ctx echo.Context) error {
	return nil
}

// ViewWorkspaceController godoc
//
//	@summary	워크스페이스 조회
//	@schemes
//	@description	워크스페이스를 조회합니다.
//	@tags			workspace
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"Workspace id"
//	@success		200	{object}	repository.FindWorkspaceRow
//	@failure		401	{object}	customerror.CustomError
//	@failure		404	{object}	customerror.CustomError
//	@failure		500	{object}	customerror.CustomError
//	@router			/api/v1/workspace/:id [GET]
func (w *WorkspaceController) ViewWorkspaceController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	workspaceInfo, err := w.WorkspaceService.FindById(*id)
	if err != nil {
		return customerror.WorkspaceNotFoundErr(err)
	}

	return ctx.JSON(http.StatusOK, workspaceInfo)
}

// UpdateWorkspaceController godoc
//
//	@summary	워크스페이스 업데이트
//	@schemes
//	@description	워크스페이스 정보를 업데이트 합니다.
//	@tags			workspace
//	@accept			json
//	@produce		json
//	@param			id		path		int									true	"Workspace id"
//	@param			body	body		repository.UpdateWorkspaceParams	true	"Workspace info"
//	@success		200		{object}	repository.Workspace
//	@failure		401		{object}	customerror.CustomError
//	@failure		404		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/workspace/:id [PUT]
func (w *WorkspaceController) UpdateWorkspaceController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	var params repository.UpdateWorkspaceParams
	err = ctx.Bind(&params)
	params.ID = int64(*id)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}

	workspace, err := w.WorkspaceService.Update(params)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}
	return ctx.JSON(http.StatusOK, workspace)
}

// DeleteWorkspaceController godoc
//
//	@summary	워크스페이스 삭제
//	@schemes
//	@description	워크스페이스를 삭제합니다.
//	@tags			workspace
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"Workspace id"
//	@success		200	{object}	repository.Workspace
//	@failure		401	{object}	customerror.CustomError
//	@failure		404	{object}	customerror.CustomError
//	@failure		500	{object}	customerror.CustomError
//	@router			/api/v1/workspace/:id [DELETE]
func (w *WorkspaceController) DeleteWorkspaceController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	err = w.WorkspaceService.Delete(*id)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}
	return nil
}

// CreateWorkspace godoc
//
//	@summary	워크스페이스 생성
//	@schemes
//	@description	워크스페이스를 생성합니다.
//	@tags			workspace
//	@accept			json
//	@produce		json
//	@param			body	body		CreateWorkspaceParam	true	"Workspace id"
//	@success		200		{object}	repository.Workspace
//	@failure		401		{object}	customerror.CustomError
//	@failure		404		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/workspace [POST]
func (w *WorkspaceController) CreateWorkspaceController(ctx echo.Context) error {
	var param repository.CreateWorkspaceParams
	err := ctx.Bind(&param)
	if err != nil {
		return customerror.InternalServerError(err)
	}

	user, err := utils.GetAuthUser(ctx)
	if err != nil {
		return err
	}

	workspace, err := w.WorkspaceService.Create(int(user.ID), param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, workspace)
}

// JoinWorkspace godoc
//
//	@summary	워크스페이스 참가
//	@schemes
//	@description	워크스페이스에 참가합니다.
//	@tags			workspace
//	@accept			json
//	@produce		json
//	@param			body	body		JoinWorkspaceParam	true	"Workspace join info"
//	@success		200		{object}	repository.Workspace
//	@failure		401		{object}	customerror.CustomError
//	@failure		404		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/workspace/join [POST]
func (w *WorkspaceController) JoinWorkspaceController(ctx echo.Context) error {
	var param JoinWorkspaceParam
	err := ctx.Bind(&param)
	if err != nil {
		return customerror.CodeNotProvide(err)
	}

	if user, ok := ctx.Get(constants.USER_CONTEXT_KEY).(*repository.User); ok {
		joinWorkspaceParam := repository.JoinWorkspaceParams{
			UserID: user.ID,
		}
		err = w.WorkspaceService.Join(param.Code, joinWorkspaceParam)
		if err != nil {
			if _, ok := err.(*customerror.CustomError); ok {
				return responses.NotAcceptable(ctx, err)
			} else {
				return responses.InternalServer(ctx, customerror.InternalServerError(err))
			}
		}
	}
	return nil
}

func InitWorkspaceController() *WorkspaceController {
	return &WorkspaceController{}
}

func (w WorkspaceController) WithWorkspaceService(service interfaces.IWorkspaceService) WorkspaceController {
	w.WorkspaceService = service
	return w
}

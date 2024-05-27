package workspace

import (
	"net/http"
	"strings"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type WorkspaceController struct {
	WorkspaceService     interfaces.IWorkspaceService
	WorkspaceCodeService interfaces.IWorkspaceCodeService
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
		return err
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
//	@success		200	{object}	responses.OkResponse
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
	return ctx.JSON(http.StatusOK, responses.OkResponse{Ok: true})
}

// CreateWorkspace godoc
//
//	@summary	워크스페이스 생성
//	@schemes
//	@description	워크스페이스를 생성합니다.
//	@tags			workspace
//	@accept			json
//	@produce		json
//	@param			body	body		repository.CreateWorkspaceParams	true	"Workspace Body"
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

	if user, ok := ctx.Get(constants.USER_CONTEXT_KEY).(*repository.FindUserByIdRow); ok {
		joinWorkspaceParam := repository.JoinWorkspaceParams{
			UserID: user.ID,
		}
		err = w.WorkspaceService.Join(param.Code, joinWorkspaceParam)
		if err != nil {
			if _, ok := err.(*customerror.CustomError); ok {
				return err
			} else {
				return customerror.InternalServerError(err)
			}
		}
		return ctx.JSON(http.StatusOK, responses.OkResponse{Ok: true})
	}
	return customerror.InternalServerError(err)
}

// FindWorkspaceCategoriesController godoc
//
//	@summary	워크스페이스 카테고리 조회
//	@schemes
//	@description	워크스페이스의 카테고리를 조회합니다.
//	@tags			workspace
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"Workspace id"
//	@success		200	{object}	[]repository.Category
//	@failure		401	{object}	customerror.CustomError
//	@failure		500	{object}	customerror.CustomError
//	@router			/api/v1/workspace/:id/category [GET]
func (w *WorkspaceController) FindWorkspaceCategoriesController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	categories, _ := w.WorkspaceService.FindCategoriesById(*id)
	return ctx.JSON(http.StatusOK, categories)
}

// FindWorkspaceCategoryBookmark godoc
//
//	@summary	워크스페이스 카테고리 - 북마크 조회
//	@schemes
//	@description	워크스페이스의 카테고리내의 북마크를  조회합니다.
//	@tags			workspace
//	@accept			json
//	@produce		json
//	@param			workspace_id	path		int	true	"Workspace id"
//	@param			category_id		path		int	true	"Category id"
//	@success		200				{object}	[]repository.Bookmark
//	@failure		401				{object}	customerror.CustomError
//	@failure		500				{object}	customerror.CustomError
//	@router			/api/v1/workspace/:workspace_id/category/:category_id [GET]
func (w *WorkspaceController) FindWorkspaceCategoryBookmark(ctx echo.Context) error {
	workspace_id, err := utils.ParseURLParam(ctx, "workspace_id")
	if err != nil {
		return err
	}
	parsed_workspace_id := int64(*workspace_id)

	category_id, err := utils.ParseURLParam(ctx, "category_id")
	if err != nil {
		return err
	}

	parsed_category_id := int64(*category_id)
	param := repository.FindWorkspaceCategoryBookmarkParams{
		WorkspaceID: &parsed_workspace_id,
		CategoryID:  &parsed_category_id,
	}
	bookmarks, _ := w.WorkspaceService.FindCategoryBookmark(param)
	return ctx.JSON(http.StatusOK, bookmarks)
}

// RegisterCategoryToWorkspaceController godoc
//
//	@summary	워크스페이스 카테고리 등록
//	@schemes
//	@description	워크스페이스에 카테고리를 등록합니다.
//	@tags			workspace
//	@accept			json
//	@produce		json
//	@param			workspace_id	path		int	true	"Workspace id"
//	@param			category_id		path		int	true	"Category id"
//	@success		200				{object}	responses.OkResponse
//	@failure		400				{object}	customerror.CustomError
//	@failure		401				{object}	customerror.CustomError
//	@failure		500				{object}	customerror.CustomError
//	@router			/api/v1/workspace/:workspace_id/category/:category_id [POST]
func (w *WorkspaceController) RegisterCategoryToWorkspaceController(ctx echo.Context) error {
	workspace_id, err := utils.ParseURLParam(ctx, "workspace_id")
	if err != nil {
		return err
	}
	parsed_workspace_id := int64(*workspace_id)

	category_id, err := utils.ParseURLParam(ctx, "category_id")
	if err != nil {
		return err
	}

	parsed_category_id := int64(*category_id)

	param := repository.RegisterCategoryToWorkspaceParams{
		WorkspaceID: parsed_workspace_id,
		CategoryID:  parsed_category_id,
	}

	err = w.WorkspaceService.RegisterCategory(param)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, responses.OkResponse{Ok: true})
}

// SearchBookmarkController godoc
//
//	@summary	워크스페이스 북마크 검색
//	@schemes
//	@description	워크스페이스의 북마크를 검색합니다.
//	@tags			workspace
//	@accept			json
//	@produce		json
//	@param			workspace_id	path		int		true	"Workspace id"
//	@param			user			query		string	false	"users id"
//	@param			category		query		string	false	"categories id"
//	@success		200				{object}	[]repository.SearchWorkspaceBookmarkRow
//	@failure		400				{object}	customerror.CustomError
//	@failure		401				{object}	customerror.CustomError
//	@failure		500				{object}	customerror.CustomError
//	@router			/api/v1/workspace/:id/bookmark [GET]
func (w *WorkspaceController) SearchBookmarkController(ctx echo.Context) error {
	user_ids := &[]int64{}
	category_ids := &[]int64{}
	workspace_id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}

	usersParam := ctx.QueryParam("user")
	if usersParam != "" {
		user_ids, err = utils.SliceValueIntoNum(strings.Split(usersParam, ","))
		if err != nil {
			return customerror.InvalidParamError(err)
		}
	}

	categoryParam := ctx.QueryParam("category")
	if categoryParam != "" {
		category_ids, err = utils.SliceValueIntoNum(strings.Split(categoryParam, ","))
		if err != nil {
			return customerror.InvalidParamError(err)
		}
	}

	parsedWorkspaceID := int64(*workspace_id)
	param := repository.SearchWorkspaceBookmarkParams{
		WorkspaceID: &parsedWorkspaceID,
		UserIds:     *user_ids,
		CategoryIds: *category_ids,
	}

	bookmarks, err := w.WorkspaceService.SearchBookmark(param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, *bookmarks)
}

// CreateWorkspaceCodeController godoc
//
//	@summary	워크스페이스 코드 생성
//	@schemes
//	@description	워크스페이스의 코드를 생성합니다.
//	@tags			workspace
//	@accept			json
//	@produce		json
//	@param			workspace_id	path		int	true	"Workspace id"
//	@success		200				{object}	repository.WorkspaceCode
//	@failure		400				{object}	customerror.CustomError
//	@failure		401				{object}	customerror.CustomError
//	@failure		500				{object}	customerror.CustomError
//	@router			/api/v1/workspace/:id/code [POST]
func (w *WorkspaceController) CreateWorkspaceCodeController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	parsedID := int64(*id)
	param := repository.CreateWorkspaceCodeParams{
		WorkspaceID: &parsedID,
	}
	workspaceCode, err := w.WorkspaceCodeService.Create(param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, workspaceCode)
}

func (w *WorkspaceController) FindWorkspaceInfoController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	workspaceInfo, err := w.WorkspaceService.FindInfoById(*id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, workspaceInfo)
}

func InitWorkspaceController() *WorkspaceController {
	return &WorkspaceController{}
}

func (w WorkspaceController) WithWorkspaceService(service interfaces.IWorkspaceService) WorkspaceController {
	w.WorkspaceService = service
	return w
}

func (w WorkspaceController) WithWorkspaceCodeService(service interfaces.IWorkspaceCodeService) WorkspaceController {
	w.WorkspaceCodeService = service
	return w
}

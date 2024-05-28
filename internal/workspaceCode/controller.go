package workspacecode

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/request"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type WorkspaceCodeController struct {
	WorkspaceCodeService interfaces.IWorkspaceCodeService
}

// PredictCategoryController godoc
//
//	@summary	카테고리 예측 API
//	@schemes
//	@description	카테고리 예측 API 입니다. 카테고리가 존재하지 않을 시 자동으로 카테고리가 생성됩니다.
//	@tags			workspaceCode
//	@accept			json
//	@produce		json
//	@param			body	body		request.PredictCategoryBody	true	"PredictCategory Body"
//	@success		200		{object}	repository.Bookmark
//	@failure		400		{object}	customerror.CustomError
//	@failure		401		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/code/predict [POST]
func (w *WorkspaceCodeController) PredictCategoryController(ctx echo.Context) error {
	var body request.PredictCategoryBody
	bytes, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return customerror.InternalServerError(err)
	}
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return customerror.InternalServerError(err)
	}
	domain := ctx.QueryParam("domain")
	if domain == "" {
		return customerror.InvalidParamError(errors.New("domain query param must be provided"))
	}
	header := ctx.Request().Header.Get(constants.USER_AGENT_HEADER)
	param := request.PredictCategoryParam{
		Code:   body.Code,
		Link:   body.Link,
		Domain: domain,
	}
	bookmark, err := w.WorkspaceCodeService.PredictCategory(param, header)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, bookmark)
}

// FindWorkspaceCodeController godoc
//
//	@summary	코드 검증 API
//	@schemes
//	@description	코드가 존재하면 해당하는 코드의 Workspace Name을 반환합니다.
//	@tags			workspaceCode
//	@accept			text/plain
//	@produce		json
//	@param			code	query		string	true	"code query"
//	@success		200		{object}	responses.FindWorkspaceResponse
//	@failure		400		{object}	customerror.CustomError
//	@failure		401		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/code [GET]
func (w *WorkspaceCodeController) FindWorkspaceCodeController(ctx echo.Context) error {
	param := ctx.QueryParam("code")
	if param == "" {
		return customerror.InvalidParamError(errors.New("query param 'code' must be provided"))
	}
	row, err := w.WorkspaceCodeService.FindByCode(param)
	if err != nil {
		return err
	}
	response := responses.FindWorkspaceResponse{
		WorkspaceName: *row.Workspace.Name,
	}
	return ctx.JSON(http.StatusOK, response)
}

func InitWorkspaceCodeController() *WorkspaceCodeController {
	return &WorkspaceCodeController{}
}

func (w WorkspaceCodeController) WithWorkspaceCodeService(service interfaces.IWorkspaceCodeService) WorkspaceCodeController {
	w.WorkspaceCodeService = service
	return w
}

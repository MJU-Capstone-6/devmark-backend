package workspace

import (
	"log"
	"net/http"

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

func (w *WorkspaceController) ViewWorkspaceController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	workspaceInfo, err := w.WorkspaceService.FindById(*id)
	if err != nil {
		return responses.NotFound(ctx, customerror.WorkspaceNotFoundErr(err))
	}

	return ctx.JSON(http.StatusOK, workspaceInfo)
}

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

func (w *WorkspaceController) CreateWorkspaceController(ctx echo.Context) error {
	var param CreateWorkspaceParam
	err := ctx.Bind(&param)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}
	workspace, err := w.WorkspaceService.Create(param.Name)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}
	return ctx.JSON(http.StatusOK, workspace)
}

func (w *WorkspaceController) JoinWorkspaceController(ctx echo.Context) error {
	var param JoinWorkspaceParam
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	err = ctx.Bind(&param)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}

	if user, ok := ctx.Get("user").(*repository.User); ok {
		joinWorkspaceParam := repository.JoinWorkspaceParams{
			WorkspaceID: int64(*id),
			UserID:      user.ID,
		}
		err = w.WorkspaceService.Join(param.Code, joinWorkspaceParam)
		log.Println(err)
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

package workspacecode

import (
	"net/http"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/request"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type WorkspaceCodeController struct {
	WorkspaceCodeService interfaces.IWorkspaceCodeService
}

func (w *WorkspaceCodeController) PredictCategoryController(ctx echo.Context) error {
	var body request.PredictCategoryBody
	err := ctx.Bind(&body)
	if err != nil {
		return customerror.InternalServerError(err)
	}
	domain := ctx.QueryParam("domain")
	user, err := utils.GetAuthUser(ctx)
	if err != nil {
		return err
	}
	param := request.PredictCategoryParam{
		Code:   body.Code,
		Link:   body.Link,
		Domain: domain,
		UserID: user.ID,
	}
	bookmark, err := w.WorkspaceCodeService.PredictCategory(param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, bookmark)
}

func InitWorkspaceCodeController() *WorkspaceCodeController {
	return &WorkspaceCodeController{}
}

func (w WorkspaceCodeController) WithWorkspaceCodeService(service interfaces.IWorkspaceCodeService) WorkspaceCodeController {
	w.WorkspaceCodeService = service
	return w
}

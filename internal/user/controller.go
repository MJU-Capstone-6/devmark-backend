package user

import (
	"log"
	"net/http"

	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService interfaces.IUserService
}

func (u *UserController) ViewOneUser(ctx echo.Context) error {
	name := ctx.Param("name")
	user, err := u.UserService.FindUserByUserName(&name)
	if err != nil {
		log.Println(err)
	}
	return ctx.JSON(http.StatusOK, user)
}

// ViewUserWorkspace godoc
//
//	@summary	유저가 참가한 워크스페이스의 리스트를 조회합니다.
//	@schemes
//	@description	유저가 참가한 워크스페이스의 리스트를 조회합니다.
//	@tags			user
//	@accept			json
//	@produce		json
//	@success		200	{object}	repository.UserWorkspaceView
//	@failure		500	{object}	customerror.CustomError
//	@router			/api/v1/user/:id/workspace [GET]
func (u *UserController) ViewUserWorkspace(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	workspace, err := u.UserService.FindJoinedWorkspace(*id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, workspace)
}

func InitController() *UserController {
	return &UserController{}
}

func (u UserController) WithUserService(service interfaces.IUserService) UserController {
	u.UserService = service
	return u
}

package user

import (
	"log"
	"net/http"

	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService         interfaces.IUserService
	NotificationService interfaces.INotificationService
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
//	@router			/api/v1/user/workspace [GET]
func (u *UserController) ViewUserWorkspace(ctx echo.Context) error {
	user, err := utils.GetAuthUser(ctx)
	if err != nil {
		return err
	}

	workspace, _ := u.UserService.FindJoinedWorkspace(int(user.ID))

	return ctx.JSON(http.StatusOK, workspace)
}

// ViewUserNotificationHistory godoc
//
//	@summary	유저의 알림 리스트를 조회합니다.
//	@schemes
//	@description	유저의 알림 리스트를 조회합니다.
//	@tags			user
//	@accept			json
//	@produce		json
//	@success		200	{object}	[]repository.FindUnreadNotificationHistoryRow
//	@failure		404 {object}	customerror.CustomError
//	@failure		500	{object}	customerror.CustomError
//	@router			/api/v1/user/notification [GET]
func (u *UserController) ViewUserNotificationHistory(ctx echo.Context) error {
	user, err := utils.GetAuthUser(ctx)
	if err != nil {
		return err
	}
	notificationHistory, err := u.NotificationService.FindUnreadNotificationHistory(user.ID)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, notificationHistory)
}

func InitController() *UserController {
	return &UserController{}
}

func (u UserController) WithUserService(service interfaces.IUserService) UserController {
	u.UserService = service
	return u
}

func (u UserController) WithNotificationService(service interfaces.INotificationService) UserController {
	u.NotificationService = service
	return u
}

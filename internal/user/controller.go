package user

import (
	"log"
	"net/http"

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

func InitController() *UserController {
	return &UserController{}
}

func (u UserController) WithUserService(service interfaces.IUserService) UserController {
	u.UserService = service
	return u
}

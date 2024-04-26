package user

import (
	"context"
	"log"
	"net/http"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService UserService
}

func (u *UserController) FindUser(ctx echo.Context) error {
	user, err := u.UserService.UserRepository.FindUserByID(context.Background(), 1)
	if err != nil {
		return err
	}
	log.Println(user)
	return nil
}

func (u *UserController) CreateUser(ctx echo.Context) error {
	param := new(repository.CreateParams)
	if err := ctx.Bind(&param); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	user, err := u.UserService.Create(context.Background(), *param)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, user)
}

func InitController(conn *pgx.Conn) *UserController {
	userService := InitUserService(conn)
	return &UserController{UserService: *userService}
}

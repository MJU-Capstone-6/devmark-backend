package user

import (
	"context"
	"log"
	"net/http"

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
	user, err := u.UserService.UserRepository.Create(context.Background(), "test")
	if err != nil {
		log.Println(err)
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}

func InitController(conn *pgx.Conn) *UserController {
	userService := InitUserService(conn)
	return &UserController{UserService: *userService}
}

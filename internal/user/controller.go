package user

import (
	"github.com/jackc/pgx/v5"
)

type UserController struct {
	UserService UserService
}

func InitController(conn *pgx.Conn) *UserController {
	userService := InitUserService(conn)
	return &UserController{UserService: *userService}
}

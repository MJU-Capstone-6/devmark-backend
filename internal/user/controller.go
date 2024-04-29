package user

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/jackc/pgx/v5"
)

type UserController struct {
	UserService UserService
}

func InitController(conn *pgx.Conn) *UserController {
	repo := repository.New(conn)
	userService := InitUserService(repo)
	return &UserController{UserService: *userService}
}

package user

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/jackc/pgx/v5"
)

type UserService struct {
	UserRepository *repository.Queries
}

func InitUserService(conn *pgx.Conn) *UserService {
	return &UserService{UserRepository: repository.New(conn)}
}

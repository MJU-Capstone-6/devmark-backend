package user

import (
	"context"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/jackc/pgx/v5"
)

type UserService struct {
	UserRepository *repository.Queries
}

func (u *UserService) FindUserByUserName(username *string) (*repository.User, error) {
	user, err := u.UserRepository.FindUserByUsername(context.Background(), username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func InitUserService(conn *pgx.Conn) *UserService {
	return &UserService{UserRepository: repository.New(conn)}
}

package user

import (
	"context"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/jackc/pgx/v5"
)

type UserService struct {
	UserRepository *repository.Queries
}

func (s *UserService) Create(ctx context.Context, params repository.CreateParams) (*repository.User, error) {
	user, err := s.UserRepository.Create(ctx, params)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func InitUserService(conn *pgx.Conn) *UserService {
	return &UserService{UserRepository: repository.New(conn)}
}

package interfaces

import (
	"context"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
)

//go:generate mockery --name IRepository
type IRepository interface {
	FindUserByUsername(context.Context, *string) (repository.User, error)
	CreateRefreshToken(context.Context, repository.CreateRefreshTokenParams) (repository.RefreshToken, error)
	FindRefreshTokenByUserID(context.Context, *int32) (repository.RefreshToken, error)
	UpdateRefreshToken(context.Context, repository.UpdateRefreshTokenParams) (repository.RefreshToken, error)
}

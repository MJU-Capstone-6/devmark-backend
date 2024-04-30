package interfaces

import (
	"context"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
)

//go:generate mockery --name IUserRepository
type IUserRepository interface {
	FindUserByUsername(context.Context, *string) (repository.User, error)
}

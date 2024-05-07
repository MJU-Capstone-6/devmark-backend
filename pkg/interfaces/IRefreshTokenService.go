package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

type IRefreshTokenService interface {
	Create(string) (*repository.RefreshToken, error)
	FindOneByUserID(id int) (*repository.RefreshToken, error)
}

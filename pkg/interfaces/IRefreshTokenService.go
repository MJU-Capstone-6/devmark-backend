package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

//go:generate mockery --name IRefreshTokenService
type IRefreshTokenService interface {
	CreateToken(string) (*repository.RefreshToken, error)
	UpdateToken(repository.UpdateRefreshTokenParams) (*repository.RefreshToken, error)
	FindOneByUserId(int) (*repository.RefreshToken, error)
}

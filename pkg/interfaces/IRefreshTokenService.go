package interfaces

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
)

//go:generate mockery --name IRefreshTokenService
type IRefreshTokenService interface {
	CreateToken(string) (*repository.RefreshToken, error)
	UpdateToken(repository.UpdateRefreshTokenParams) (*repository.RefreshToken, error)
	FindOneByUserId(int) (*repository.RefreshToken, error)
	RefreshTokens(string) (*responses.RefreshAccessTokenResponse, error)
}

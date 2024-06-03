package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

//go:generate mockery --name IRecommendLinkService
type IRecommendLinkService interface {
	Create(repository.CreateRecommendLinkParams) (*repository.RecommendLink, error)
}

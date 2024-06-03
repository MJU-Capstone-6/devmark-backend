package recommendlink

import (
	"context"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type RecommendLinkService struct {
	Repository interfaces.IRepository
}

func (r *RecommendLinkService) Create(param repository.CreateRecommendLinkParams) (*repository.RecommendLink, error) {
	recommendLink, err := r.Repository.CreateRecommendLink(context.Background(), param)
	if err != nil {
		return nil, customerror.RecommendLinkCreationFail(err)
	}
	return &recommendLink, nil
}

func InitRecommendLinkService() *RecommendLinkService {
	return &RecommendLinkService{}
}

func (r RecommendLinkService) WithRepository(repo interfaces.IRepository) RecommendLinkService {
	r.Repository = repo
	return r
}

package refreshtoken

import (
	"context"
	"log"
	"strconv"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type RefreshTokenService struct {
	Repository  interfaces.IRepository
	JwtService  interfaces.IJWTService
	UserService interfaces.IUserService
}

func (r *RefreshTokenService) CreateToken(token string) (*repository.RefreshToken, error) {
	log.Println(token)
	parsedToken, err := r.JwtService.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	parsedValue, err := strconv.Atoi(parsedToken.Get(constants.TOKEN_DATA_KEY))
	if err != nil {
		return nil, customerror.TokenNotValidError(err)
	}
	userId := int32(parsedValue)

	params := repository.CreateRefreshTokenParams{
		UserID: &userId,
		Token:  &token,
	}
	refreshToken, err := r.Repository.CreateRefreshToken(context.Background(), params)
	if err != nil {
		return nil, customerror.RefreshTokenCreationFailed(err)
	}
	return &refreshToken, nil
}

func (r *RefreshTokenService) UpdateToken(params repository.UpdateRefreshTokenParams) (*repository.RefreshToken, error) {
	updatedRefreshToken, err := r.Repository.UpdateRefreshToken(context.Background(), params)
	if err != nil {
		return nil, customerror.RefreshTokenUpdateFailed(err)
	}
	return &updatedRefreshToken, nil
}

func (r *RefreshTokenService) FindOneByUserId(id int) (*repository.RefreshToken, error) {
	userId := int32(id)
	refreshToken, err := r.Repository.FindRefreshTokenByUserID(context.Background(), &userId)
	if err != nil {
		return nil, customerror.RefreshTokenNotFound(err)
	}
	return &refreshToken, nil
}

func (r *RefreshTokenService) RefreshTokens(token string) (*responses.RefreshAccessTokenResponse, error) {
	if refreshToken, err := r.JwtService.VerifyToken(token); err != nil {
		return nil, customerror.TokenNotValidError(err)
	} else {
		userID, err := strconv.Atoi(refreshToken.Get(constants.TOKEN_DATA_KEY))
		if err != nil {
			return nil, customerror.InternalServerError(err)
		}
		accessToken, err := r.JwtService.GenerateToken(userID, constants.ACCESSTOKEN_EXPIRED_TIME)
		if err != nil {
			return nil, customerror.TokenSignFail(err)
		}
		newRefreshToken, err := r.JwtService.GenerateToken(userID, constants.REFRESH_TOKEN_EXPIRED_TIME)
		if err != nil {
			return nil, err
		}
		findRefreshToken, err := r.FindOneByUserId(userID)
		if err != nil {
			return nil, err
		}
		param := repository.UpdateRefreshTokenParams{
			Token: &newRefreshToken,
			ID:    findRefreshToken.ID,
		}
		_, err = r.UpdateToken(param)
		if err != nil {
			return nil, err
		}

		return &responses.RefreshAccessTokenResponse{
			AccessToken:  accessToken,
			RefreshToken: newRefreshToken,
		}, nil
	}
}

func InitRefreshTokenService(repo interfaces.IRepository, service interfaces.IJWTService) *RefreshTokenService {
	refreshService := RefreshTokenService{}.WithRepository(repo).WithJWTService(service)

	return &refreshService
}

func (r RefreshTokenService) WithRepository(repo interfaces.IRepository) RefreshTokenService {
	r.Repository = repo
	return r
}

func (r RefreshTokenService) WithJWTService(service interfaces.IJWTService) RefreshTokenService {
	r.JwtService = service
	return r
}

func (r RefreshTokenService) WithUserService(service interfaces.IUserService) RefreshTokenService {
	r.UserService = service
	return r
}

package auth

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type AuthService struct {
	Repository     interfaces.IRepository
	UserService    interfaces.IUserService
	JWTService     interfaces.IJWTService
	RefreshService interfaces.IRefreshTokenService
}

func (a *AuthService) KakaoSignUp(nickname string, provider string) (*responses.GetKakaoInfoResponse, error) {
	var userId int
	var user *repository.User
	var refreshToken *repository.RefreshToken

	user, err := a.UserService.FindUserByUserName(&nickname)
	if err != nil {
		params := repository.CreateUserParams{Username: &nickname, Provider: &provider}
		user, err = a.UserService.CreateUser(params)
		userId = int(user.ID)
		if err != nil {
			return nil, customerror.UserCreationFail(err)
		}

		tokenString, err := a.JWTService.GenerateToken(userId, constants.REFRESH_TOKEN_EXPIRED_TIME)
		if err != nil {
			return nil, customerror.TokenCreationFailed(err)
		}
		refreshToken, err = a.RefreshService.CreateToken(tokenString)
		if err != nil {
			return nil, customerror.TokenCreationFailed(err)
		}
		refreshTokenId := int32(refreshToken.ID)

		user.RefreshToken = &refreshTokenId
	} else {
		userId = int(user.ID)
		tokenString, err := a.JWTService.GenerateToken(userId, constants.REFRESH_TOKEN_EXPIRED_TIME)
		if err != nil {
			return nil, customerror.TokenCreationFailed(err)
		}
		refreshToken, err = a.RefreshService.FindOneByUserId(int(user.ID))
		if err != nil {
			return nil, customerror.TokenNotFound(err)
		}
		updateRefreshTokenParam := repository.UpdateRefreshTokenParams{
			Token: &tokenString,
			ID:    refreshToken.ID,
		}
		refreshToken, err = a.RefreshService.UpdateToken(updateRefreshTokenParam)
		if err != nil {
			return nil, customerror.TokenUpdateFailed(err)
		}
	}
	accessToken, err := a.JWTService.GenerateToken(userId, constants.ACCESSTOKEN_EXPIRED_TIME)
	if err != nil {
		return nil, customerror.TokenCreationFailed(err)
	}
	refreshTokenID := int32(refreshToken.ID)
	updateUserParam := repository.UpdateUserParams{
		RefreshToken: &refreshTokenID,
		Username:     user.Username,
		ID:           user.ID,
	}
	_, err = a.UserService.UpdateUser(updateUserParam)
	if err != nil {
		return nil, customerror.UserUpdateFail(err)
	}
	return &responses.GetKakaoInfoResponse{AccessToken: accessToken, RefreshToken: *refreshToken.Token}, nil
}

func InitAuthService(repo interfaces.IRepository, userService interfaces.IUserService, jwtService interfaces.IJWTService, refreshTokenService interfaces.IRefreshTokenService) *AuthService {
	authService := AuthService{}.WithUserService(userService).WithRepository(repo).WithJWTService(jwtService).WithRefreshTokenService(refreshTokenService)
	return &authService
}

func (a AuthService) WithRepository(repo interfaces.IRepository) AuthService {
	a.Repository = repo
	return a
}

func (a AuthService) WithUserService(service interfaces.IUserService) AuthService {
	a.UserService = service
	return a
}

func (a AuthService) WithJWTService(service interfaces.IJWTService) AuthService {
	a.JWTService = service
	return a
}

func (a AuthService) WithRefreshTokenService(service interfaces.IRefreshTokenService) AuthService {
	a.RefreshService = service
	return a
}

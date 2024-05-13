package auth

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type AuthService struct {
	Repository     interfaces.IRepository
	UserService    interfaces.IUserService
	JWTService     interfaces.IJWTService
	RefreshService interfaces.IRefreshTokenService
}

func (a *AuthService) KakaoSignUp(nickname string, provider string, ctx echo.Context) error {
	var userId int
	var user *repository.User
	var refreshToken *repository.RefreshToken

	user, err := a.UserService.FindUserByUserName(&nickname)
	if err != nil {
		params := repository.CreateUserParams{Username: &nickname, Provider: &provider}
		user, err = a.UserService.CreateUser(params)
		userId = int(user.ID)
		if err != nil {
			return responses.InternalServer(ctx, customerror.InternalServerError(err))
		}

		tokenString, err := a.JWTService.GenerateToken(userId, constants.REFRESH_TOKEN_EXPIRED_TIME)
		if err != nil {
			return responses.InternalServer(ctx, customerror.InternalServerError(err))
		}
		refreshToken, err = a.RefreshService.CreateToken(tokenString)
		if err != nil {
			return responses.InternalServer(ctx, customerror.TokenCreationFailed(err))
		}
		refreshTokenId := int32(refreshToken.ID)

		user.RefreshToken = &refreshTokenId
	} else {
		userId = int(user.ID)
		tokenString, err := a.JWTService.GenerateToken(userId, constants.REFRESH_TOKEN_EXPIRED_TIME)
		if err != nil {
			return responses.InternalServer(ctx, customerror.TokenCreationFailed(err))
		}
		refreshToken, err = a.RefreshService.FindOneByUserId(int(user.ID))
		if err != nil {
			return responses.InternalServer(ctx, customerror.TokenCreationFailed(err))
		}
		updateRefreshTokenParam := repository.UpdateRefreshTokenParams{
			Token: &tokenString,
			ID:    refreshToken.ID,
		}
		refreshToken, err = a.RefreshService.UpdateToken(updateRefreshTokenParam)
		if err != nil {
			return responses.InternalServer(ctx, customerror.TokenCreationFailed(err))
		}
	}
	accessToken, err := a.JWTService.GenerateToken(userId, constants.ACCESSTOKEN_EXPIRED_TIME)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}
	refreshTokenID := int32(refreshToken.ID)
	updateUserParam := repository.UpdateUserParams{
		RefreshToken: &refreshTokenID,
		Username:     user.Username,
		ID:           user.ID,
	}
	_, err = a.UserService.UpdateUser(updateUserParam)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}
	return responses.OK(ctx, GetKakaoInfoResponse{AccessToken: accessToken, RefreshToken: *refreshToken.Token})
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

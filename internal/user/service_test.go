package user

import (
	"context"
	"errors"
	"testing"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces/mocks"
	"github.com/stretchr/testify/require"
)

func TestUserService_FindUserByUserName(t *testing.T) {
	tests := []struct {
		before func(*testing.T, *mocks.IRepository)
		expect func(*testing.T, *UserService)
		name   string
	}{
		{
			name: "FindUserByUsername/find_user_failed",
			before: func(_ *testing.T, repo *mocks.IRepository) {
				str := ""
				repo.On("FindUserByUsername", context.Background(), &str).Return(repository.User{}, errors.New("")).Once()
			},
			expect: func(t *testing.T, service *UserService) {
				str := ""
				user, err := service.FindUserByUserName(&str)
				require.Error(t, err)
				require.Nil(t, user)
			},
		},
		{
			name: "FindUserByUsername/find_user_success",
			before: func(_ *testing.T, repo *mocks.IRepository) {
				str := ""
				repo.On("FindUserByUsername", context.Background(), &str).Return(repository.User{}, nil).Once()
			},
			expect: func(t *testing.T, service *UserService) {
				str := ""
				user, err := service.FindUserByUserName(&str)
				require.Nil(t, err)
				require.NotNil(t, user)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewIRepository(t)
			userService := InitUserService(repo)
			tt.before(t, repo)
			tt.expect(t, userService)
		})
	}
}

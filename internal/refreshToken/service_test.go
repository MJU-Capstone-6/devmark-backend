package refreshtoken

import (
	"context"
	"crypto/ed25519"
	"reflect"
	"testing"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	"github.com/MJU-Capstone-6/devmark-backend/internal/jwtToken"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces/mocks"
	"github.com/o1egl/paseto"
)

func TestRefreshTokenService_CreateToken(t *testing.T) {
	type fields struct {
		Repository interfaces.IRepository
		JwtService interfaces.IJWTService
	}
	type args struct {
		token string
	}
	pubKey, privateKey, _ := ed25519.GenerateKey(nil)
	jwtService := jwtToken.InitJWTService(pubKey, privateKey, "")
	token, err := jwtService.GenerateToken(1, constants.ACCESSTOKEN_EXPIRED_TIME)
	repo := mocks.NewIRepository(t)
	service := mocks.NewIJWTService(t)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *repository.RefreshToken
		wantErr bool
	}{
		{
			name: "CreateRefreshToken_Success",
			fields: fields{
				Repository: repo,
				JwtService: service,
			},
			args: args{
				token: token,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RefreshTokenService{
				Repository: tt.fields.Repository,
				JwtService: tt.fields.JwtService,
			}
			id := int32(0)
			refreshtoken := repository.RefreshToken{}
			pasetoToken := paseto.JSONToken{}
			pasetoToken.Set(constants.TOKEN_DATA_KEY, "0")
			repo.On("CreateRefreshToken", context.Background(), repository.CreateRefreshTokenParams{
				Token:  &tt.args.token,
				UserID: &id,
			}).Return(refreshtoken, nil)

			service.On("VerifyToken", tt.args.token).Return(pasetoToken, nil)

			_, err := r.CreateToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshTokenService.CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/*
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("RefreshTokenService.CreateToken() = %v, want %v", got, tt.want)
				}*/
		})
	}
}

func TestRefreshTokenService_UpdateToken(t *testing.T) {
	type fields struct {
		Repository interfaces.IRepository
		JwtService interfaces.IJWTService
	}
	type args struct {
		params repository.UpdateRefreshTokenParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *repository.RefreshToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RefreshTokenService{
				Repository: tt.fields.Repository,
				JwtService: tt.fields.JwtService,
			}
			got, err := r.UpdateToken(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshTokenService.UpdateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefreshTokenService.UpdateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRefreshTokenService_FindOneByUserId(t *testing.T) {
	type fields struct {
		Repository interfaces.IRepository
		JwtService interfaces.IJWTService
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *repository.RefreshToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RefreshTokenService{
				Repository: tt.fields.Repository,
				JwtService: tt.fields.JwtService,
			}
			got, err := r.FindOneByUserId(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshTokenService.FindOneByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefreshTokenService.FindOneByUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitRefreshTokenService(t *testing.T) {
	type args struct {
		repo    interfaces.IRepository
		service interfaces.IJWTService
	}
	tests := []struct {
		name string
		args args
		want RefreshTokenService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitRefreshTokenService(tt.args.repo, tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitRefreshTokenService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRefreshTokenService_WithRepository(t *testing.T) {
	type fields struct {
		Repository interfaces.IRepository
		JwtService interfaces.IJWTService
	}
	type args struct {
		repo interfaces.IRepository
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   RefreshTokenService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RefreshTokenService{
				Repository: tt.fields.Repository,
				JwtService: tt.fields.JwtService,
			}
			if got := r.WithRepository(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefreshTokenService.WithRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRefreshTokenService_WithJWTService(t *testing.T) {
	type fields struct {
		Repository interfaces.IRepository
		JwtService interfaces.IJWTService
	}
	type args struct {
		service interfaces.IJWTService
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   RefreshTokenService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RefreshTokenService{
				Repository: tt.fields.Repository,
				JwtService: tt.fields.JwtService,
			}
			if got := r.WithJWTService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefreshTokenService.WithJWTService() = %v, want %v", got, tt.want)
			}
		})
	}
}

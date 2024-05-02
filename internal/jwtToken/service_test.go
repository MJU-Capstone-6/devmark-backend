package jwtToken

import (
	"crypto/ed25519"
	"reflect"
	"testing"
	"time"

	"github.com/o1egl/paseto"
)

func TestJWTService_GenerateToken(t *testing.T) {
	type args struct {
		id      int
		expTime time.Time
	}

	publicKey, privateKey, nil := ed25519.GenerateKey(nil)
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		{
			name: "GenerateToken_success",
			args: args{
				id:      1,
				expTime: time.Now().Add(24 * time.Hour),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JWTService{
				PublicKey:  publicKey,
				PrivateKey: privateKey,
				Footer:     "test",
			}
			got, err := j.GenerateToken(tt.args.id, tt.args.expTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTService.GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var parsedJSONToken paseto.JSONToken
			var footer string
			if err := paseto.NewV2().Verify(*got, publicKey, &parsedJSONToken, &footer); err != nil {
				t.Errorf("Token is invalid error = %v", err)
			}
			/*
				if got != tt.want {
					t.Errorf("JWTService.GenerateToken() = %v, want %v", got, tt.want)
				}*/
		})
	}
}

func TestJWTService_DecryptToken(t *testing.T) {
	type args struct {
		token string
	}

	publicKey, privateKey, nil := ed25519.GenerateKey(nil)
	jsonToken := paseto.JSONToken{}
	generateToken, _ := paseto.NewV2().Sign(privateKey, &jsonToken, "test")
	tests := []struct {
		name    string
		args    args
		want    *paseto.JSONToken
		wantErr bool
	}{
		{
			name: "VerifyToken_failed",
			args: args{
				token: "",
			},
			wantErr: true,
		},
		{
			name: "VerifyToken_success",
			args: args{
				token: generateToken,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JWTService{
				PublicKey:  publicKey,
				PrivateKey: privateKey,
				Footer:     "test",
			}
			_, err := j.VerifyToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTService.VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/*
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("JWTService.VerifyToken() = %v, want %v", got, tt.want)
				}*/
		})
	}
}

func TestInitJWTService(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)

	tests := []struct {
		name string
		want *JWTService
	}{
		{
			name: "InitJWTService_success",
			want: &JWTService{
				PublicKey:  publicKey,
				PrivateKey: privateKey,
				Footer:     "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitJWTService(publicKey, privateKey, "test"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitJWTService() = %v, want %v", got, tt.want)
			}
		})
	}
}

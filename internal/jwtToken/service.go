package jwtToken

import (
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/o1egl/paseto"
)

type JWTService struct {
	PublicKey  ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
	Footer     string
}

func (j *JWTService) GenerateToken(id int, expTime time.Time) (string, error) {
	now := time.Now()

	jsonToken := paseto.JSONToken{
		IssuedAt:   now,
		Expiration: expTime,
	}
	jsonToken.Set(constants.TOKEN_DATA_KEY, fmt.Sprintf("%d", id))
	token, err := paseto.NewV2().Sign(j.PrivateKey, jsonToken, nil)
	if err != nil {
		return "", customerror.TokenSignFail(err)
	}
	return token, nil
}

func (j *JWTService) VerifyToken(token string) (paseto.JSONToken, error) {
	var parsedJSONToken paseto.JSONToken
	err := paseto.NewV2().Verify(token, j.PublicKey, &parsedJSONToken, nil)
	if err != nil {
		return parsedJSONToken, customerror.TokenVerifyFail(err)
	}
	return parsedJSONToken, nil
}

func InitJWTService(pubKey ed25519.PublicKey, privateKey ed25519.PrivateKey, footer string) *JWTService {
	return &JWTService{
		PublicKey:  pubKey,
		PrivateKey: privateKey,
		Footer:     footer,
	}
}

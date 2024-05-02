package utils

import (
	"crypto/ed25519"
	"encoding/hex"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KeyGeneratorFunc func(token []byte) ([]byte, error)

func GenerateKey(token string, keyGenerator KeyGeneratorFunc) ([]byte, error) {
	decodedString, err := hex.DecodeString(token)
	if err != nil {
		return nil, err
	}
	return keyGenerator(decodedString)
}

func GeneratePublicKey(token []byte) ([]byte, error) {
	return ed25519.PublicKey(token), nil
}

func GeneratePrivateKey(token []byte) ([]byte, error) {
	return ed25519.PrivateKey(token), nil
}

func Unauthorized(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusUnauthorized, data)
}

func OK(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, data)
}

func InternalServer(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusInternalServerError, data)
}

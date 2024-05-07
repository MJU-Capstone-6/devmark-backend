package utils

import (
	"crypto/ed25519"
	"encoding/hex"
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

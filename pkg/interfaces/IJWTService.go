package interfaces

import (
	"time"

	"github.com/o1egl/paseto"
)

//go:generate mockery --name IJWTService
type IJWTService interface {
	GenerateToken(int, time.Time) (string, error)
	DecryptToken(string) (*paseto.JSONToken, error)
	Verify(string) error
}

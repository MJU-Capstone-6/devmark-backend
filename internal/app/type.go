package app

import (
	"crypto/ed25519"

	"github.com/MJU-Capstone-6/devmark-backend/internal/config"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type Application struct {
	Repository repository.Queries
	DB         *pgx.Conn
	Handler    *echo.Echo
	Config     *config.Config
	PubKey     ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
}

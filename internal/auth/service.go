package auth

import "github.com/jackc/pgx/v5"

type AuthService struct {
	Conn *pgx.Conn
}

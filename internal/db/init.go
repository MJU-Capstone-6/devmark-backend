package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/MJU-Capstone-6/devmark-backend/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db   *pgxpool.Pool
	once sync.Once
)
var BASE_URL = "postgres://%s:%s@%s:%s/%s?sslmode=disable"

func InitDBbyConfig(ctx context.Context, c config.DB) (*pgxpool.Pool, error) {
	if db == nil {
		once.Do(func() {
			conn, err := setDBByConfig(ctx, c)
			if err != nil {
				log.Fatal(err)
			}
			db = conn
		})
	} else {
		return nil, errors.New("DB already configured")
	}
	return db, nil
}

func InitDBbyURL(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	if db == nil {
		once.Do(func() {
			conn, err := setDBByURL(ctx, dbURL)
			if err != nil {
				log.Fatal(err)
			}
			db = conn
		})
	} else {
		return nil, errors.New("DB already configured")
	}
	return db, nil
}

func setDBByConfig(ctx context.Context, c config.DB) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf(BASE_URL, c.Username, c.Password, c.Host, c.Port, c.Name)
	conn, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func setDBByURL(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

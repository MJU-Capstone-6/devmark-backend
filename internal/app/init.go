package app

import (
	"context"
	"crypto/ed25519"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/MJU-Capstone-6/devmark-backend/internal/config"
	"github.com/MJU-Capstone-6/devmark-backend/internal/db"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	app  *Application
	once sync.Once
)

func InitApplication() (*Application, error) {
	if app == nil {
		once.Do(func() {
			err := setApplication()
			if err != nil {
				log.Fatal("something went wrong while configure application")
			}
		})
	} else {
		return nil, errors.New("Application already configured")
	}
	return app, nil
}

func (app *Application) Run() error {
	if err := app.Handler.Start(fmt.Sprintf(":%s", app.Config.App.Port)); err != nil {
		return err
	}
	defer app.DB.Close(context.Background())
	return nil
}

func runWithPostgresContainer(ctx context.Context, dbConfig config.DB) (*postgres.PostgresContainer, error) {
	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:14-alpine"),
		postgres.WithDatabase(dbConfig.Name),
		postgres.WithUsername(dbConfig.Username),
		postgres.WithPassword(dbConfig.Password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)))
	if err != nil {
		return nil, err
	}
	return postgresContainer, nil
}

func setApplication() error {
	ctx := context.Background()
	applicationConfig, err := config.InitConfig()
	if err != nil {
		return err
	}

	postgresContainer, err := runWithPostgresContainer(ctx, applicationConfig.DB)
	if err != nil {
		return err
	}

	db, err := db.InitDB(ctx, applicationConfig.DB)
	if err != nil {
		return err
	}

	handler := echo.New()
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return err
	}
	app = &Application{
		DB:         db,
		Config:     applicationConfig,
		Repository: *repository.New(db),
		PubKey:     publicKey,
		PrivateKey: privateKey,
		Handler:    handler,
	}

	app.InitRoutes()
	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

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
	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	"github.com/MJU-Capstone-6/devmark-backend/internal/db"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/google/generative-ai-go/genai"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"google.golang.org/api/option"
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
				log.Println(err)
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
	defer app.DB.Close()
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
	/*
		postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image:        "postgres:14-alpine",
				ExposedPorts: []string{"5432/tcp"},
				WaitingFor: wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).
					WithStartupTimeout(5 * time.Second),
				Env: map[string]string{
					"POSTGRES_USER":     dbConfig.Username,
					"POSTGRES_DB":       dbConfig.Name,
					"POSTGRES_PASSWORD": dbConfig.Password,
				},
			},
			Started: true,
		})*/
	return postgresContainer, nil
}

func setApplication() error {
	var dbConn *pgxpool.Pool

	ctx := context.Background()
	applicationConfig, err := config.InitConfig()
	if err != nil {
		return err
	}
	if applicationConfig.App.IsDevMode {
		postgresContainer, err := runWithPostgresContainer(ctx, applicationConfig.DB)
		if err != nil {
			return err
		}
		dbURL, _ := postgresContainer.ConnectionString(ctx)
		dbConn, err = db.InitDBbyURL(ctx, dbURL)
		if err != nil {
			return err
		}
		err = db.Migration(fmt.Sprintf("%ssslmode=disable", dbURL))
		if err != nil {
			log.Println(err)
		}

	} else {
		dbConn, err = db.InitDBbyConfig(ctx, applicationConfig.DB)
		if err != nil {
			return err
		}
	}

	handler := echo.New()
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return err
	}
	gptClient := openai.NewClient(applicationConfig.OpenAI.ClientKey)
	geminiClient, err := genai.NewClient(context.Background(), option.WithAPIKey(applicationConfig.GoogleAPI.ClientKey))
	if err != nil {
		return err
	}
	geminiModel := geminiClient.GenerativeModel(constants.GEMINI_FLASH_MODEL)
	app = &Application{
		DB:           dbConn,
		Config:       applicationConfig,
		Repository:   *repository.New(dbConn),
		PubKey:       publicKey,
		PrivateKey:   privateKey,
		Handler:      handler,
		GPTClient:    gptClient,
		GeminiClient: geminiClient,
		GeminiModel:  geminiModel,
	}

	app.InitRoutes()

	return nil
}

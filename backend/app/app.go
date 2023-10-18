package app

import (
	"context"
	"fmt"
	"github.com/artemKapitonov/school-library/backend/internal/config"
	"github.com/artemKapitonov/school-library/backend/internal/usecase/storage"
	"github.com/artemKapitonov/school-library/backend/pkg/client/postgresql"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"

	"github.com/artemKapitonov/school-library/backend/internal/endpoint"
	"github.com/artemKapitonov/school-library/backend/internal/usecase"
)

type App struct {
	Storage  *storage.Storage
	Endpoint *endpoint.Endpoints
	UseCase  *usecase.UseCase
}

func New() *App {
	var app = &App{}
	ctx := context.TODO()

	if err := config.Init(); err != nil {
		panic(err)
	}

	err := godotenv.Load(".env")
	if err != nil {
		panic("env")
	}

	var cfg = postgresql.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.SSLMode"),
	}

	fmt.Println(cfg)

	db, err := postgresql.ConnectToDB(ctx, cfg)
	if err != nil {
		panic("DB connect")
	}

	fmt.Println("Connection success")

	app.Storage = storage.New(db)

	app.UseCase = usecase.New(app.Storage)

	app.Endpoint = endpoint.New(&ctx, app.UseCase)

	return app
}

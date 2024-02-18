package app

import (
	"context"
	"log/slog"

	"github.com/artemKapitonov/school-library/backend/internal/storage"
	"github.com/artemKapitonov/school-library/backend/internal/transport"
	"github.com/artemKapitonov/school-library/backend/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	ctx       context.Context
	Transport *transport.Transport
	Storage   *storage.Storage
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func New() *App {
	var app = &App{}

	var cfg logging.Config
	cleanenv.ReadConfig("backend/config/logger.yaml", &cfg)

	logger := logging.New(cfg)
	slog.SetDefault(logger.Logger)

	db, err := storage.ConnectToDB()
	if err != nil {
		slog.Error("Failed connect to db", err)
		panic(err)
	}
	slog.Info("Succesfuly connect ro database")

	app.Storage = storage.New(db)

	app.Transport = transport.New(app.Storage)

	return app
}

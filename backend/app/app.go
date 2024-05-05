package app

import (
	"context"
	"log/slog"

	"github.com/artemKapitonov/school-library/backend/internal/storage"
	"github.com/artemKapitonov/school-library/backend/internal/transport"
	"github.com/artemKapitonov/school-library/backend/migrations"
	"github.com/artemKapitonov/school-library/backend/pkg/logging"
)

type App struct {
	ctx       context.Context
	Transport *transport.Transport
	Storage   *storage.Storage
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OnShutdown(ctx context.Context) {
	err := a.Storage.DB.Close()
	if err != nil {
		slog.Warn("Storage don't close Error:", err)
	} else {
		slog.Info("Storage succusully closed")
	}
}

func New() *App {
	var app = &App{}

	var cfg = logging.Config{}

	logger := logging.New(cfg)
	slog.SetDefault(logger.Logger)

	db, err := storage.Connect()
	if err != nil {
		slog.Error("Failed connect to db", err)
		panic(err)
	}
	slog.Info("Succesfuly connect to database")

	if err := migrations.Create(db); err != nil {
		panic(err)
	}

	app.Storage = storage.New(db)

	app.Transport = transport.New(app.Storage)

	return app
}

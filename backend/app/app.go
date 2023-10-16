package app

import (
	"context"
	"github.com/artemKapitonov/school-library/backend/internal/endpoint"
	"github.com/artemKapitonov/school-library/backend/internal/usecase"
)

type App struct {
	Endpoint *endpoint.Endpoints
	UseCase  *usecase.UseCase
}

func New() *App {
	var app = &App{}

	app.Endpoint = endpoint.New(context.Background())

	app.UseCase = usecase.New()

	return app
}

func (*App) Hello() {
	return
}

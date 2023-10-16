package app

import (
	"github.com/artemKapitonov/school-library/backend/internal/endpoint"
	"github.com/artemKapitonov/school-library/backend/internal/usecase"
)

type App struct {
	Endpoint *endpoint.Endpoints
	UseCase  *usecase.UseCase
}

func New() *App {
	var app = &App{}

	app.UseCase = usecase.New()

	app.Endpoint = endpoint.New()

	return app
}

func (*App) Hello() {
	return
}

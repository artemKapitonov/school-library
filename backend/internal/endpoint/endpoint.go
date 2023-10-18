package endpoint

import (
	"context"
	"fmt"
	"github.com/artemKapitonov/school-library/backend/internal/usecase"
)

type Endpoints struct {
	ctx *context.Context
	Class
	Books
	Students
}

func New(ctx *context.Context, useCase *usecase.UseCase) *Endpoints {
	return &Endpoints{ctx: ctx}
}

// Startup is called when the endpoint starts. The context is saved
// So we can call the runtime methods
func (a *Endpoints) Startup(ctx context.Context) {
	c := &ctx
	a.ctx = c
}

// Greet returns a greeting for the given name
func (a *Endpoints) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

package endpoint

import (
	"context"
	"fmt"
)

type Endpoints struct {
	ctx context.Context
}

func New(ctx context.Context) *Endpoints {
	return &Endpoints{ctx: ctx}
}

// Startup is called when the endpoint starts. The context is saved
// So we can call the runtime methods
func (a *Endpoints) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *Endpoints) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

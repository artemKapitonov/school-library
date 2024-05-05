package main

import (
	"embed"
	"log/slog"

	"github.com/artemKapitonov/school-library/backend/app"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := app.New()

	// Create application with options
	err := wails.Run(&options.App{
		Width: 1500,
		Title: "school-library",
		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		OnStartup:  app.Startup,
		OnShutdown: app.OnShutdown,
		Bind: []interface{}{
			app.Transport,
		},
	})
	if err != nil {
		slog.Error("Failed to start application Error:", err)
	}

}

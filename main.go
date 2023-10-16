package main

import (
	"embed"
	"github.com/artemKapitonov/school-library/backend/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	a := app.New()
	// Create an instance of the  structure
	point := a.Endpoint

	// Create  with options
	err := wails.Run(&options.App{
		Title:  "school-library",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        point.Startup,
		Bind: []interface{}{
			point,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

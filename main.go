package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
//Embed empacota arquivos estaticos em compilados. 
var assets embed.FS

func main() {
	// Create an instance of the app structure
	NewArrow := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "NewArrow",
		Width:  728,
		Height: 414,
		AssetServer: &assetserver.Options{
			Assets: assets,
			Handler: NewMusicAssetHandler(),
		},
		Bind: []interface{}{
			NewArrow,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
 
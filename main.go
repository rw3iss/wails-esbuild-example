package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

// Todo: read in the static filenames from client-solid/build/asset-manifest.json

func main() {

	js := mewn.String("./client/build/main.js")
	css := mewn.String("./client/build/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "example",
		JS:     js,
		CSS:    css,
		Colour: "#2255aa",
	})
	app.Bind(&Counter{})
	app.Run()
}

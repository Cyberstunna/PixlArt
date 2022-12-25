package ui

func Setup(app *AppInit) {
	swatchesContainer := BuildWatches(app)

	app.PixlWindow.SetContent(swatchesContainer)
}

package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowTableWindow(parent fyne.Window, dbName string) {
	window := fyne.CurrentApp().NewWindow("Tables - " + dbName)

	// TODO: Implement table listing and management
	content := container.NewVBox(
		widget.NewLabel("Tables in "+dbName),
		widget.NewLabel("Coming soon..."),
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(600, 400))
	window.Show()
}

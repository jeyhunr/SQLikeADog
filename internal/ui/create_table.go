package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowCreateTableWindow(parent fyne.Window, dbName string) {
	window := fyne.CurrentApp().NewWindow("Create Table - " + dbName)

	// TODO: Implement table creation form
	content := container.NewVBox(
		widget.NewLabel("Create Table in "+dbName),
		widget.NewLabel("Coming soon..."),
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(600, 400))
	window.Show()
}

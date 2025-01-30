package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowEditTableWindow(parent fyne.Window, dbName, tableName string) {
	window := fyne.CurrentApp().NewWindow("Edit Table - " + tableName)

	// TODO: Implement table editing interface
	content := container.NewVBox(
		widget.NewLabel("Edit Table "+tableName+" in "+dbName),
		widget.NewLabel("Coming soon..."),
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(600, 400))
	window.Show()
}

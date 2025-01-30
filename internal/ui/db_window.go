package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jeyhunr/SQLikeADog/internal/db"
)

func ShowDBWindow(parent fyne.Window) {
	window := fyne.CurrentApp().NewWindow("Databases")

	databases, err := db.ListDatabases()
	if err != nil {
		log.Printf("Error listing databases: %v", err)
		return
	}

	var items []string
	for _, database := range databases {
		items = append(items, database)
	}

	list := widget.NewList(
		func() int {
			return len(items)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(items[i])
		})

	content := container.NewVBox(
		widget.NewLabel("Available Databases"),
		list,
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(400, 300))
	window.Show()
}

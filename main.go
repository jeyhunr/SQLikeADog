package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("MySQL Client")

	// Create UI components
	queryEntry := widget.NewMultiLineEntry()
	queryEntry.SetPlaceHolder("Enter SQL query here...")

	resultLabel := widget.NewLabel("Results will be displayed here")

	executeButton := widget.NewButton("Execute", func() {
		query := queryEntry.Text
		results, err := executeQuery(query)
		if err != nil {
			resultLabel.SetText(fmt.Sprintf("Error: %v", err))
		} else {
			resultLabel.SetText(results)
		}
	})

	// Layout
	content := container.NewVBox(
		queryEntry,
		executeButton,
		resultLabel,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.ShowAndRun()
}

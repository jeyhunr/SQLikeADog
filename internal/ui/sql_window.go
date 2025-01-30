package ui

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jeyhunr/SQLikeADog/internal/db"
)

func ShowSQLWindow(parent fyne.Window) {
	window := fyne.CurrentApp().NewWindow("Execute SQL")

	// Create SQL input
	sqlInput := widget.NewMultiLineEntry()
	sqlInput.SetPlaceHolder("Enter your SQL query here...")

	// Create results area
	resultsArea := widget.NewTextGrid()

	// Execute button
	executeBtn := widget.NewButton("Execute", func() {
		if sqlInput.Text == "" {
			ShowErrorPopUp("Please enter a SQL query", window.Canvas())
			return
		}

		// Execute query and show results
		columns, rows, err := db.ExecuteQuery(sqlInput.Text)
		if err != nil {
			ShowErrorPopUp("Query execution failed: "+err.Error(), window.Canvas())
			return
		}

		// Format and display results
		var result string
		result = "Columns: " + strings.Join(columns, ", ") + "\n\n"
		for _, row := range rows {
			result += strings.Join(row, " | ") + "\n"
		}
		resultsArea.SetText(result)
	})

	// Layout
	content := container.NewBorder(
		container.NewVBox(
			sqlInput,
			executeBtn,
		),
		nil, nil, nil,
		container.NewVScroll(resultsArea),
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(800, 600))
	window.Show()
}

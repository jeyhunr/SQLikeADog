package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jeyhunr/SQLikeADog/internal/auth"
	"github.com/jeyhunr/SQLikeADog/internal/db"
)

type MainWindow struct {
	win fyne.Window
}

func NewMainWindow() *MainWindow {
	myWindow := fyne.CurrentApp().NewWindow("SQLikeADog - Main")
	return &MainWindow{
		win: myWindow,
	}
}

func (mw *MainWindow) Show() {
	mw.win.SetContent(mw.makeUI())
	mw.win.Resize(fyne.NewSize(800, 600))
	mw.win.Show()
}

func (mw *MainWindow) makeUI() fyne.CanvasObject {
	// Create left sidebar for databases
	databases, err := db.ListDatabases()
	if err != nil {
		log.Printf("Error listing databases: %v", err)
		return container.NewVBox(widget.NewLabel("Error loading databases"))
	}

	// Create database list
	dbList := widget.NewList(
		func() int { return len(databases) },
		func() fyne.CanvasObject { return widget.NewLabel("Template") },
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(databases[lii])
		},
	)

	rightHeader := widget.NewLabel("Select a database to view tables")
	rightContent := container.NewMax()

	var showTablesList func(string) // Declare the function
	showTablesList = func(selectedDB string) {
		tables, err := db.ListTables(selectedDB)
		if err != nil {
			ShowErrorPopUp("Error listing tables: "+err.Error(), mw.win.Canvas())
			return
		}

		tableList := widget.NewList(
			func() int { return len(tables) },
			func() fyne.CanvasObject { return widget.NewLabel("Template") },
			func(lii widget.ListItemID, co fyne.CanvasObject) {
				co.(*widget.Label).SetText(tables[lii])
			},
		)

		// Handle table selection
		tableList.OnSelected = func(id widget.ListItemID) {
			selectedTable := tables[id]

			// Create back button
			backButton := widget.NewButton("Back to Tables", func() {
				rightHeader.SetText("Tables in " + selectedDB)
				showTablesList(selectedDB)
			})

			// Get table data
			columns, data, err := db.GetTableData(selectedDB, selectedTable)
			if err != nil {
				ShowErrorPopUp("Error fetching table data: "+err.Error(), mw.win.Canvas())
				return
			}

			// Create table widget
			table := widget.NewTable(
				func() (int, int) {
					return len(data) + 1, len(columns)
				},
				func() fyne.CanvasObject {
					label := widget.NewLabel("Template")
					label.Alignment = fyne.TextAlignLeading
					label.Wrapping = fyne.TextTruncate
					return container.NewPadded(label)
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					label := o.(*fyne.Container).Objects[0].(*widget.Label)

					if i.Row == 0 {
						// Header row
						label.TextStyle = fyne.TextStyle{Bold: true}
						label.SetText(columns[i.Col])
					} else {
						// Data rows
						value := data[i.Row-1][i.Col]
						// Clean up the value string
						if value == "NULL" {
							label.SetText("NULL")
						} else if len(value) > 2 && value[0] == '[' && value[len(value)-1] == ']' {
							// Remove array brackets
							label.SetText(value[1 : len(value)-1])
						} else {
							if len(value) > 50 {
								value = value[:47] + "..."
							}
							label.SetText(value)
						}
						label.TextStyle = fyne.TextStyle{}
					}
				})

			// Set column widths based on content
			for i := 0; i < len(columns); i++ {
				maxWidth := len(columns[i]) * 10 // Base width on header length
				for _, row := range data {
					if len(row[i])*8 > maxWidth {
						maxWidth = len(row[i]) * 8
					}
				}
				table.SetColumnWidth(i, float32(maxWidth))
			}

			// Update header and content
			rightHeader.SetText("Table: " + selectedTable)
			rightContent.Objects = []fyne.CanvasObject{
				container.NewBorder(
					backButton,
					nil,
					nil,
					nil,
					container.NewVScroll(table),
				),
			}
			rightContent.Refresh()
		}

		rightHeader.SetText("Tables in " + selectedDB)
		rightContent.Objects = []fyne.CanvasObject{container.NewVScroll(tableList)}
		rightContent.Refresh()
	}

	// Now we can use showTablesList in the OnSelected handler
	dbList.OnSelected = func(id widget.ListItemID) {
		selectedDB := databases[id]
		showTablesList(selectedDB)
	}

	// Logout button
	logoutButton := widget.NewButton("Logout", func() {
		if err := auth.DeleteCredentials(); err != nil {
			ShowErrorPopUp("Failed to logout: "+err.Error(), mw.win.Canvas())
			return
		}

		loginWindow := NewLoginWindow()
		loginWindow.Show()
		mw.win.Close()
	})

	// Create left sidebar container
	leftContent := container.NewBorder(
		widget.NewLabel("Databases"),
		logoutButton,
		nil,
		nil,
		container.NewVScroll(dbList),
	)

	// Create right container
	rightContainer := container.NewBorder(
		rightHeader,
		nil,
		nil,
		nil,
		rightContent,
	)

	// Create split container
	split := container.NewHSplit(
		leftContent,
		rightContainer,
	)
	split.SetOffset(0.3)

	return split
}

func (mw *MainWindow) GetWindow() fyne.Window {
	return mw.win
}

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

	// Create right content area (initially empty)
	rightHeader := widget.NewLabel("Select a database to view tables")
	rightContent := container.NewMax()

	// Handle database selection
	dbList.OnSelected = func(id widget.ListItemID) {
		selectedDB := databases[id]
		tables, err := db.ListTables(selectedDB)
		if err != nil {
			ShowErrorPopUp("Error listing tables: "+err.Error(), mw.win.Canvas())
			return
		}

		// Create table list
		tableList := widget.NewList(
			func() int { return len(tables) },
			func() fyne.CanvasObject { return widget.NewLabel("Template") },
			func(lii widget.ListItemID, co fyne.CanvasObject) {
				co.(*widget.Label).SetText(tables[lii])
			},
		)

		// Update right header and content
		rightHeader.SetText("Tables in " + selectedDB)
		rightContent.Objects = []fyne.CanvasObject{container.NewVScroll(tableList)}
		rightContent.Refresh()
	}

	// Logout button
	logoutButton := widget.NewButton("Logout", func() {
		if err := auth.DeleteCredentials(); err != nil {
			ShowErrorPopUp("Failed to logout: "+err.Error(), mw.win.Canvas())
			return
		}

		// Show login window
		loginWindow := NewLoginWindow()
		loginWindow.Show()
		mw.win.Close()
	})

	// Create left sidebar container with header and scrollable list
	leftContent := container.NewBorder(
		widget.NewLabel("Databases"), // Top
		logoutButton,                 // Bottom
		nil,                          // Left
		nil,                          // Right
		container.NewVScroll(dbList), // Center content (scrollable)
	)

	// Create right container with header and scrollable content
	rightContainer := container.NewBorder(
		rightHeader, // Top
		nil,         // Bottom
		nil,         // Left
		nil,         // Right
		rightContent,
	)

	// Create split container
	split := container.NewHSplit(
		leftContent,
		rightContainer,
	)
	split.SetOffset(0.3) // 30% width for the left sidebar

	return split
}

func (mw *MainWindow) GetWindow() fyne.Window {
	return mw.win
}

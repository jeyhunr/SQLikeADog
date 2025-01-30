package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jeyhunr/SQLikeADog/internal/auth"
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
	// Logout button
	logoutButton := widget.NewButton("Logout", func() {
		if err := auth.DeleteCredentials(); err != nil {
			widget.NewPopUp(widget.NewLabel("Failed to logout: "+err.Error()), mw.win.Canvas())
			return
		}

		// Redirect to the login window
		loginWindow := NewLoginWindow()
		loginWindow.Show()
		mw.win.Close()
	})

	return container.NewVBox(
		logoutButton,
	)
}

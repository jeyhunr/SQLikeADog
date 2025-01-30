package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jeyhunr/SQLikeADog/internal/auth"
	"github.com/jeyhunr/SQLikeADog/internal/db"
)

type LoginWindow struct {
	win fyne.Window
}

func NewLoginWindow() *LoginWindow {
	myWindow := fyne.CurrentApp().NewWindow("SQLikeADog - Login")
	return &LoginWindow{
		win: myWindow,
	}
}

func (lw *LoginWindow) Show() {
	lw.win.SetContent(lw.makeUI())
	lw.win.Resize(fyne.NewSize(400, 300))
	lw.win.Show()
}

func (lw *LoginWindow) makeUI() fyne.CanvasObject {
	dbHost := widget.NewEntry()
	dbHost.SetPlaceHolder("Enter DB Host")

	dbUser := widget.NewEntry()
	dbUser.SetPlaceHolder("Enter DB User")

	dbPassword := widget.NewEntry()
	dbPassword.SetPlaceHolder("Enter DB Password")
	dbPassword.Password = true

	dbName := widget.NewEntry()
	dbName.SetPlaceHolder("Enter DB Name")

	loginButton := widget.NewButton("Login", func() {
		creds := auth.Credentials{
			Host:     dbHost.Text,
			User:     dbUser.Text,
			Password: dbPassword.Text,
			DBName:   dbName.Text,
		}

		// Save credentials to JSON
		if err := auth.SaveCredentials(creds); err != nil {
			widget.NewPopUp(widget.NewLabel("Failed to save credentials: "+err.Error()), lw.win.Canvas())
			return
		}

		// Connect to the database
		if err := db.Connect(creds.Host, creds.User, creds.Password, creds.DBName); err != nil {
			widget.NewPopUp(widget.NewLabel("Failed to connect: "+err.Error()), lw.win.Canvas())
			return
		}

		// Redirect to the main window
		mainWindow := NewMainWindow()
		mainWindow.Show()
		lw.win.Close()
	})

	return container.NewVBox(
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		loginButton,
	)
}

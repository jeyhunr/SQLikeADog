package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

	dbPort := widget.NewEntry()
	dbPort.SetPlaceHolder("Enter Port (default: 3306)")
	dbPort.SetText("3306")

	dbUser := widget.NewEntry()
	dbUser.SetPlaceHolder("Enter DB User")

	dbPassword := widget.NewEntry()
	dbPassword.SetPlaceHolder("Enter DB Password")
	dbPassword.Password = true

	dbName := widget.NewEntry()
	dbName.SetPlaceHolder("Enter DB Name (optional)")

	loginButton := widget.NewButton("Login", func() {
		// Validate required input fields
		if dbHost.Text == "" || dbUser.Text == "" || dbPassword.Text == "" {
			ShowErrorPopUp("Required fields: Host, User, and Password", lw.win.Canvas())
			return
		}

		creds := auth.Credentials{
			Host:     dbHost.Text,
			Port:     dbPort.Text,
			User:     dbUser.Text,
			Password: dbPassword.Text,
			DBName:   dbName.Text, // Optional
		}

		// Save credentials to JSON
		if err := auth.SaveCredentials(creds); err != nil {
			ShowErrorPopUp("Failed to save credentials: "+err.Error(), lw.win.Canvas())
			return
		}

		// Connect to the database
		if err := db.Connect(creds.Host, creds.Port, creds.User, creds.Password, creds.DBName); err != nil {
			// Remove saved credentials if connection fails
			_ = auth.DeleteCredentials()
			ShowErrorPopUp("Database connection failed:\n"+err.Error(), lw.win.Canvas())
			return
		}

		// Redirect to the main window
		mainWindow := NewMainWindow()
		mainWindow.Show()
		lw.win.Close()
	})

	return container.NewVBox(
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
		loginButton,
	)
}

// Add this helper function for error popups
func ShowErrorPopUp(message string, canvas fyne.Canvas) {
	// Create close button
	closeButton := widget.NewButton("Close", nil)

	errorContainer := container.NewVBox(
		widget.NewLabel(message),
		container.NewHBox(
			layout.NewSpacer(),
			closeButton,
		),
	)
	errorContainer.Resize(fyne.NewSize(300, -1))

	popup := widget.NewModalPopUp(errorContainer, canvas)
	popup.Resize(fyne.NewSize(300, 100))

	// Set close button action after popup is created
	closeButton.OnTapped = popup.Hide

	// Position at the bottom
	popup.Move(fyne.NewPos(
		(canvas.Size().Width-popup.Size().Width)/2,
		canvas.Size().Height-popup.Size().Height-20,
	))

	popup.Show()
}

func (lw *LoginWindow) GetWindow() fyne.Window {
	return lw.win
}

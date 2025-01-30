package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/jeyhunr/SQLikeADog/internal/auth"
	"github.com/jeyhunr/SQLikeADog/internal/db"
	"github.com/jeyhunr/SQLikeADog/internal/ui"
)

func main() {
	myApp := app.New()

	// Check if credentials exist
	creds, err := auth.LoadCredentials()
	if err != nil {
		// No credentials found, show login window
		loginWindow := ui.NewLoginWindow()
		loginWindow.Show()
	} else {
		// Credentials found, connect to the database and show the main window
		if err := db.Connect(creds.Host, creds.Port, creds.User, creds.Password, creds.DBName); err != nil {
			// If connection fails, delete credentials and show login window
			_ = auth.DeleteCredentials()
			loginWindow := ui.NewLoginWindow()
			ui.ShowErrorPopUp("Failed to connect to database:\n"+err.Error(), loginWindow.GetWindow().Canvas())
			loginWindow.Show()
		} else {
			mainWindow := ui.NewMainWindow()
			mainWindow.Show()
		}
	}

	myApp.Run()
}

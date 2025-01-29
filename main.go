package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("SQLikeADog")

	// Check if configuration file exists
	if _, err := os.Stat("dbconfig.json"); err == nil {
		// Configuration exists, show the database page
		showDatabasePage(myWindow)
	} else {
		// No configuration, show the connection setup page
		showConnectionSetupPage(myWindow)
	}

	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}

func showConnectionSetupPage(myWindow fyne.Window) {
	// Create input fields for database connection details
	userEntry := widget.NewEntry()
	userEntry.SetPlaceHolder("Enter username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password")

	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("Enter host")

	dbNameEntry := widget.NewEntry()
	dbNameEntry.SetPlaceHolder("Enter database name")

	// Create a button to save the configuration
	saveButton := widget.NewButton("Save Configuration", func() {
		config := DBConfig{
			User:     userEntry.Text,
			Password: passwordEntry.Text,
			Host:     hostEntry.Text,
			DBName:   dbNameEntry.Text,
		}

		err := saveConfig(config, "dbconfig.json")
		if err != nil {
			log.Printf("Error saving config: %v", err)
		} else {
			log.Println("Configuration saved successfully")
			showDatabasePage(myWindow)
		}
	})

	// Layout
	content := container.NewVBox(
		userEntry,
		passwordEntry,
		hostEntry,
		dbNameEntry,
		saveButton,
	)

	myWindow.SetContent(content)
}

func showDatabasePage(myWindow fyne.Window) {
	// Create a logout button
	logoutButton := widget.NewButton("Logout", func() {
		err := os.Remove("dbconfig.json")
		if err != nil {
			log.Printf("Error removing config: %v", err)
		} else {
			log.Println("Logged out successfully")
			showConnectionSetupPage(myWindow)
		}
	})

	// Layout
	content := container.NewVBox(
		widget.NewLabel("Welcome to the Database Page"),
		logoutButton,
	)

	myWindow.SetContent(content)
}

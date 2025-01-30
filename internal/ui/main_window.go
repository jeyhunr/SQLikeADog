package ui

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jeyhunr/SQLikeADog/internal/utils"
)

func ShowMainWindow(app fyne.App) {
	myWindow := app.NewWindow("SQLikeADog")

	if _, err := os.Stat("dbconfig.json"); err == nil {
		showDatabasePage(myWindow)
	} else {
		showConnectionSetupPage(myWindow)
	}

	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}

func showConnectionSetupPage(myWindow fyne.Window) {
	userEntry := widget.NewEntry()
	userEntry.SetPlaceHolder("Enter username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password")

	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("Enter host")

	dbNameEntry := widget.NewEntry()
	dbNameEntry.SetPlaceHolder("Enter database name")

	saveButton := widget.NewButton("Save Configuration", func() {
		config := utils.DBConfig{
			User:     userEntry.Text,
			Password: passwordEntry.Text,
			Host:     hostEntry.Text,
			DBName:   dbNameEntry.Text,
		}

		err := utils.SaveConfig(config, "dbconfig.json")
		if err != nil {
			log.Printf("Error saving config: %v", err)
		} else {
			log.Println("Configuration saved successfully")
			showDatabasePage(myWindow)
		}
	})

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
	logoutButton := widget.NewButton("Logout", func() {
		err := os.Remove("dbconfig.json")
		if err != nil {
			log.Printf("Error removing config: %v", err)
		} else {
			log.Println("Logged out successfully")
			showConnectionSetupPage(myWindow)
		}
	})

	content := container.NewVBox(
		widget.NewLabel("Welcome to the Database Page"),
		logoutButton,
	)

	myWindow.SetContent(content)
}

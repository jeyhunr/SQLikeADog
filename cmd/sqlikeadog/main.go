package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/jeyhunr/SQLikeADog/internal/ui"
)

func main() {
	myApp := app.New()
	ui.ShowMainWindow(myApp)
}

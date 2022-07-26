package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Приложение")
	w.Resize(fyne.NewSize(400, 320))

	w.SetContent(container.NewVBox())

	w.ShowAndRun()

}

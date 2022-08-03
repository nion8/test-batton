package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Приложение")
	w.Resize(fyne.NewSize(400, 320))

	checks := widget.NewCheckGroup([]string{"Check1", "Check2", "Check3"}, func() {

	})

	w.SetContent(container.NewVBox())

	w.ShowAndRun()

}

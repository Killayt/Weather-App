package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MakeGUI() {
	a := app.New()
	w := a.NewWindow("Weather K")

	w.Resize(fyne.Size{Width: 500, Height: 700})

	// FIND FORM
	input := widget.NewEntry()
	input.SetPlaceHolder("City")

	content := container.NewVBox(input, widget.NewButton("Find", func() {
		log.Println("Content was", input.Text)
	}))

	w.SetContent(content)

	// MAIN FORM

	// w.SetContent(widget.NewButton("Find"))

	w.ShowAndRun()
}

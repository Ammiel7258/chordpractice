package ui

import (
	"chordpractice/handlers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoadHomeWindow(window fyne.Window) fyne.CanvasObject {
	label := widget.NewLabel("Label:")
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Place holder")

	button := widget.NewButton("Greet me", func() {
		handlers.HandleGreet(entry, label)
	})

	return container.NewVBox(label, entry, button)
}

func GetHomeContent(window fyne.Window) fyne.CanvasObject {

  bpmSelector := drawBPMSelector(3)

  return container.NewVBox(bpmSelector)
}

func drawBPMSelector(bpm int) fyne.CanvasObject {
  label := widget.NewLabel("BPM:")
  entry := widget.NewEntry()
  entry.SetPlaceHolder("What BPM would you like to practice with?")

  button := widget.NewButton("Continue", func() {

  })

  return container.NewHBox(label, entry, button)
}

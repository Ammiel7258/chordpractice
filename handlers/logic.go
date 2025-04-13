package handlers

import "fyne.io/fyne/v2/widget"

func HandleGreet(entry *widget.Entry, label *widget.Label) {
  name := entry.Text

  if name == "" {
    label.SetText("Please enter your name.")
  } else {
    label.SetText("Hello, " + name + "!")
  }
}

// When this is called, handle the submit button for the bpm
func HandleSubmit(entry *widget.Entry, label *widget.Label) {

}

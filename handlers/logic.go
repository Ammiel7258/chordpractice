package handlers

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/widget"
)

func HandleGreet(entry *widget.Entry, label *widget.Label) {
	name := entry.Text

	if name == "" {
		label.SetText("Please enter your name.")
	} else {
		label.SetText("Hello, " + name + "!")
	}
}

// When this is called, handle the submit button for the bpm
func HandleSubmit(entry *widget.Entry, label *widget.Label) int {
	text := entry.Text

	bpm, err := strconv.Atoi(text)
	if err != nil {
    label.SetText("Invalid BPM entered, setting to 60!")
		return 60
	}

  label.SetText(fmt.Sprintf("BPM set to %s", bpm))
	return bpm
}

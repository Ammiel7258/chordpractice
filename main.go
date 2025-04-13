package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chordpractice/chords"
	"chordpractice/ui"

	"fyne.io/fyne/v2/app"
)

var unusedChords []string
var usedChords []string
var bpm int

func main() {

	chordPracticeApp := app.New()

	backgroundWindow := chordPracticeApp.NewWindow("Guitar Chord Practice App")

	content := ui.LoadHomeWindow(backgroundWindow)

	backgroundWindow.SetContent(content)

  backgroundWindow.ShowAndRun()
}

func menu() {
	fmt.Println("What bpm will you be strumming at?")
	fmt.Scan(&bpm)

	fmt.Printf("Displaying chords to play at %d bmp! Press CTRL + C to stop!\n", bpm)

	unusedChords = chords.UnusedChords
	usedChords = chords.UsedChords

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	ticker := time.NewTicker(time.Minute / time.Duration(bpm))

	for {
		<-ticker.C

		fmt.Printf("%s\n", getRandomChord())
	}
}

func getRandomChord() string {
	if len(unusedChords) < 1 {
		unusedChords = append(unusedChords, usedChords[0:]...)
		usedChords = usedChords[:0]
	}

	index := rand.Intn(len(unusedChords))
	chord := unusedChords[index]

	usedChords = append(usedChords, chord)
	unusedChords = append(unusedChords[:index], unusedChords[index+1:]...)

	return chord
}

/*
	"bytes"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"

	fileBytes, err := os.ReadFile("metronome.mp3")
	if err != nil {
		panic("reading metronome.mp3 failed: " + err.Error())
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
  if err != nil {
    panic("mp3.NewDecoder failed: " + err.Error())
  }

  op := &oto.NewContextOptions{
    SampleRate: 44100,
    ChannelCount: 2,
    Format: oto.FormatSignedInt16LE,
  }

  otoCtx, readyChan, err := oto.NewContext(op)
  if err != nil {
    panic("oto.NewContext failed: " + err.Error())
  }
  <-readyChan

  player := otoCtx.NewPlayer(decodedMp3)

  player.Play()

  for player.IsPlaying() {
    time.Sleep(time.Millisecond)
  }

  err = player.Close()
  if err != nil {
      panic("player.Close failed: " + err.Error())
  }
*/

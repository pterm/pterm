package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Fullscreen takes over the whole terminal; WithCenter centers the
	// content within it.
	area, _ := pterm.DefaultArea.WithFullscreen().WithCenter().Start()

	// Each Update redraws the area in place instead of appending new lines.
	for i := range 5 {
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	// Stop clears the area and restores the terminal.
	area.Stop()
}

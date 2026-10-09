package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Fullscreen clears the terminal and gives the area the whole screen.
	area, _ := pterm.DefaultArea.WithFullscreen().Start()

	// Each Update redraws the area in place instead of appending new lines.
	for i := range 5 {
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	// Stop clears the area and restores the terminal.
	area.Stop()
}

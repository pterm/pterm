package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// WithCenter horizontally centers the area's content in the terminal.
	area, _ := pterm.DefaultArea.WithCenter().Start()

	// Each Update redraws the area in place instead of appending new lines.
	for i := range 5 {
		area.Update(pterm.Sprintfln("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	area.Stop()
}

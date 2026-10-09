package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	area, _ := pterm.DefaultArea.Start()

	// Each Update redraws the area in place instead of appending new lines.
	for i := range 5 {
		area.Update(pterm.Sprintfln("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	area.Stop()
}

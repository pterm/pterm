package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Start a new fullscreen area. This will return an area instance and an error.
	// The underscore (_) is used to ignore the error.
	area, _ := pterm.DefaultArea.WithFullscreen().Start()

	// Loop 5 times to update the area content.
	for i := 0; i < 5; i++ {
		// Update the content of the area with the current count.
		// The Sprintf function is used to format the string.
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))

		// Pause for a second before the next update.
		time.Sleep(time.Second)
	}

	// Stop the area after all updates are done.
	area.Stop()
}

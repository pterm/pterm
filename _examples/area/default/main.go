package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Start a new default area and get a reference to it.
	// The second return value is an error which is ignored here.
	area, _ := pterm.DefaultArea.Start()

	// Loop 5 times
	for i := 0; i < 5; i++ {
		// Update the content of the area dynamically.
		// Here we're just displaying the current count.
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))

		// Pause for a second before the next update.
		time.Sleep(time.Second)
	}

	// Stop the area after all updates are done.
	// This will clean up and free resources used by the area.
	area.Stop()
}

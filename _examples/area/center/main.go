package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Start a new default area in the center of the terminal.
	// The Start() function returns the created area and an error.
	area, _ := pterm.DefaultArea.WithCenter().Start()

	// Loop 5 times to simulate a dynamic update.
	for i := 0; i < 5; i++ {
		// Update the content of the area with the current count.
		// The Sprintf function is used to format the string.
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))

		// Pause for a second to simulate a time-consuming task.
		time.Sleep(time.Second)
	}

	// Stop the area after all updates are done.
	area.Stop()
}

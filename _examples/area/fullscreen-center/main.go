package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Initialize a new PTerm area with fullscreen and center options
	// The Start() function returns the created area and an error (ignored here)
	area, _ := pterm.DefaultArea.WithFullscreen().WithCenter().Start()

	// Loop 5 times to demonstrate dynamic content update
	for i := 0; i < 5; i++ {
		// Update the content of the area with the current count
		// The Sprintf function is used to format the string with the count
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))

		// Pause for a second
		time.Sleep(time.Second)
	}

	// Stop the area after all updates are done
	// This will clear the area and return the terminal to its normal state
	area.Stop()
}

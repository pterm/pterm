package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	timer, _ := pterm.DefaultTimer.Start() // Start the timer printer.
	time.Sleep(time.Second * 5)            // Simulate 5 seconds of work.
	timer.Stop()                           // Stop the timer.
}

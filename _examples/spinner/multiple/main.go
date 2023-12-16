package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Create a multi printer. This allows multiple spinners to print simultaneously.
	multi := pterm.DefaultMultiPrinter

	// Create and start spinner 1 with a new writer from the multi printer.
	// The spinner will display the message "Spinner 1".
	spinner1, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 1")

	// Create and start spinner 2 with a new writer from the multi printer.
	// The spinner will display the message "Spinner 2".
	spinner2, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 2")

	// Create and start spinner 3 with a new writer from the multi printer.
	// The spinner will display the message "Spinner 3".
	spinner3, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 3")

	// Start the multi printer. This will start printing all the spinners.
	multi.Start()

	// Wait for 1 second.
	time.Sleep(time.Millisecond * 1000)

	// Stop spinner 1 with a success message.
	spinner1.Success("Spinner 1 is done!")

	// Wait for 750 milliseconds.
	time.Sleep(time.Millisecond * 750)

	// Stop spinner 2 with a failure message.
	spinner2.Fail("Spinner 2 failed!")

	// Wait for 500 milliseconds.
	time.Sleep(time.Millisecond * 500)

	// Stop spinner 3 with a warning message.
	spinner3.Warning("Spinner 3 has a warning!")

	// Stop the multi printer. This will stop printing all the spinners.
	multi.Stop()
}

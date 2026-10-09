package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// The spinner animation is just a sequence of frames, so any set of
	// strings works. WithStyle colors the animation.
	spinner, _ := pterm.DefaultSpinner.
		WithSequence("▁", "▃", "▅", "▇", "▅", "▃").
		WithStyle(pterm.NewStyle(pterm.FgCyan)).
		Start("Uploading assets...")

	time.Sleep(time.Second * 2)

	// The text can change while the spinner keeps running.
	spinner.UpdateText("Finalizing upload...")
	time.Sleep(time.Second * 2)

	spinner.Success("Upload complete")
}

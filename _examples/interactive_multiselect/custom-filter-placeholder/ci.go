package main

import (
	"os"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

// ------ Automation for CI ------
// Simulates the user's keystrokes when CI=true, so the demo can run
// unattended and be recorded for the example animation. You can ignore this file.
func init() {
	if os.Getenv("CI") == "true" {
		go func() {
			time.Sleep(time.Second)

			// Filter the list down to the fuzzy searching options.
			for _, s := range "fuzzy" {
				keyboard.SimulateKeyPress(s)
				time.Sleep(time.Millisecond * 150)
			}

			time.Sleep(time.Millisecond * 500)

			// Toggle a few of the matches.
			keyboard.SimulateKeyPress(keys.Enter)
			time.Sleep(time.Millisecond * 300)
			keyboard.SimulateKeyPress(keys.Down)
			time.Sleep(time.Millisecond * 300)
			keyboard.SimulateKeyPress(keys.Enter)
			time.Sleep(time.Millisecond * 300)
			keyboard.SimulateKeyPress(keys.Down)
			time.Sleep(time.Millisecond * 300)
			keyboard.SimulateKeyPress(keys.Enter)
			time.Sleep(time.Millisecond * 500)

			// Confirm the selection.
			keyboard.SimulateKeyPress(keys.Tab)
			time.Sleep(time.Millisecond * 500)
		}()
	}
}

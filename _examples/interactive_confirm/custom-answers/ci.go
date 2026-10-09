package main

import (
	"os"
	"time"

	"atomicgo.dev/keyboard"
)

// ------ Automation for CI ------
// Simulates the user's keystrokes when CI=true, so the demo can run
// unattended and be recorded for the example animation. You can ignore this file.
func init() {
	if os.Getenv("CI") == "true" {
		go func() {
			time.Sleep(time.Second * 2)

			// "a" is the shorthand for the custom confirm text "Apply".
			keyboard.SimulateKeyPress('a')
		}()
	}
}

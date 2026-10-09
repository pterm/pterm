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
			input := "Hello, World!"
			for _, r := range input {
				if r == '\n' {
					keyboard.SimulateKeyPress(keys.Enter)
				} else {
					keyboard.SimulateKeyPress(r)
				}
				time.Sleep(time.Millisecond * 250)
			}

			keyboard.SimulateKeyPress(keys.Enter)
		}()
	}
}

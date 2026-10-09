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
			keyboard.SimulateKeyPress(keys.Down)
			time.Sleep(time.Millisecond * 100)
			keyboard.SimulateKeyPress(keys.Space)

			time.Sleep(time.Millisecond * 300)

			keyboard.SimulateKeyPress(keys.Down)
			time.Sleep(time.Millisecond * 100)
			keyboard.SimulateKeyPress(keys.Space)

			time.Sleep(time.Millisecond * 300)
			keyboard.SimulateKeyPress(keys.Enter)
		}()
	}
}

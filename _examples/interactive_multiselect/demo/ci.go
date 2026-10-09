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
			for i := 0; i < 10; i++ {
				keyboard.SimulateKeyPress(keys.Down)
				if i%2 == 0 {
					time.Sleep(time.Millisecond * 100)
					keyboard.SimulateKeyPress(keys.Enter)
				}
				time.Sleep(time.Millisecond * 500)
			}
			time.Sleep(time.Second)

			// Filter the list down to the fuzzy searching options.
			for _, s := range "fuzzy" {
				keyboard.SimulateKeyPress(s)
				time.Sleep(time.Millisecond * 150)
			}

			time.Sleep(time.Second)

			for i := 0; i < 2; i++ {
				keyboard.SimulateKeyPress(keys.Down)
				time.Sleep(time.Millisecond * 300)
			}

			keyboard.SimulateKeyPress(keys.Enter)
			time.Sleep(time.Millisecond * 350)
			keyboard.SimulateKeyPress(keys.Tab)
		}()
	}
}

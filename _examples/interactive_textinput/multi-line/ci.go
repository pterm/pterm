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
			input := "1111111\n2222222"
			for _, r := range input {
				if r == '\n' {
					keyboard.SimulateKeyPress(keys.Enter)
				} else {
					keyboard.SimulateKeyPress(r)
				}
				time.Sleep(time.Millisecond * 250)
			}

			for i := 0; i < 7; i++ {
				keyboard.SimulateKeyPress(keys.Left)
				time.Sleep(time.Millisecond * 150)
			}

			keyboard.SimulateKeyPress(keys.Backspace)
			time.Sleep(time.Millisecond * 500)
			keyboard.SimulateKeyPress(keys.Enter)
			time.Sleep(time.Millisecond * 500)
			input = "33333333\n4\n5555555"
			for _, r := range input {
				if r == '\n' {
					keyboard.SimulateKeyPress(keys.Enter)
				} else {
					keyboard.SimulateKeyPress(r)
				}
				time.Sleep(time.Millisecond * 250)
			}

			// Tab submits the multi-line input.
			keyboard.SimulateKeyPress(keys.Tab)
		}()
	}
}

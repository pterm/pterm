package main

import (
	"os"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

// ------ Automation for CI ------
// You can ignore this function, it is used to automatically run the demo and generate the example animation in our CI system.
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

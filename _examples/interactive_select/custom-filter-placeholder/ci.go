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

			// Type some search text
			for _, s := range "fuzzy" {
				keyboard.SimulateKeyPress(s)
				time.Sleep(time.Millisecond * 150)
			}

			time.Sleep(time.Second)

			// Navigate down a couple times
			keyboard.SimulateKeyPress(keys.Down)
			time.Sleep(time.Millisecond * 300)
			keyboard.SimulateKeyPress(keys.Down)
			time.Sleep(time.Millisecond * 300)

			// Select an option
			keyboard.SimulateKeyPress(keys.Enter)
			time.Sleep(time.Millisecond * 500)
		}()
	}
}

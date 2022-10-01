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

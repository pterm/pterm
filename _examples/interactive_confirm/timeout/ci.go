package main

import (
	"os"
	"time"

	"atomicgo.dev/keyboard"
)

// ------ Automation for CI ------
// You can ignore this function, it is used to automatically run the demo and generate the example animation in our CI system.
func init() {
	if os.Getenv("CI") == "true" {
		go func() {
			time.Sleep(time.Second * 2)
			keyboard.SimulateKeyPress('y')
		}()
	}
}

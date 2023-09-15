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
			for i := 0; i < 10; i++ {
				keyboard.SimulateKeyPress("j")
				time.Sleep(time.Millisecond * 250)
			}
			time.Sleep(time.Second)

			keyboard.SimulateKeyPress(keys.CtrlF)
			for _, s := range "fuzzy" {
				keyboard.SimulateKeyPress(s)
				time.Sleep(time.Millisecond * 150)
			}

			time.Sleep(time.Second)

			keyboard.SimulateKeyPress(keys.CtrlF)
			for i := 0; i < 2; i++ {
				keyboard.SimulateKeyPress("j")
				time.Sleep(time.Millisecond * 300)
			}

			keyboard.SimulateKeyPress(keys.Enter)
		}()
	}
}

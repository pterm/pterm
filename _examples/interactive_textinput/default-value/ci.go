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
			time.Sleep(time.Millisecond * 1500)
			input := "Hello; World!"
			for _, r := range []rune(input) {
				keyboard.SimulateKeyPress(r)
				time.Sleep(time.Millisecond * 250)
			}

			for i := 0; i < 7; i++ {
				keyboard.SimulateKeyPress(keys.Left)
				time.Sleep(time.Millisecond * 150)
			}

			keyboard.SimulateKeyPress(keys.Backspace)
			time.Sleep(time.Millisecond * 500)
			keyboard.SimulateKeyPress(',')
			time.Sleep(time.Millisecond * 500)
			keyboard.SimulateKeyPress(keys.Enter)
		}()
	}
}

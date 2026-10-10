package main

import "github.com/pterm/pterm"

func main() {
	// Options and their handles can be anything. Handles default to the
	// first letter of each option; here "keep" and "kill" would clash, so
	// they get their own.
	answer, _ := pterm.Continue.
		WithOptions([]string{"keep", "kill", "restart"}).
		WithHandles([]string{"k", "x", "r"}).
		WithDefaultValue("restart").
		Show("Port 8080 is taken by node (pid 4211)")

	pterm.Info.Printfln("You chose to %s it", answer)
}

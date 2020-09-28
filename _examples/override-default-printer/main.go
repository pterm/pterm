package main

import "github.com/pterm/pterm"

func main() {

	pterm.Error.Println("This is the default Error")

	pterm.Error.Prefix = pterm.Prefix{
		Text:  "OVERRIDE",
		Style: pterm.Style{pterm.BgCyan, pterm.FgRed},
	}

	pterm.Error.Println("This is the default Error after the prefix was overridden")
}

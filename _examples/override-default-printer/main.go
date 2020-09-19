package main

import "github.com/pterm/pterm"

func main() {

	pterm.ErrorPrinter.Println("This is the default ErrorPrinter")

	pterm.ErrorPrinter.Prefix = pterm.Prefix{
		Text:  "OVERRIDE",
		Style: pterm.Style{pterm.BgCyan, pterm.FgRed},
	}

	pterm.ErrorPrinter.Println("This is the default ErrorPrinter after the prefix was overridden")
}

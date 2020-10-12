package main

import "github.com/pterm/pterm"

func main() {
	// Simple Println with different colored words.
	pterm.Println(pterm.Red("Hello, ") + pterm.Green("World") + pterm.Cyan("!"))
	pterm.Println(pterm.Red("Even " + pterm.Cyan("nested ") + pterm.Green("colors ") + "are supported!"))
	pterm.FgBlack.Println("FgBlack")
	pterm.FgRed.Println("FgRed")
	pterm.FgGreen.Println("FgGreen")
	pterm.FgYellow.Println("FgYellow")
	pterm.FgBlue.Println("FgBlue")
	pterm.FgMagenta.Println("FgMagenta")
	pterm.FgCyan.Println("FgCyan")
	pterm.FgWhite.Println("FgWhite")

	pterm.FgLightRed.Println("FgLightRed")
	pterm.FgLightGreen.Println("FgLightGreen")
	pterm.FgLightYellow.Println("FgLightYellow")
	pterm.FgLightBlue.Println("FgLightBlue")
	pterm.FgLightMagenta.Println("FgLightMagenta")
	pterm.FgLightCyan.Println("FgLightCyan")
	pterm.FgLightWhite.Println("FgLightWhite")
}

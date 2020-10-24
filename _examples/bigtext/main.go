package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("PTerm")).Render()

	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("P", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).
		Render()
}

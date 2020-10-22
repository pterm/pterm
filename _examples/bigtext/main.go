package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultBigText.WithLetters(pterm.NewLettersFromText("PTerm")).Render()

	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromTextWithStyle("P", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromTextWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).
		Render()
}

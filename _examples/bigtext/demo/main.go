package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Big ASCII-art text in the default theme style. Handy for title screens.
	pterm.DefaultBigText.WithLetters(putils.LettersFromString("PTerm")).Render()

	// Each letter group can carry its own style, so parts of the text can be
	// colored independently.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Term", pterm.FgLightMagenta.ToStyle()),
	).Render()

	// TrueColor works too. On terminals without TrueColor support, PTerm
	// downsamples the RGB value automatically.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithRGB("PTerm", pterm.NewRGB(255, 215, 0)),
	).Render()
}

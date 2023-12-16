package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Create a large text with the LetterStyle from the standard theme.
	// This is useful for creating title screens.
	pterm.DefaultBigText.WithLetters(putils.LettersFromString("PTerm")).Render()

	// Create a large text with differently colored letters.
	// Here, the first letter 'P' is colored cyan and the rest 'Term' is colored light magenta.
	// This can be used to highlight specific parts of the text.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Term", pterm.FgLightMagenta.ToStyle()),
	).Render()

	// Create a large text with a specific RGB color.
	// This can be used when you need a specific color that is not available in the standard colors.
	// Here, the color is gold (RGB: 255, 215, 0).
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithRGB("PTerm", pterm.NewRGB(255, 215, 0)),
	).Render()
}

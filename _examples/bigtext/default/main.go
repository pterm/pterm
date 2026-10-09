package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// BigText takes Letters rather than a plain string; putils converts one
	// into the other.
	letters := putils.LettersFromString("PTerm")

	pterm.DefaultBigText.WithLetters(letters).Render()
}

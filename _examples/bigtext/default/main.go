package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	pterm.DefaultBigText.WithLetters(putils.LettersFromString("PTerm")).Render()
}

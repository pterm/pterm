package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Print a block of text centered in the terminal
	pterm.DefaultCenter.Println("This text is centered!\nIt centers the whole block by default.\nIn that way you can do stuff like this:")

	// Generate BigLetters and store in 's'
	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("PTerm")).Srender()

	// Print the BigLetters 's' centered in the terminal
	pterm.DefaultCenter.Println(s)

	// Print each line of the text separately centered in the terminal
	pterm.DefaultCenter.WithCenterEachLineSeparately().Println("This text is centered!\nBut each line is\ncentered\nseparately")
}

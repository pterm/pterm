package main

import "github.com/pterm/pterm"

func main() {
	// Customize the DefaultHeader with a cyan background, black text, and a margin of 15.
	pterm.DefaultHeader.WithMargin(15).WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("This is a custom header!")

	// Define a new HeaderPrinter with a red background, black text, and a margin of 20.
	newHeader := pterm.HeaderPrinter{
		TextStyle:       pterm.NewStyle(pterm.FgBlack),
		BackgroundStyle: pterm.NewStyle(pterm.BgRed),
		Margin:          20,
	}

	// Print the custom header using the new HeaderPrinter.
	newHeader.Println("This is a custom header!")
}

package main

import "github.com/pterm/pterm"

func main() {
	// All available options: https://pkg.go.dev/github.com/pterm/pterm#HeaderPrinter

	// Build on top of DefaultHeader
	pterm.DefaultHeader. // Use DefaultHeader as base
				WithMargin(15).                                    // Set Margin to 15
				WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)). // Set BackgroundStyle to Cyan
				WithTextStyle(pterm.NewStyle(pterm.FgBlack)).      // Set TextStyle to Black
				Println("This is a custom header!")                // Print header
	// Instead of printing the header you can set it to a variable.
	// You can then reuse your custom header.

	// Making a completely new HeaderPrinter
	newHeader := pterm.HeaderPrinter{
		TextStyle:       pterm.NewStyle(pterm.FgBlack),
		BackgroundStyle: pterm.NewStyle(pterm.BgRed),
		Margin:          20,
	}

	newHeader.Println("This is a custom header!")

}

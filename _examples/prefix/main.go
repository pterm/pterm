package main

import "github.com/pterm/pterm"

func main() {
	// Enable debug messages.
	pterm.EnableDebugMessages()

	pterm.Debug.Println("Hello, World!")                                                // Print Debug.
	pterm.Info.Println("Hello, World!")                                                 // Print Info.
	pterm.Success.Println("Hello, World!")                                              // Print Success.
	pterm.Warning.Println("Hello, World!")                                              // Print Warning.
	pterm.Error.Println("Errors show the filename and linenumber inside the terminal!") // Print Error.
	pterm.Info.WithShowLineNumber().Println("Other PrefixPrinters can do that too!")    // Print Info with line number.
	// Temporarily set Fatal to false, so that the CI won't actually panic.
	pterm.Fatal.WithFatal(false).Println("Hello, World!") // Print Fatal.

	// Spacer
	pterm.DefaultSection.WithBottomPadding(0).Println("Custom PrefixPrinters can be used too!")

	// --- Custom PrefixPrinter ---
	customPrefixPrinter := pterm.PrefixPrinter{
		Prefix: pterm.Prefix{
			Text:  "CUSTOM",
			Style: pterm.NewStyle(pterm.FgGreen, pterm.BgLightWhite),
		},
		Scope: pterm.Scope{
			Text:  "custom",
			Style: &pterm.ThemeDefault.ScopeStyle, // Use default theme scope style.
		},
		MessageStyle: pterm.NewStyle(pterm.FgRed),
	}

	customPrefixPrinter.Println("Custom PrefixPrinter")
}

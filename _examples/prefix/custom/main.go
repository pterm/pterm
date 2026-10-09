package main

import "github.com/pterm/pterm"

func main() {
	// A PrefixPrinter is just a struct, so custom printers can be built from
	// scratch. The built-ins (Info, Success, ...) are constructed the same way.
	deploy := pterm.PrefixPrinter{
		Prefix: pterm.Prefix{
			Text:  "DEPLOY",
			Style: pterm.NewStyle(pterm.BgLightMagenta, pterm.FgBlack),
		},
		// The scope is printed after the prefix in brackets. Useful for
		// tagging messages with a subsystem or component name.
		Scope: pterm.Scope{
			Text:  "database",
			Style: pterm.NewStyle(pterm.FgGray),
		},
		MessageStyle: pterm.NewStyle(pterm.FgLightMagenta),
	}

	deploy.Println("Running migrations...")
	deploy.Println("Migrations complete!")

	// Existing printers can be tweaked on the fly. With* methods return a
	// modified copy, so pterm.Info itself stays untouched.
	pterm.Info.WithScope(pterm.Scope{
		Text:  "api",
		Style: pterm.NewStyle(pterm.BgGray, pterm.FgLightWhite),
	}).Println("Listening on port 8080")
}

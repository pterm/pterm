### prefix/demo

![Animation](https://vhs.charm.sh/vhs-6vnMXqfVTtzXZnMoifwpNo.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Debug messages are hidden by default. Enable them so pterm.Debug prints.
	pterm.EnableDebugMessages()

	pterm.Debug.Println("Hello, World!")
	pterm.Info.Println("Hello, World!")
	pterm.Success.Println("Hello, World!")
	pterm.Warning.Println("Hello, World!")

	// Error prints the filename and line number of the call site.
	pterm.Error.Println("Errors show the filename and linenumber inside the terminal!")

	// Any PrefixPrinter can show line numbers via WithShowLineNumber.
	pterm.Info.WithShowLineNumber().Println("Other PrefixPrinters can do that too!")

	// Fatal would normally terminate the program. WithFatal(false) turns that
	// off so this demo keeps running.
	pterm.Fatal.WithFatal(false).Println("Hello, World!")
}
```

</details>

### prefix/custom

![Animation](https://vhs.charm.sh/vhs-67iqEMVRli3nNeB7gYEPxb.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
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
```

</details>


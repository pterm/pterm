# prefix/demo

![Animation](https://vhs.charm.sh/vhs-6vnMXqfVTtzXZnMoifwpNo.gif)

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

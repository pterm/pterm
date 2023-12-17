# prefix/demo

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Enable debug messages in PTerm.
	pterm.EnableDebugMessages()

	// Print a debug message with PTerm.
	pterm.Debug.Println("Hello, World!")

	// Print an informational message with PTerm.
	pterm.Info.Println("Hello, World!")

	// Print a success message with PTerm.
	pterm.Success.Println("Hello, World!")

	// Print a warning message with PTerm.
	pterm.Warning.Println("Hello, World!")

	// Print an error message with PTerm. This will also display the filename and line number in the terminal.
	pterm.Error.Println("Errors show the filename and linenumber inside the terminal!")

	// Print an informational message with PTerm, with line number.
	// This demonstrates that other PrefixPrinters can also display line numbers.
	pterm.Info.WithShowLineNumber().Println("Other PrefixPrinters can do that too!")

	// Temporarily set Fatal to false, so that the CI won't crash.
	// This will print a fatal message with PTerm, but won't terminate the program.
	pterm.Fatal.WithFatal(false).Println("Hello, World!")
}

```

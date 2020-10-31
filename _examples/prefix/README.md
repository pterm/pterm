# prefix

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Enable debug messages.
	pterm.EnableDebugMessages()

	pterm.Debug.Println("Hello, World!")   // Print Debug.
	pterm.Info.Println("Hello, World!")    // Print Info.
	pterm.Success.Println("Hello, World!") // Print Success.
	pterm.Warning.Println("Hello, World!") // Print Warning.
	pterm.Error.Println("Hello, World!")   // Print Error.
	// Temporarily set Fatal to false, so that the CI won't crash.
	pterm.Fatal.WithFatal(false).Println("Hello, World!") // Print Fatal.
}

```

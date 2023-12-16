# style/demo

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define a primary style with light cyan foreground, gray background, and bold text
	primary := pterm.NewStyle(pterm.FgLightCyan, pterm.BgGray, pterm.Bold)

	// Define a secondary style with light green foreground, white background, and italic text
	secondary := pterm.NewStyle(pterm.FgLightGreen, pterm.BgWhite, pterm.Italic)

	// Print "Hello, World!" with the primary style
	primary.Println("Hello, World!")

	// Print "Hello, World!" with the secondary style
	secondary.Println("Hello, World!")
}

```

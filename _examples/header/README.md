### header/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/header/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Print a default header.
	// This uses the default settings of PTerm to print a header.
	pterm.DefaultHeader.Println("This is the default header!")

	// Print a spacer line for better readability.
	pterm.Println()

	// Print a full-width header.
	// This uses the WithFullWidth() option of PTerm to print a header that spans the full width of the terminal.
	pterm.DefaultHeader.WithFullWidth().Println("This is a full-width header.")
}

```

</details>

### header/custom

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/header/custom/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
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

```

</details>


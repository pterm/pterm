# Quick Start - Add PTerm to your project

## Adding PTerm to your project

Just make sure your project has [`go modules`]() integration enabled.
Adding PTerm to your project is as simple as executing this comment inside your project directory:

```commandline
go get -u github.com/pterm/pterm
```

## Write your first `Hello World` application with PTerm

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a new header as a fork from pterm.DefaultHeader.
	// ┌ new header variable
	// │                 ┌ Fork it from the default header
	// │                 │            ┌ Set options
	header := pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgRed))

	// Print the header centered in your terminal.
	//      ┌ Use the default CenterPrinter
	//      │              ┌ Print a string ending with a new line
	//      │              │      ┌ Use our new header to format the input string
	pterm.DefaultCenter.Println(header.Sprint("Hello, World"))

	// Print a big text to the terminal.
	//          ┌ Use the default BigTextPrinter
	//          │              ┌ Set the Letters option
	//          │              │                   ┌ Generate Letters from string
	//          │              │                   │                            ┌ Render output to the console
	_ = pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Hello")).Render()

	// ┌──────────────────────────────────────────────────────────┐
	// │There are many more features, which are waiting for you :)│
	// └──────────────────────────────────────────────────────────┘

	// TODO: If you want, you can try to make the big text centered.
}

```
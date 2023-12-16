# section/demo

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a section with level one and print it.
	pterm.DefaultSection.Println("This is a section!")

	// Print an informational message.
	pterm.Info.Println("And here is some text.\nThis text could be anything.\nBasically it's just a placeholder")

	// Create a section with level two and print it.
	pterm.DefaultSection.WithLevel(2).Println("This is another section!")

	// Print another informational message.
	pterm.Info.Println("And this is\nmore placeholder text")
}

```

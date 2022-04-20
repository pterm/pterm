# section/demo

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Print a section with level one.
	pterm.DefaultSection.Println("This is a section!")
	// Print placeholder.
	pterm.Info.Println("And here is some text.\nThis text could be anything.\nBasically it's just a placeholder")

	// Print a section with level two.
	pterm.DefaultSection.WithLevel(2).Println("This is another section!")
	// Print placeholder.
	pterm.Info.Println("And this is\nmore placeholder text")
}

```

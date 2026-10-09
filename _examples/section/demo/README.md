# section/demo

![Animation](https://vhs.charm.sh/vhs-1Wv6LwKgaw4o36p5quakV1.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Sections structure long output into headed blocks, like chapters.
	pterm.DefaultSection.Println("This is a section!")
	pterm.Info.Println("And here is some text.\nThis text could be anything.\nBasically it's just a placeholder")

	// Higher levels render as smaller subsections.
	pterm.DefaultSection.WithLevel(2).Println("This is another section!")
	pterm.Info.Println("And this is\nmore placeholder text")
}
```

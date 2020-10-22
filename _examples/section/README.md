# section

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultSection.Println("This is a section!")
	pterm.Info.Println("And here is some letters.\nThis letters could be anything.\nBasically it's just a placeholder")
	pterm.DefaultSection.Println("This is another section!")
	pterm.Info.Println("And this is\nmore placeholder letters")
}

```

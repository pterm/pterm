### header/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/header/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Print a default header.
	pterm.DefaultHeader.Println("This is the default header!")
	pterm.Println() // spacer
	pterm.DefaultHeader.WithFullWidth().Println("This is a full-width header.")
}

```

</details>


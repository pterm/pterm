### interactive_continue/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_continue/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	result, _ := pterm.DefaultInteractiveContinue.Show()
	pterm.Println() // Blank line
	pterm.Info.Printfln("You answered: %s", result)
}

```

</details>


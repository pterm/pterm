### center/demo

![Animation](https://vhs.charm.sh/vhs-3bSfOBmdMLfHkRMvVNkHNe.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// By default the whole block is centered as one unit, so the lines keep
	// their left alignment relative to each other.
	pterm.DefaultCenter.Println("This text is centered!\nIt centers the whole block by default.\nIn that way you can do stuff like this:")

	// That makes it safe to center multiline output from other printers,
	// like a BigText rendered to a string.
	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("PTerm")).Srender()
	pterm.DefaultCenter.Println(s)

	// WithCenterEachLineSeparately centers every line on its own instead.
	pterm.DefaultCenter.WithCenterEachLineSeparately().Println("This text is centered!\nBut each line is\ncentered\nseparately")
}
```

</details>


### bulletlist/customized

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/bulletlist/customized/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Print a customized list with different styles and levels.
	pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "Blue", TextStyle: pterm.NewStyle(pterm.FgBlue), BulletStyle: pterm.NewStyle(pterm.FgRed)},
		{Level: 1, Text: "Green", TextStyle: pterm.NewStyle(pterm.FgGreen), Bullet: "-", BulletStyle: pterm.NewStyle(pterm.FgLightWhite)},
		{Level: 2, Text: "Cyan", TextStyle: pterm.NewStyle(pterm.FgCyan), Bullet: ">", BulletStyle: pterm.NewStyle(pterm.FgYellow)},
	}).Render()
}

```

</details>

### bulletlist/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/bulletlist/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Print a list with different levels.
	// Useful to generate lists automatically from data.
	pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "Level 0"},
		{Level: 1, Text: "Level 1"},
		{Level: 2, Text: "Level 2"},
	}).Render()

	// Convert a text to a list and print it.
	putils.BulletListFromString(`0
 1
  2
   3`, " ").Render()
}

```

</details>


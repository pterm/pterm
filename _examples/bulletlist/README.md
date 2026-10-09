### bulletlist/demo

![Animation](https://vhs.charm.sh/vhs-6ybH1SVurQNm5WaatnNAkF.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Level controls the indentation depth of each item.
	bulletListItems := []pterm.BulletListItem{
		{Level: 0, Text: "Level 0"},
		{Level: 1, Text: "Level 1"},
		{Level: 2, Text: "Level 2"},
	}

	pterm.DefaultBulletList.WithItems(bulletListItems).Render()

	// Alternatively, build a list from an indented string. The second
	// argument is the indent unit; one leading space equals one level here.
	text := `0
 1
  2
   3`

	putils.BulletListFromString(text, " ").Render()
}
```

</details>

### bulletlist/customized

![Animation](https://vhs.charm.sh/vhs-dcYDU9coxQOWTWt9hyDt0.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Every item can override the bullet symbol and style the bullet and
	// text independently. Level controls the indentation depth.
	bulletListItems := []pterm.BulletListItem{
		{
			Level:       0,
			Text:        "Blue",
			TextStyle:   pterm.NewStyle(pterm.FgBlue),
			BulletStyle: pterm.NewStyle(pterm.FgRed),
		},
		{
			Level:       1,
			Text:        "Green",
			TextStyle:   pterm.NewStyle(pterm.FgGreen),
			Bullet:      "-",
			BulletStyle: pterm.NewStyle(pterm.FgLightWhite),
		},
		{
			Level:       2,
			Text:        "Cyan",
			TextStyle:   pterm.NewStyle(pterm.FgCyan),
			Bullet:      ">",
			BulletStyle: pterm.NewStyle(pterm.FgYellow),
		},
	}

	pterm.DefaultBulletList.WithItems(bulletListItems).Render()
}
```

</details>


# bulletlist/customized

![Animation](https://vhs.charm.sh/vhs-dcYDU9coxQOWTWt9hyDt0.gif)

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

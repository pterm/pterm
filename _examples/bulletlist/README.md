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
	// Define a list of bullet list items with different levels.
	bulletListItems := []pterm.BulletListItem{
		{Level: 0, Text: "Level 0"}, // Level 0 item
		{Level: 1, Text: "Level 1"}, // Level 1 item
		{Level: 2, Text: "Level 2"}, // Level 2 item
	}

	// Use the default bullet list style to render the list items.
	pterm.DefaultBulletList.WithItems(bulletListItems).Render()

	// Define a string with different levels of indentation.
	text := `0
 1
  2
   3`

	// Convert the indented string to a bullet list and render it.
	putils.BulletListFromString(text, " ").Render()
}

```

</details>

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
	// Define a list of bullet list items with different styles and levels.
	bulletListItems := []pterm.BulletListItem{
		{
			Level:       0,                            // Level 0 (top level)
			Text:        "Blue",                       // Text to display
			TextStyle:   pterm.NewStyle(pterm.FgBlue), // Text color
			BulletStyle: pterm.NewStyle(pterm.FgRed),  // Bullet color
		},
		{
			Level:       1,                                  // Level 1 (sub-item)
			Text:        "Green",                            // Text to display
			TextStyle:   pterm.NewStyle(pterm.FgGreen),      // Text color
			Bullet:      "-",                                // Custom bullet symbol
			BulletStyle: pterm.NewStyle(pterm.FgLightWhite), // Bullet color
		},
		{
			Level:       2,                              // Level 2 (sub-sub-item)
			Text:        "Cyan",                         // Text to display
			TextStyle:   pterm.NewStyle(pterm.FgCyan),   // Text color
			Bullet:      ">",                            // Custom bullet symbol
			BulletStyle: pterm.NewStyle(pterm.FgYellow), // Bullet color
		},
	}

	// Create a bullet list with the defined items and render it.
	pterm.DefaultBulletList.WithItems(bulletListItems).Render()
}

```

</details>


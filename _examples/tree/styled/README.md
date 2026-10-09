# tree/styled

![Animation](https://vhs.charm.sh/vhs-6695FaObbMrAZ8wcnhMb6T.gif)

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	tree := pterm.TreeNode{
		Text: "project",
		Children: []pterm.TreeNode{
			{Text: "cmd", Children: []pterm.TreeNode{
				{Text: "main.go"},
			}},
			{Text: "internal", Children: []pterm.TreeNode{
				{Text: "server.go"},
				{Text: "config.go"},
			}},
			{Text: "go.mod"},
			{Text: "README.md"},
		},
	}

	// WithTreeStyle colors the branch lines, WithTextStyle the node text.
	pterm.DefaultTree.
		WithRoot(tree).
		WithTreeStyle(pterm.NewStyle(pterm.FgLightBlue)).
		WithTextStyle(pterm.NewStyle(pterm.FgLightGreen)).
		Render()
}
```

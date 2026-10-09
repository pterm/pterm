# tree/demo

![Animation](https://vhs.charm.sh/vhs-7jq6l6QS4HH2hVPCWqkncl.gif)

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// A tree is described by nesting TreeNodes; each node holds its text and
	// its children.
	tree := pterm.TreeNode{
		Text: "Top node",
		Children: []pterm.TreeNode{{
			Text: "Child node",
			Children: []pterm.TreeNode{
				{Text: "Grandchild node"},
				{Text: "Grandchild node"},
				{Text: "Grandchild node"},
			},
		}},
	}

	pterm.DefaultTree.WithRoot(tree).Render()
}
```

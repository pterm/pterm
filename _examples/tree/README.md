### tree/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/tree/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define a tree structure using pterm.TreeNode
	tree := pterm.TreeNode{
		// The top node of the tree
		Text: "Top node",
		// The children of the top node
		Children: []pterm.TreeNode{{
			// A child node
			Text: "Child node",
			// The children of the child node
			Children: []pterm.TreeNode{
				// Grandchildren nodes
				{Text: "Grandchild node"},
				{Text: "Grandchild node"},
				{Text: "Grandchild node"},
			},
		}},
	}

	// Render the tree with the defined structure as the root
	pterm.DefaultTree.WithRoot(tree).Render()
}

```

</details>

### tree/from-leveled-list

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/tree/from-leveled-list/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Define a leveled list to represent the structure of the directories.
	leveledList := pterm.LeveledList{
		{Level: 0, Text: "C:"},
		{Level: 1, Text: "Users"},
		{Level: 1, Text: "Windows"},
		{Level: 1, Text: "Programs"},
		{Level: 1, Text: "Programs(x86)"},
		{Level: 1, Text: "dev"},
		{Level: 0, Text: "D:"},
		{Level: 0, Text: "E:"},
		{Level: 1, Text: "Movies"},
		{Level: 1, Text: "Music"},
		{Level: 2, Text: "LinkinPark"},
		{Level: 1, Text: "Games"},
		{Level: 2, Text: "Shooter"},
		{Level: 3, Text: "CallOfDuty"},
		{Level: 3, Text: "CS:GO"},
		{Level: 3, Text: "Battlefield"},
		{Level: 4, Text: "Battlefield 1"},
		{Level: 4, Text: "Battlefield 2"},
		{Level: 0, Text: "F:"},
		{Level: 1, Text: "dev"},
		{Level: 2, Text: "dops"},
		{Level: 2, Text: "PTerm"},
	}

	// Convert the leveled list into a tree structure.
	root := putils.TreeFromLeveledList(leveledList)
	root.Text = "Computer" // Set the root node text.

	// Render the tree structure using the default tree printer.
	pterm.DefaultTree.WithRoot(root).Render()
}

```

</details>


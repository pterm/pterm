### tree/demo

![Animation](https://vhs.charm.sh/vhs-7jq6l6QS4HH2hVPCWqkncl.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### tree/from-leveled-list

![Animation](https://vhs.charm.sh/vhs-5m4lW2DbvrN4pgynZUdVvO.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// A LeveledList is a flat alternative to nesting TreeNodes by hand: each
	// entry states its own depth.
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

	// TreeFromLeveledList converts the list into a TreeNode; the returned
	// root just needs a name.
	root := putils.TreeFromLeveledList(leveledList)
	root.Text = "Computer"

	pterm.DefaultTree.WithRoot(root).Render()
}
```

</details>

### tree/styled

![Animation](https://vhs.charm.sh/vhs-6695FaObbMrAZ8wcnhMb6T.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>


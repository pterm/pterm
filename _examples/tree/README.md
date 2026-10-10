### tree/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Descriptions line up in one muted column, however deep the node sits.
	pterm.Tree.WithNodes(
		pterm.TreeNode{Text: "doctor", Description: "diagnostics"},
		pterm.TreeNode{Text: "mcp", Description: "manage MCP servers", Children: []pterm.TreeNode{
			{Text: "start", Description: "boots a server"},
			{Text: "stop", Description: "stops a running server"},
		}},
		pterm.TreeNode{Text: "disk", Description: "local disk utilities", Children: []pterm.TreeNode{
			{Text: "usage", Description: "shows what takes up space"},
			{Text: "clean"},
		}},
	).Render()
}
```

</details>

### tree/from-indented

![Animation](from-indented/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Write the tree the way you would sketch it in a README. Each level
	// only has to be indented further than its parent.
	files := putils.TreeFromIndented(`
cmd
  pterm
    main.go
internal
  screen
    screen.go
    region.go
  text
    width.go
go.mod`)

	_ = pterm.Tree.WithRoot(files).Render()
}
```

</details>

### tree/from-leveled-list

![Animation](from-leveled-list/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A LeveledList is a flat alternative to nesting TreeNodes by hand: each
	// item states its own depth.
	list := pterm.LeveledList{
		{Level: 0, Text: "C:"},
		{Level: 1, Text: "Users"},
		{Level: 1, Text: "Windows"},
		{Level: 1, Text: "Programs"},
		{Level: 0, Text: "D:"},
		{Level: 1, Text: "Movies"},
		{Level: 1, Text: "Games"},
		{Level: 2, Text: "Shooter"},
		{Level: 3, Text: "Battlefield"},
		{Level: 0, Text: "E:"},
		{Level: 1, Text: "dev"},
		{Level: 2, Text: "pterm"},
	}

	root := list.Tree()
	root.Text = "Computer"

	pterm.Tree.WithRoot(root).Render()
}
```

</details>

### tree/root

![Animation](root/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// The text of the root goes on the first line, above its children.
	pterm.Tree.WithRoot(pterm.TreeNode{
		Text: "pterm",
		Children: []pterm.TreeNode{
			{Text: "_examples", Children: []pterm.TreeNode{
				{Text: "table"},
				{Text: "tree"},
			}},
			{Text: "internal", Children: []pterm.TreeNode{
				{Text: "screen", Description: "the renderer"},
				{Text: "text", Description: "width, wrapping and cutting"},
			}},
			{Text: "go.mod"},
			{Text: "README.md"},
		},
	}).Render()
}
```

</details>

### tree/styled

![Animation](styled/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Every part has its own tone. Here the node names are Accent and the
	// branches get one more column of line per level.
	pterm.Tree.
		WithTextStyle(pterm.ToneAccent.Style()).
		WithLineTone(pterm.ToneMuted).
		WithIndent(1).
		WithNodes(
			pterm.TreeNode{Text: "cmd", Children: []pterm.TreeNode{
				{Text: "main.go"},
			}},
			pterm.TreeNode{Text: "internal", Children: []pterm.TreeNode{
				{Text: "server.go"},
				{Text: "config.go"},
			}},
			pterm.TreeNode{Text: "go.mod"},
		).
		Render()
}
```

</details>


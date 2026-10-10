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

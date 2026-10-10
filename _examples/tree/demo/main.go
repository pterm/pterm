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

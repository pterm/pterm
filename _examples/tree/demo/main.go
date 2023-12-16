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

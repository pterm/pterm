package putils

import "github.com/pterm/pterm"

// TreeFromLeveledList converts a TreeItems list to a TreeNode and returns it.
func TreeFromLeveledList(leveledListItems pterm.LeveledList) pterm.TreeNode {
	if len(leveledListItems) == 0 {
		return pterm.TreeNode{}
	}

	root := &pterm.TreeNode{
		Children: []pterm.TreeNode{},
	}

	for i, record := range leveledListItems {
		last := root

		if record.Level < 0 {
			record.Level = 0
			leveledListItems[i].Level = 0
		}

		if len(leveledListItems)-1 != i {
			if leveledListItems[i+1].Level-1 > record.Level {
				leveledListItems[i+1].Level = record.Level + 1
			}
		}

		for i := 0; i < record.Level; i++ {
			lastIndex := len(last.Children) - 1
			last = &last.Children[lastIndex]
		}
		last.Children = append(last.Children, pterm.TreeNode{
			Children: []pterm.TreeNode{},
			Text:     record.Text,
		})
	}

	return *root
}

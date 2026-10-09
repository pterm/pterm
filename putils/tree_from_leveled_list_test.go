package putils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestTreeFromLeveledListEmpty(t *testing.T) {
	assert.Equal(t, pterm.TreeNode{}, TreeFromLeveledList(pterm.LeveledList{}))
	assert.Equal(t, pterm.TreeNode{}, TreeFromLeveledList(nil))
}

func TestTreeFromLeveledListFlat(t *testing.T) {
	tree := TreeFromLeveledList(pterm.LeveledList{
		{Level: 0, Text: "a"},
		{Level: 0, Text: "b"},
	})

	expected := pterm.TreeNode{
		Children: []pterm.TreeNode{
			{Children: []pterm.TreeNode{}, Text: "a"},
			{Children: []pterm.TreeNode{}, Text: "b"},
		},
	}
	assert.Equal(t, expected, tree)
}

func TestTreeFromLeveledListNested(t *testing.T) {
	tree := TreeFromLeveledList(pterm.LeveledList{
		{Level: 0, Text: "root"},
		{Level: 1, Text: "child"},
		{Level: 2, Text: "grandchild"},
		{Level: 1, Text: "second child"},
		{Level: 0, Text: "second root"},
	})

	expected := pterm.TreeNode{
		Children: []pterm.TreeNode{
			{
				Text: "root",
				Children: []pterm.TreeNode{
					{
						Text: "child",
						Children: []pterm.TreeNode{
							{Children: []pterm.TreeNode{}, Text: "grandchild"},
						},
					},
					{Children: []pterm.TreeNode{}, Text: "second child"},
				},
			},
			{Children: []pterm.TreeNode{}, Text: "second root"},
		},
	}
	assert.Equal(t, expected, tree)
}

func TestTreeFromLeveledListClampsLevelJumps(t *testing.T) {
	// A level may only increase by one per item; bigger jumps are clamped.
	tree := TreeFromLeveledList(pterm.LeveledList{
		{Level: 0, Text: "a"},
		{Level: 5, Text: "b"},
	})

	expected := pterm.TreeNode{
		Children: []pterm.TreeNode{
			{
				Text: "a",
				Children: []pterm.TreeNode{
					{Children: []pterm.TreeNode{}, Text: "b"},
				},
			},
		},
	}
	assert.Equal(t, expected, tree)
}

func TestTreeFromLeveledListNegativeLevelBecomesRootLevel(t *testing.T) {
	tree := TreeFromLeveledList(pterm.LeveledList{
		{Level: -5, Text: "a"},
		{Level: 0, Text: "b"},
	})

	expected := pterm.TreeNode{
		Children: []pterm.TreeNode{
			{Children: []pterm.TreeNode{}, Text: "a"},
			{Children: []pterm.TreeNode{}, Text: "b"},
		},
	}
	assert.Equal(t, expected, tree)
}

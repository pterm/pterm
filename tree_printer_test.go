package pterm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreePrinterNilPrint(t *testing.T) {
	proxyToDevNull()
	TreePrinter{}.Render()
	TreePrinter{}.WithRoot(NewTreeFromLeveledList(LeveledList{LeveledListItem{Text: "Hello, World!", Level: 0}})).Render()
}

func TestTreePrinter_Render(t *testing.T) {
	DefaultTree.WithRoot(NewTreeFromLeveledList([]LeveledListItem{
		{Level: 0, Text: fmt.Sprint("Hello, World!")},
		{Level: 0, Text: "0.0"},
		{Level: 1, Text: "0.1"},
		{Level: 1, Text: "0.2"},
		{Level: 0, Text: "1.0"},
		{Level: 0, Text: "2.0"},
		{Level: 1, Text: "2.1"},
		{Level: 1, Text: "2.2"},
		{Level: 2, Text: "2.2.1"},
		{Level: 1, Text: "2.3"},
	})).Render()
}

func TestTreePrinter_NewTreeFromLeveledList(t *testing.T) {
	p := DefaultTree
	p2 := p.WithRoot(NewTreeFromLeveledList(LeveledList{
		LeveledListItem{Level: 0, Text: "0.0"},
		LeveledListItem{Level: 1, Text: "0.1"},
		LeveledListItem{Level: 1, Text: "0.2"},
		LeveledListItem{Level: 0, Text: "1.0"},
		LeveledListItem{Level: 0, Text: "2.0"},
		LeveledListItem{Level: 1, Text: "2.1"},
		LeveledListItem{Level: 1, Text: "2.2"},
		LeveledListItem{Level: 2, Text: "2.2.1"},
		LeveledListItem{Level: 1, Text: "2.3"}}))

	assert.Equal(t, NewTreeFromLeveledList(LeveledList{
		LeveledListItem{Level: 0, Text: "0.0"},
		LeveledListItem{Level: 1, Text: "0.1"},
		LeveledListItem{Level: 1, Text: "0.2"},
		LeveledListItem{Level: 0, Text: "1.0"},
		LeveledListItem{Level: 0, Text: "2.0"},
		LeveledListItem{Level: 1, Text: "2.1"},
		LeveledListItem{Level: 1, Text: "2.2"},
		LeveledListItem{Level: 2, Text: "2.2.1"},
		LeveledListItem{Level: 1, Text: "2.3"}}), p2.Root)
	assert.Empty(t, p.Root)
}

func TestTreePrinter_NewTreeFromLeveledListLevelInvalidIncrease(t *testing.T) {
	p := DefaultTree
	p2 := p.WithRoot(NewTreeFromLeveledList(LeveledList{
		LeveledListItem{Level: 0, Text: "0.0"},
		LeveledListItem{Level: 1, Text: "0.1"},
		LeveledListItem{Level: 1, Text: "0.2"},
		LeveledListItem{Level: 0, Text: "1.0"},
		LeveledListItem{Level: 0, Text: "2.0"},
		LeveledListItem{Level: 1, Text: "2.1"},
		LeveledListItem{Level: 1, Text: "2.2"},
		LeveledListItem{Level: 2, Text: "2.2.1"},
		LeveledListItem{Level: 10, Text: "2.3"}}))

	assert.Equal(t, NewTreeFromLeveledList(LeveledList{
		LeveledListItem{Level: 0, Text: "0.0"},
		LeveledListItem{Level: 1, Text: "0.1"},
		LeveledListItem{Level: 1, Text: "0.2"},
		LeveledListItem{Level: 0, Text: "1.0"},
		LeveledListItem{Level: 0, Text: "2.0"},
		LeveledListItem{Level: 1, Text: "2.1"},
		LeveledListItem{Level: 1, Text: "2.2"},
		LeveledListItem{Level: 2, Text: "2.2.1"},
		LeveledListItem{Level: 3, Text: "2.3"}}), p2.Root)
	assert.Empty(t, p.Root)
}

func TestTreePrinter_NewTreeFromLeveledListEmptyList(t *testing.T) {
	p := DefaultTree
	p2 := p.WithRoot(NewTreeFromLeveledList(LeveledList{}))

	assert.Equal(t, NewTreeFromLeveledList(LeveledList{}), p2.Root)
	assert.Empty(t, p.Root)
}

func TestTreePrinter_NewTreeFromLeveledListNegativeLevel(t *testing.T) {
	p := DefaultTree
	p2 := p.WithRoot(NewTreeFromLeveledList(LeveledList{
		LeveledListItem{Level: 0, Text: "0.0"},
		LeveledListItem{Level: 1, Text: "0.1"},
		LeveledListItem{Level: 1, Text: "0.2"},
		LeveledListItem{Level: 0, Text: "1.0"},
		LeveledListItem{Level: 0, Text: "2.0"},
		LeveledListItem{Level: 1, Text: "2.1"},
		LeveledListItem{Level: 1, Text: "2.2"},
		LeveledListItem{Level: 2, Text: "2.2.1"},
		LeveledListItem{Level: -5, Text: "2.3"}}))

	assert.Equal(t, NewTreeFromLeveledList(LeveledList{
		LeveledListItem{Level: 0, Text: "0.0"},
		LeveledListItem{Level: 1, Text: "0.1"},
		LeveledListItem{Level: 1, Text: "0.2"},
		LeveledListItem{Level: 0, Text: "1.0"},
		LeveledListItem{Level: 0, Text: "2.0"},
		LeveledListItem{Level: 1, Text: "2.1"},
		LeveledListItem{Level: 1, Text: "2.2"},
		LeveledListItem{Level: 2, Text: "2.2.1"},
		LeveledListItem{Level: 0, Text: "2.3"}}), p2.Root)
	assert.Empty(t, p.Root)
}

func TestTreePrinter_WithHorizontalString(t *testing.T) {
	p := TreePrinter{}
	p2 := p.WithHorizontalString("-")

	assert.Equal(t, "-", p2.HorizontalString)
	assert.Empty(t, p.HorizontalString)
}

func TestTreePrinter_WithRoot(t *testing.T) {
	p := TreePrinter{}
	p2 := p.WithRoot(TreeNode{
		Children: nil,
		Text:     "Hello, World!",
	})

	assert.Equal(t, TreeNode{
		Children: nil,
		Text:     "Hello, World!",
	}, p2.Root)
	assert.Empty(t, p.Root)
}

func TestTreePrinter_WithTreeStyle(t *testing.T) {
	p := TreePrinter{}
	s := NewStyle(FgRed, BgRed, Bold)
	p2 := p.WithTreeStyle(s)

	assert.Equal(t, s, p2.TreeStyle)
	assert.Empty(t, p.TreeStyle)
}

func TestTreePrinter_WithTextStyle(t *testing.T) {
	p := TreePrinter{}
	s := NewStyle(FgRed, BgRed, Bold)
	p2 := p.WithTextStyle(s)

	assert.Equal(t, s, p2.TextStyle)
	assert.Empty(t, p.TextStyle)
}

func TestTreePrinter_WithTopRightCornerString(t *testing.T) {
	p := TreePrinter{}
	p2 := p.WithTopRightCornerString("-")

	assert.Equal(t, "-", p2.TopRightCornerString)
	assert.Empty(t, p.TopRightCornerString)
}

func TestTreePrinter_WithTopRightDownStringOngoing(t *testing.T) {
	p := TreePrinter{}
	p2 := p.WithTopRightDownStringOngoing("-")

	assert.Equal(t, "-", p2.TopRightDownString)
	assert.Empty(t, p.TopRightDownString)
}

func TestTreePrinter_WithVerticalString(t *testing.T) {
	p := TreePrinter{}
	p2 := p.WithVerticalString("-")

	assert.Equal(t, "-", p2.VerticalString)
	assert.Empty(t, p.VerticalString)
}

func TestTreePrinter_WithIndent(t *testing.T) {
	p := TreePrinter{}
	p2 := p.WithIndent(3)

	assert.Equal(t, 3, p2.Indent)
	assert.Empty(t, p.Indent)
}

func TestTreePrinter_WithIndentInvalid(t *testing.T) {
	p := TreePrinter{}
	p2 := p.WithIndent(0)

	assert.Equal(t, 1, p2.Indent)
	assert.Empty(t, p.Indent)
}

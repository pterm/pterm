package pterm

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestTreePrinterNilPrint(t *testing.T) {
	proxyToDevNull()
	Tree{}.Render()
	Tree{}.WithRoot(NewTreeFromLeveledList(LeveledList{LeveledListItem{Text: "Hello, World!", Level: 0}})).Render()
}

func TestTree_Render(t *testing.T) {
	internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
		DefaultTree.WithRoot(NewTreeFromLeveledList([]LeveledListItem{
			{Level: 0, Text: fmt.Sprint(a)},
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
	})
}

func TestTree_NewTreeFromLeveledList(t *testing.T) {
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

func TestTree_NewTreeFromLeveledListLevelInvalidIncrease(t *testing.T) {
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

func TestTree_NewTreeFromLeveledListEmptyList(t *testing.T) {
	p := DefaultTree
	p2 := p.WithRoot(NewTreeFromLeveledList(LeveledList{}))

	assert.Equal(t, NewTreeFromLeveledList(LeveledList{}), p2.Root)
	assert.Empty(t, p.Root)
}

func TestTree_NewTreeFromLeveledListNegativeLevel(t *testing.T) {
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

func TestTree_WithHorizontalString(t *testing.T) {
	p := Tree{}
	p2 := p.WithHorizontalString("-")

	assert.Equal(t, "-", p2.HorizontalString)
	assert.Empty(t, p.HorizontalString)
}

func TestTree_WithRoot(t *testing.T) {
	p := Tree{}
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

func TestTree_WithTreeStyle(t *testing.T) {
	p := Tree{}
	s := NewStyle(FgRed, BgRed, Bold)
	p2 := p.WithTreeStyle(s)

	assert.Equal(t, s, p2.TreeStyle)
	assert.Empty(t, p.TreeStyle)
}

func TestTree_WithTextStyle(t *testing.T) {
	p := Tree{}
	s := NewStyle(FgRed, BgRed, Bold)
	p2 := p.WithTextStyle(s)

	assert.Equal(t, s, p2.TextStyle)
	assert.Empty(t, p.TextStyle)
}

func TestTree_WithTopRightCornerString(t *testing.T) {
	p := Tree{}
	p2 := p.WithTopRightCornerString("-")

	assert.Equal(t, "-", p2.TopRightCornerString)
	assert.Empty(t, p.TopRightCornerString)
}

func TestTree_WithTopRightDownStringOngoing(t *testing.T) {
	p := Tree{}
	p2 := p.WithTopRightDownStringOngoing("-")

	assert.Equal(t, "-", p2.TopRightDownString)
	assert.Empty(t, p.TopRightDownString)
}

func TestTree_WithVerticalString(t *testing.T) {
	p := Tree{}
	p2 := p.WithVerticalString("-")

	assert.Equal(t, "-", p2.VerticalString)
	assert.Empty(t, p.VerticalString)
}

func TestTree_WithIndent(t *testing.T) {
	p := Tree{}
	p2 := p.WithIndent(3)

	assert.Equal(t, 3, p2.Indent)
	assert.Empty(t, p.Indent)
}

func TestTree_WithIndentInvalid(t *testing.T) {
	p := Tree{}
	p2 := p.WithIndent(0)

	assert.Equal(t, 1, p2.Indent)
	assert.Empty(t, p.Indent)
}

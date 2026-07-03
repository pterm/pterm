package pterm_test

import (
	"os"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestTreePrinterNilPrint(t *testing.T) {
	_ = pterm.TreePrinter{}.Render()

	printer := pterm.TreePrinter{}.WithRoot(pterm.NewTreeFromLeveledList(pterm.LeveledList{pterm.LeveledListItem{Text: "Hello, World!", Level: 0}}))
	content, err := printer.Srender()

	assert.NoError(t, err)
	assert.NotNil(t, content)
}

func TestTreePrinter_Render(_ *testing.T) {
	_ = pterm.DefaultTree.WithRoot(pterm.NewTreeFromLeveledList([]pterm.LeveledListItem{
		{Level: 0, Text: "Hello, World!"},
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
	p := pterm.DefaultTree
	p2 := p.WithRoot(pterm.NewTreeFromLeveledList(pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "0.0"},
		pterm.LeveledListItem{Level: 1, Text: "0.1"},
		pterm.LeveledListItem{Level: 1, Text: "0.2"},
		pterm.LeveledListItem{Level: 0, Text: "1.0"},
		pterm.LeveledListItem{Level: 0, Text: "2.0"},
		pterm.LeveledListItem{Level: 1, Text: "2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.2"},
		pterm.LeveledListItem{Level: 2, Text: "2.2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.3"}}))

	assert.Equal(t, pterm.NewTreeFromLeveledList(pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "0.0"},
		pterm.LeveledListItem{Level: 1, Text: "0.1"},
		pterm.LeveledListItem{Level: 1, Text: "0.2"},
		pterm.LeveledListItem{Level: 0, Text: "1.0"},
		pterm.LeveledListItem{Level: 0, Text: "2.0"},
		pterm.LeveledListItem{Level: 1, Text: "2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.2"},
		pterm.LeveledListItem{Level: 2, Text: "2.2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.3"}}), p2.Root)
	assert.Zero(t, p.Root)
}

func TestTreePrinter_NewTreeFromLeveledListLevelInvalidIncrease(t *testing.T) {
	p := pterm.DefaultTree
	p2 := p.WithRoot(pterm.NewTreeFromLeveledList(pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "0.0"},
		pterm.LeveledListItem{Level: 1, Text: "0.1"},
		pterm.LeveledListItem{Level: 1, Text: "0.2"},
		pterm.LeveledListItem{Level: 0, Text: "1.0"},
		pterm.LeveledListItem{Level: 0, Text: "2.0"},
		pterm.LeveledListItem{Level: 1, Text: "2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.2"},
		pterm.LeveledListItem{Level: 2, Text: "2.2.1"},
		pterm.LeveledListItem{Level: 10, Text: "2.3"}}))

	assert.Equal(t, pterm.NewTreeFromLeveledList(pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "0.0"},
		pterm.LeveledListItem{Level: 1, Text: "0.1"},
		pterm.LeveledListItem{Level: 1, Text: "0.2"},
		pterm.LeveledListItem{Level: 0, Text: "1.0"},
		pterm.LeveledListItem{Level: 0, Text: "2.0"},
		pterm.LeveledListItem{Level: 1, Text: "2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.2"},
		pterm.LeveledListItem{Level: 2, Text: "2.2.1"},
		pterm.LeveledListItem{Level: 3, Text: "2.3"}}), p2.Root)
	assert.Zero(t, p.Root)
}

func TestTreePrinter_NewTreeFromLeveledListEmptyList(t *testing.T) {
	p := pterm.DefaultTree
	p2 := p.WithRoot(pterm.NewTreeFromLeveledList(pterm.LeveledList{}))

	assert.Equal(t, pterm.NewTreeFromLeveledList(pterm.LeveledList{}), p2.Root)
	assert.Zero(t, p.Root)
}

func TestTreePrinter_NewTreeFromLeveledListNegativeLevel(t *testing.T) {
	p := pterm.DefaultTree
	p2 := p.WithRoot(pterm.NewTreeFromLeveledList(pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "0.0"},
		pterm.LeveledListItem{Level: 1, Text: "0.1"},
		pterm.LeveledListItem{Level: 1, Text: "0.2"},
		pterm.LeveledListItem{Level: 0, Text: "1.0"},
		pterm.LeveledListItem{Level: 0, Text: "2.0"},
		pterm.LeveledListItem{Level: 1, Text: "2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.2"},
		pterm.LeveledListItem{Level: 2, Text: "2.2.1"},
		pterm.LeveledListItem{Level: -5, Text: "2.3"}}))

	assert.Equal(t, pterm.NewTreeFromLeveledList(pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "0.0"},
		pterm.LeveledListItem{Level: 1, Text: "0.1"},
		pterm.LeveledListItem{Level: 1, Text: "0.2"},
		pterm.LeveledListItem{Level: 0, Text: "1.0"},
		pterm.LeveledListItem{Level: 0, Text: "2.0"},
		pterm.LeveledListItem{Level: 1, Text: "2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.2"},
		pterm.LeveledListItem{Level: 2, Text: "2.2.1"},
		pterm.LeveledListItem{Level: 0, Text: "2.3"}}), p2.Root)
	assert.Zero(t, p.Root)
}

func TestTreePrinter_WithHorizontalString(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithHorizontalString("-")

	assert.Equal(t, "-", p2.HorizontalString)
	assert.Zero(t, p.HorizontalString)
}

func TestTreePrinter_WithRoot(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithRoot(pterm.TreeNode{
		Children: nil,
		Text:     "Hello, World!",
	})

	assert.Equal(t, pterm.TreeNode{
		Children: nil,
		Text:     "Hello, World!",
	}, p2.Root)
	assert.Zero(t, p.Root)
}

func TestTreePrinter_WithTreeStyle(t *testing.T) {
	p := pterm.TreePrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithTreeStyle(s)

	assert.Equal(t, s, p2.TreeStyle)
	assert.Zero(t, p.TreeStyle)
}

func TestTreePrinter_WithTextStyle(t *testing.T) {
	p := pterm.TreePrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithTextStyle(s)

	assert.Equal(t, s, p2.TextStyle)
	assert.Zero(t, p.TextStyle)
}

func TestTreePrinter_WithTopRightCornerString(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithTopRightCornerString("-")

	assert.Equal(t, "-", p2.TopRightCornerString)
	assert.Zero(t, p.TopRightCornerString)
}

func TestTreePrinter_WithTopRightDownStringOngoing(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithTopRightDownStringOngoing("-")

	assert.Equal(t, "-", p2.TopRightDownString)
	assert.Zero(t, p.TopRightDownString)
}

func TestTreePrinter_WithVerticalString(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithVerticalString("-")

	assert.Equal(t, "-", p2.VerticalString)
	assert.Zero(t, p.VerticalString)
}

func TestTreePrinter_WithIndent(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithIndent(3)

	assert.Equal(t, 3, p2.Indent)
	assert.Zero(t, p.Indent)
}

func TestTreePrinter_WithIndentInvalid(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithIndent(0)

	assert.Equal(t, 1, p2.Indent)
	assert.Zero(t, p.Indent)
}

func TestTreePrinter_WithWriter(t *testing.T) {
	p := pterm.TreePrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	assert.Equal(t, s, p2.Writer)
	assert.Zero(t, p.Writer)
}

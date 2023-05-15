package pterm_test

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestTreePrinterNilPrint(t *testing.T) {
	pterm.TreePrinter{}.Render()
	printer := pterm.TreePrinter{}.WithRoot(pterm.NewTreeFromLeveledList(pterm.LeveledList{pterm.LeveledListItem{Text: "Hello, World!", Level: 0}}))
	content, err := printer.Srender()

	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)
}

func TestTreePrinter_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.TreePrinter{}, "WithTopRightDownStringOngoing")
}

func TestTreePrinter_Render(t *testing.T) {
	pterm.DefaultTree.WithRoot(pterm.NewTreeFromLeveledList([]pterm.LeveledListItem{
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

	testza.AssertEqual(t, pterm.NewTreeFromLeveledList(pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "0.0"},
		pterm.LeveledListItem{Level: 1, Text: "0.1"},
		pterm.LeveledListItem{Level: 1, Text: "0.2"},
		pterm.LeveledListItem{Level: 0, Text: "1.0"},
		pterm.LeveledListItem{Level: 0, Text: "2.0"},
		pterm.LeveledListItem{Level: 1, Text: "2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.2"},
		pterm.LeveledListItem{Level: 2, Text: "2.2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.3"}}), p2.Root)
	testza.AssertZero(t, p.Root)
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

	testza.AssertEqual(t, pterm.NewTreeFromLeveledList(pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "0.0"},
		pterm.LeveledListItem{Level: 1, Text: "0.1"},
		pterm.LeveledListItem{Level: 1, Text: "0.2"},
		pterm.LeveledListItem{Level: 0, Text: "1.0"},
		pterm.LeveledListItem{Level: 0, Text: "2.0"},
		pterm.LeveledListItem{Level: 1, Text: "2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.2"},
		pterm.LeveledListItem{Level: 2, Text: "2.2.1"},
		pterm.LeveledListItem{Level: 3, Text: "2.3"}}), p2.Root)
	testza.AssertZero(t, p.Root)
}

func TestTreePrinter_NewTreeFromLeveledListEmptyList(t *testing.T) {
	p := pterm.DefaultTree
	p2 := p.WithRoot(pterm.NewTreeFromLeveledList(pterm.LeveledList{}))

	testza.AssertEqual(t, pterm.NewTreeFromLeveledList(pterm.LeveledList{}), p2.Root)
	testza.AssertZero(t, p.Root)
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

	testza.AssertEqual(t, pterm.NewTreeFromLeveledList(pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "0.0"},
		pterm.LeveledListItem{Level: 1, Text: "0.1"},
		pterm.LeveledListItem{Level: 1, Text: "0.2"},
		pterm.LeveledListItem{Level: 0, Text: "1.0"},
		pterm.LeveledListItem{Level: 0, Text: "2.0"},
		pterm.LeveledListItem{Level: 1, Text: "2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.2"},
		pterm.LeveledListItem{Level: 2, Text: "2.2.1"},
		pterm.LeveledListItem{Level: 0, Text: "2.3"}}), p2.Root)
	testza.AssertZero(t, p.Root)
}

package pterm_test

import (
	"os"
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestTreePrinter_WithHorizontalString(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithHorizontalString("-")

	testza.AssertEqual(t, "-", p2.HorizontalString)
	testza.AssertZero(t, p.HorizontalString)
}

func TestTreePrinter_WithRoot(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithRoot(pterm.TreeNode{
		Children: nil,
		Text:     "Hello, World!",
	})

	testza.AssertEqual(t, pterm.TreeNode{
		Children: nil,
		Text:     "Hello, World!",
	}, p2.Root)
	testza.AssertZero(t, p.Root)
}

func TestTreePrinter_WithTreeStyle(t *testing.T) {
	p := pterm.TreePrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithTreeStyle(s)

	testza.AssertEqual(t, s, p2.TreeStyle)
	testza.AssertZero(t, p.TreeStyle)
}

func TestTreePrinter_WithTextStyle(t *testing.T) {
	p := pterm.TreePrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithTextStyle(s)

	testza.AssertEqual(t, s, p2.TextStyle)
	testza.AssertZero(t, p.TextStyle)
}

func TestTreePrinter_WithTopRightCornerString(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithTopRightCornerString("-")

	testza.AssertEqual(t, "-", p2.TopRightCornerString)
	testza.AssertZero(t, p.TopRightCornerString)
}

func TestTreePrinter_WithTopRightDownStringOngoing(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithTopRightDownStringOngoing("-")

	testza.AssertEqual(t, "-", p2.TopRightDownString)
	testza.AssertZero(t, p.TopRightDownString)
}

func TestTreePrinter_WithVerticalString(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithVerticalString("-")

	testza.AssertEqual(t, "-", p2.VerticalString)
	testza.AssertZero(t, p.VerticalString)
}

func TestTreePrinter_WithIndent(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithIndent(3)

	testza.AssertEqual(t, 3, p2.Indent)
	testza.AssertZero(t, p.Indent)
}

func TestTreePrinter_WithIndentInvalid(t *testing.T) {
	p := pterm.TreePrinter{}
	p2 := p.WithIndent(0)

	testza.AssertEqual(t, 1, p2.Indent)
	testza.AssertZero(t, p.Indent)
}

func TestTreePrinter_WithWriter(t *testing.T) {
	p := pterm.TreePrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	testza.AssertEqual(t, s, p2.Writer)
	testza.AssertZero(t, p.Writer)
}

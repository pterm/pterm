package pterm_test

import (
	"testing"

	"github.com/pterm/pterm/internal/testhelper"
	"github.com/pterm/pterm"
)

func TestBar_WithLabel(t *testing.T) {
	p := pterm.Bar{}
	s := "X"
	p2 := p.WithLabel(s)

	testhelper.AssertEqual(t, s, p2.Label)
	testhelper.AssertZero(t, p.Label)
}

func TestBar_WithStyle(t *testing.T) {
	p := pterm.Bar{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p2 := p.WithStyle(s)

	testhelper.AssertEqual(t, s, p2.Style)
	testhelper.AssertZero(t, p.Style)
}

func TestBar_WithValue(t *testing.T) {
	p := pterm.Bar{}
	s := 1337
	p2 := p.WithValue(s)

	testhelper.AssertEqual(t, s, p2.Value)
	testhelper.AssertZero(t, p.Value)
}

func TestBar_WithLabelStyle(t *testing.T) {
	p := pterm.Bar{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p2 := p.WithLabelStyle(s)

	testhelper.AssertEqual(t, s, p2.LabelStyle)
	testhelper.AssertZero(t, p.LabelStyle)
}

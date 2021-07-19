package pterm_test

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestBar_WithLabel(t *testing.T) {
	p := pterm.Bar{}
	s := "X"
	p2 := p.WithLabel(s)

	testza.AssertEqual(t, s, p2.Label)
	testza.AssertZero(t, p.Label)
}

func TestBar_WithStyle(t *testing.T) {
	p := pterm.Bar{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p2 := p.WithStyle(s)

	testza.AssertEqual(t, s, p2.Style)
	testza.AssertZero(t, p.Style)
}

func TestBar_WithValue(t *testing.T) {
	p := pterm.Bar{}
	s := 1337
	p2 := p.WithValue(s)

	testza.AssertEqual(t, s, p2.Value)
	testza.AssertZero(t, p.Value)
}

func TestBar_WithLabelStyle(t *testing.T) {
	p := pterm.Bar{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p2 := p.WithLabelStyle(s)

	testza.AssertEqual(t, s, p2.LabelStyle)
	testza.AssertZero(t, p.LabelStyle)
}

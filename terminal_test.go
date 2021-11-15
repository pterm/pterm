package pterm_test

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestSetForcedTerminalSize(t *testing.T) {
	pterm.SetForcedTerminalSize(1, 1)
	w, h, _ := pterm.GetTerminalSize()
	testza.AssertEqual(t, w, 1)
	testza.AssertEqual(t, h, 1)
	w = pterm.GetTerminalWidth()
	h = pterm.GetTerminalHeight()
	testza.AssertEqual(t, w, 1)
	testza.AssertEqual(t, h, 1)
}

package pterm_test

import (
	"os"
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
	"golang.org/x/term"
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

func TestGetTerminalSizeAutodetect(t *testing.T) {
	// enable autodetection
	pterm.SetForcedTerminalSize(0, 0)
	autoW, autoH, _ := term.GetSize(int(os.Stdout.Fd()))
	w, h, _ := pterm.GetTerminalSize()
	testza.AssertEqual(t, w, autoW)
	testza.AssertEqual(t, h, autoH)
	// disable autodetection
	pterm.SetForcedTerminalSize(terminalWidth, terminalHeight)
}

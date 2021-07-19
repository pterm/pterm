package pterm_test

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestGetTerminalHeight(t *testing.T) {
	testza.AssertNotZero(t, pterm.GetTerminalHeight())
}

func TestGetTerminalWidth(t *testing.T) {
	testza.AssertNotZero(t, pterm.GetTerminalWidth())
}

func TestGetTerminalSize(t *testing.T) {
	w, h, _ := pterm.GetTerminalSize()
	testza.AssertNotZero(t, w)
	testza.AssertNotZero(t, h)
}

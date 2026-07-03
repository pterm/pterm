package pterm_test

import (
	"os"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
	"golang.org/x/term"
)

func TestSetForcedTerminalSize(t *testing.T) {
	pterm.SetForcedTerminalSize(1, 1)

	w, h, _ := pterm.GetTerminalSize()
	assert.Equal(t, w, 1)
	assert.Equal(t, h, 1)

	w = pterm.GetTerminalWidth()
	h = pterm.GetTerminalHeight()

	assert.Equal(t, w, 1)
	assert.Equal(t, h, 1)
}

func autodetectTerminalSize() (int, int) {
	expectedW, expectedH, _ := term.GetSize(int(os.Stdout.Fd()))
	// CI goes to fallback values
	if expectedW <= 0 {
		expectedW = pterm.FallbackTerminalWidth
		expectedH = pterm.FallbackTerminalHeight
	}

	return expectedW, expectedH
}

func TestGetTerminalSizeAutodetect(t *testing.T) {
	// enable autodetection
	pterm.SetForcedTerminalSize(0, 0)

	expectedW, expectedH := autodetectTerminalSize()
	w, h, _ := pterm.GetTerminalSize()
	assert.Equal(t, expectedW, w)
	assert.Equal(t, expectedH, h)
	// disable autodetection
	pterm.SetForcedTerminalSize(terminalWidth, terminalHeight)
}

func TestGetTerminalSizeAutodetect2(t *testing.T) {
	// enable autodetection
	pterm.SetForcedTerminalSize(0, 0)

	expectedW, expectedH := autodetectTerminalSize()
	w := pterm.GetTerminalWidth()
	h := pterm.GetTerminalHeight()

	assert.Equal(t, expectedW, w)
	assert.Equal(t, expectedH, h)
	// disable autodetection
	pterm.SetForcedTerminalSize(terminalWidth, terminalHeight)
}

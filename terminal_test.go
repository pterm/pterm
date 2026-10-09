package pterm_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/term"

	"github.com/pterm/pterm"
)

// restoreForcedTerminalSize puts the terminal size back to the value TestMain
// forces for the whole suite.
func restoreForcedTerminalSize(t *testing.T) {
	t.Helper()
	t.Cleanup(func() { pterm.SetForcedTerminalSize(terminalWidth, terminalHeight) })
}

func TestSetForcedTerminalSizeIsHonored(t *testing.T) {
	restoreForcedTerminalSize(t)

	pterm.SetForcedTerminalSize(123, 45)

	w, h, err := pterm.GetTerminalSize()
	assert.NoError(t, err, "a forced size is always detectable")
	assert.Equal(t, 123, w)
	assert.Equal(t, 45, h)

	assert.Equal(t, 123, pterm.GetTerminalWidth())
	assert.Equal(t, 45, pterm.GetTerminalHeight())
}

// autodetectTerminalSize returns the size pterm should report when no size is
// forced: the real terminal size, or the fallback values when there is none
// (CI, piped output).
func autodetectTerminalSize() (int, int) {
	expectedW, expectedH, _ := term.GetSize(int(os.Stdout.Fd()))
	if expectedW <= 0 {
		expectedW = pterm.FallbackTerminalWidth
	}

	if expectedH <= 0 {
		expectedH = pterm.FallbackTerminalHeight
	}

	return expectedW, expectedH
}

func TestGetTerminalSizeAutodetection(t *testing.T) {
	restoreForcedTerminalSize(t)

	// Setting the forced size to zero re-enables autodetection.
	pterm.SetForcedTerminalSize(0, 0)

	expectedW, expectedH := autodetectTerminalSize()

	w, h, _ := pterm.GetTerminalSize()
	assert.Equal(t, expectedW, w)
	assert.Equal(t, expectedH, h)

	assert.Equal(t, expectedW, pterm.GetTerminalWidth())
	assert.Equal(t, expectedH, pterm.GetTerminalHeight())
}

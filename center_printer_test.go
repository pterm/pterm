package pterm_test

// Behavioral tests for CenterPrinter.
//
// Builder methods, Print*/Sprint* delegation and the global styling invariants
// are covered generically in contract_test.go. This file verifies the
// centering math against the 80-column terminal forced by TestMain.

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestCenterPrinterCentersSingleLine(t *testing.T) {
	// "Hello" is 5 cells wide: (80-5)/2 = 37 cells of left padding.
	assert.Equal(t, strings.Repeat(" ", 37)+"Hello\n", pterm.DefaultCenter.Sprint("Hello"))
}

func TestCenterPrinterCentersBlockAsAWhole(t *testing.T) {
	// By default the whole block is centered by its widest line (3 cells):
	// every line gets the same (80-3)/2 = 38 cells of padding, preserving the
	// relative alignment inside the block.
	expected := strings.Repeat(" ", 38) + "aaa\n" +
		strings.Repeat(" ", 38) + "a\n"

	assert.Equal(t, expected, pterm.DefaultCenter.Sprint("aaa\na"))
}

func TestCenterPrinterCentersEachLineSeparately(t *testing.T) {
	// With CenterEachLineSeparately every line gets its own padding:
	// (80-3)/2 = 38 for "aaa" and (80-1)/2 = 39 for "a". This must differ
	// from the block mode above.
	expected := strings.Repeat(" ", 38) + "aaa\n" +
		strings.Repeat(" ", 39) + "a\n"

	p := pterm.DefaultCenter.WithCenterEachLineSeparately()

	assert.Equal(t, expected, p.Sprint("aaa\na"))
	assert.NotEqual(t, pterm.DefaultCenter.Sprint("aaa\na"), p.Sprint("aaa\na"))
}

func TestCenterPrinterCentersByVisibleWidth(t *testing.T) {
	t.Run("unicode", func(t *testing.T) {
		// "汉字" occupies 4 terminal cells: (80-4)/2 = 38 cells of padding.
		assert.Equal(t, strings.Repeat(" ", 38)+"汉字\n", pterm.DefaultCenter.Sprint("汉字"))
	})

	t.Run("ANSI codes are ignored for the width", func(t *testing.T) {
		colored := pterm.FgRed.Sprint("Hello")

		// The escape codes must not count towards the width: same 37-cell
		// padding as for the plain "Hello".
		assert.Equal(t, strings.Repeat(" ", 37)+colored+"\n", pterm.DefaultCenter.Sprint(colored))
	})
}

func TestCenterPrinterLineWiderThanTerminalIsNotPadded(t *testing.T) {
	wide := strings.Repeat("a", 100)

	assert.Equal(t, wide+"\n", pterm.DefaultCenter.Sprint(wide))
	assert.Equal(t, wide+"\n", pterm.DefaultCenter.WithCenterEachLineSeparately().Sprint(wide))
}

func TestCenterPrinterRawOutputPassesTextThrough(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableStyling()

	assert.Equal(t, "Hello", pterm.DefaultCenter.Sprint("Hello"))
}

func TestCenterPrinterRespectsForcedTerminalWidth(t *testing.T) {
	pterm.SetForcedTerminalSize(20, 60)
	t.Cleanup(func() { pterm.SetForcedTerminalSize(terminalWidth, terminalHeight) })

	// (20-4)/2 = 8 cells of padding on a 20-column terminal.
	assert.Equal(t, strings.Repeat(" ", 8)+"text\n", pterm.DefaultCenter.Sprint("text"))
}

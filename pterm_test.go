package pterm_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

// EnableStyling/DisableStyling must flip both RawOutput and PrintColor
// coherently, observable through the actual output.
func TestStylingToggle(t *testing.T) {
	restoreGlobalStyling(t)

	pterm.DisableStyling()

	assert.True(t, pterm.RawOutput)
	assert.False(t, pterm.PrintColor)

	// No color codes in raw mode ...
	assert.Equal(t, "x", pterm.FgRed.Sprint("x"))
	// ... and no line-overwriting control characters either.
	assert.Equal(t, "raw", captureStdout(func(_ io.Writer) { pterm.Printo("raw") }))

	pterm.EnableStyling()

	assert.False(t, pterm.RawOutput)
	assert.True(t, pterm.PrintColor)
	assert.Equal(t, "\x1b[31mx\x1b[0m", pterm.FgRed.Sprint("x"))
	assert.Equal(t, "\rstyled", captureStdout(func(_ io.Writer) { pterm.Printo("styled") }))
}

// The debug toggle must control whether the Debug printer produces output.
func TestDebugMessagesToggle(t *testing.T) {
	t.Cleanup(pterm.DisableDebugMessages)

	pterm.DisableDebugMessages()
	assert.Empty(t, captureStdout(func(_ io.Writer) { pterm.Debug.Println("hidden") }),
		"Debug must print nothing while debug messages are disabled")

	pterm.EnableDebugMessages()

	out := captureStdout(func(_ io.Writer) { pterm.Debug.Println("visible") })
	assert.Contains(t, stripANSI(out), "visible")
}

// RecalculateTerminalSize must derive the default printer dimensions from the
// (forced) terminal size.
func TestRecalculateTerminalSize(t *testing.T) {
	t.Cleanup(func() { pterm.SetForcedTerminalSize(terminalWidth, terminalHeight) })

	// SetForcedTerminalSize recalculates internally.
	pterm.SetForcedTerminalSize(120, 30)

	assert.Equal(t, 120*2/3, pterm.DefaultBarChart.Width)
	assert.Equal(t, 30*2/3, pterm.DefaultBarChart.Height)
	assert.Equal(t, 120, pterm.DefaultParagraph.MaxWidth)

	// A direct call must restore dimensions that were changed out-of-band.
	pterm.DefaultBarChart.Width = 0
	pterm.DefaultParagraph.MaxWidth = 0

	pterm.RecalculateTerminalSize()

	assert.Equal(t, 120*2/3, pterm.DefaultBarChart.Width)
	assert.Equal(t, 120, pterm.DefaultParagraph.MaxWidth)
}

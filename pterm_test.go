package pterm_test

import (
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestDisableDebugMessages(t *testing.T) {
	pterm.EnableDebugMessages()

	pterm.DisableDebugMessages()
	assert.False(t, pterm.PrintDebugMessages)
}

func TestEnableDebugMessages(t *testing.T) {
	pterm.EnableDebugMessages()
	assert.True(t, pterm.PrintDebugMessages)
}

func TestDisableOutput(t *testing.T) {
	pterm.DisableOutput()
	assert.False(t, pterm.Output)
}

func TestEnableOutput(t *testing.T) {
	pterm.DisableOutput()
	pterm.EnableOutput()
	assert.True(t, pterm.Output)
}

func TestDisableStyling(t *testing.T) {
	pterm.EnableStyling()

	pterm.DisableStyling()
	assert.True(t, pterm.RawOutput)
}

func TestEnableStyling(t *testing.T) {
	pterm.DisableStyling()

	pterm.EnableStyling()
	assert.False(t, pterm.RawOutput)
}

func TestRecalculateTerminalSize(t *testing.T) {
	// save existing values
	prevBarChartWidth := pterm.DefaultBarChart.Width
	prevBarChartHeight := pterm.DefaultBarChart.Height
	prevParagraphMaxWidth := pterm.DefaultParagraph.MaxWidth
	w := pterm.GetTerminalWidth()
	h := pterm.GetTerminalHeight()
	// double the terminal size
	pterm.SetForcedTerminalSize(w*2, h*2)
	// assert the values doubled
	assert.Equal(t, prevBarChartWidth*2, pterm.DefaultBarChart.Width)
	assert.Equal(t, prevBarChartHeight*2, pterm.DefaultBarChart.Height)
	assert.Equal(t, prevParagraphMaxWidth*2, pterm.DefaultParagraph.MaxWidth)
	// revert the terminal size
	pterm.SetForcedTerminalSize(w, h)
}

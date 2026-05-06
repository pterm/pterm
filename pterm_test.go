package pterm_test

import (
	"testing"

	"github.com/pterm/pterm/internal/testhelper"
	"github.com/pterm/pterm"
)

func TestDisableDebugMessages(t *testing.T) {
	pterm.PrintDebugMessages = true

	pterm.DisableDebugMessages()
	testhelper.AssertFalse(t, pterm.PrintDebugMessages)
}

func TestEnableDebugMessages(t *testing.T) {
	pterm.EnableDebugMessages()
	testhelper.AssertTrue(t, pterm.PrintDebugMessages)
}

func TestDisableOutput(t *testing.T) {
	pterm.DisableOutput()
	testhelper.AssertFalse(t, pterm.Output)
}

func TestEnableOutput(t *testing.T) {
	pterm.DisableOutput()
	pterm.EnableOutput()
	testhelper.AssertTrue(t, pterm.Output)
}

func TestDisableStyling(t *testing.T) {
	pterm.RawOutput = false

	pterm.DisableStyling()
	testhelper.AssertTrue(t, pterm.RawOutput)
}

func TestEnableStyling(t *testing.T) {
	pterm.RawOutput = true

	pterm.EnableStyling()
	testhelper.AssertFalse(t, pterm.RawOutput)
}

func TestInterfaceImplementation(t *testing.T) {
	// If a printer doesn't fit into the slice, the printer doesn't has the right interface anymore.
	_ = []pterm.TextPrinter{&pterm.DefaultBasicText, pterm.DefaultBox, pterm.DefaultCenter, &pterm.DefaultHeader, &pterm.DefaultParagraph, &pterm.Info, &pterm.DefaultSection, pterm.FgRed, pterm.NewRGB(0, 0, 0)}
	_ = []pterm.LivePrinter{&pterm.DefaultProgressbar, &pterm.DefaultSpinner}
	_ = []pterm.RenderPrinter{pterm.DefaultBarChart, pterm.DefaultBigText, pterm.DefaultBulletList, pterm.DefaultPanel, pterm.DefaultTable, pterm.DefaultTree}
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
	testhelper.AssertEqual(t, prevBarChartWidth*2, pterm.DefaultBarChart.Width)
	testhelper.AssertEqual(t, prevBarChartHeight*2, pterm.DefaultBarChart.Height)
	testhelper.AssertEqual(t, prevParagraphMaxWidth*2, pterm.DefaultParagraph.MaxWidth)
	// revert the terminal size
	pterm.SetForcedTerminalSize(w, h)
}

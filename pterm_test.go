package pterm_test

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestDisableDebugMessages(t *testing.T) {
	pterm.PrintDebugMessages.Store(true)
	pterm.DisableDebugMessages()
	testza.AssertFalse(t, pterm.PrintDebugMessages.Load())
}

func TestEnableDebugMessages(t *testing.T) {
	pterm.EnableDebugMessages()
	testza.AssertTrue(t, pterm.PrintDebugMessages.Load())
}

func TestDisableOutput(t *testing.T) {
	pterm.DisableOutput()
	testza.AssertFalse(t, pterm.Output.Load())
}

func TestEnableOutput(t *testing.T) {
	pterm.DisableOutput()
	pterm.EnableOutput()
	testza.AssertTrue(t, pterm.Output.Load())
}

func TestDisableStyling(t *testing.T) {
	pterm.RawOutput.Store(false)
	pterm.DisableStyling()
	testza.AssertTrue(t, pterm.RawOutput.Load())
}

func TestEnableStyling(t *testing.T) {
	pterm.RawOutput.Store(true)
	pterm.EnableStyling()
	testza.AssertFalse(t, pterm.RawOutput.Load())
}

func TestInterfaceImplementation(t *testing.T) {
	// If a printer doesn't fit into the slice, the printer doesn't has the right interface anymore.
	_ = []pterm.TextPrinter{&pterm.DefaultBasicText, pterm.DefaultBox, pterm.DefaultCenter, &pterm.DefaultHeader, &pterm.DefaultParagraph, &pterm.Info, &pterm.DefaultSection, pterm.FgRed, pterm.NewRGB(0, 0, 0)}
	_ = []pterm.LivePrinter{pterm.DefaultProgressbar, &pterm.DefaultSpinner}
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
	testza.AssertEqual(t, prevBarChartWidth*2, pterm.DefaultBarChart.Width)
	testza.AssertEqual(t, prevBarChartHeight*2, pterm.DefaultBarChart.Height)
	testza.AssertEqual(t, prevParagraphMaxWidth*2, pterm.DefaultParagraph.MaxWidth)
	// revert the terminal size
	pterm.SetForcedTerminalSize(w, h)
}

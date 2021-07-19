package pterm_test

import (
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestDisableDebugMessages(t *testing.T) {
	pterm.PrintDebugMessages = true
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
	pterm.RawOutput = false
	pterm.DisableStyling()
	assert.True(t, pterm.RawOutput)
}

func TestEnableStyling(t *testing.T) {
	pterm.RawOutput = true
	pterm.EnableStyling()
	assert.False(t, pterm.RawOutput)
}

func TestInterfaceImplementation(t *testing.T) {
	// If a printer doesn't fit into the slice, the printer doesn't has the right interface anymore.
	_ = []pterm.TextPrinter{&pterm.DefaultBasicText, pterm.DefaultBox, pterm.DefaultCenter, &pterm.DefaultHeader, &pterm.DefaultParagraph, &pterm.Info, &pterm.DefaultSection, pterm.FgRed, pterm.NewRGB(0, 0, 0)}
	_ = []pterm.LivePrinter{pterm.DefaultProgressbar, &pterm.DefaultSpinner}
	_ = []pterm.RenderPrinter{pterm.DefaultBarChart, pterm.DefaultBigText, pterm.DefaultBulletList, pterm.DefaultPanel, pterm.DefaultTable, pterm.DefaultTree}
}

package pterm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisableDebugMessages(t *testing.T) {
	PrintDebugMessages = true
	DisableDebugMessages()
	assert.False(t, PrintDebugMessages)
}

func TestEnableDebugMessages(t *testing.T) {
	EnableDebugMessages()
	assert.True(t, PrintDebugMessages)
}

func TestDisableOutput(t *testing.T) {
	DisableOutput()
	assert.False(t, Output)
}

func TestEnableOutput(t *testing.T) {
	DisableOutput()
	EnableOutput()
	assert.True(t, Output)
}

func TestDisableStyling(t *testing.T) {
	RawOutput = false
	DisableStyling()
	assert.True(t, RawOutput)
}

func TestEnableStyling(t *testing.T) {
	RawOutput = true
	EnableStyling()
	assert.False(t, RawOutput)
}

func TestInterfaceImplementation(t *testing.T) {
	// If a printer doesn't fit into the slice, the printer doesn't has the right interface anymore.
	_ = []TextPrinter{&DefaultBasicText, DefaultBox, DefaultCenter, &DefaultHeader, &DefaultParagraph, &Info, &DefaultSection, FgRed, NewRGB(0, 0, 0)}
	_ = []LivePrinter{DefaultProgressbar, &DefaultSpinner}
	_ = []RenderPrinter{DefaultBarChart, DefaultBigText, DefaultBulletList, DefaultPanel, DefaultTable, DefaultTree}
}

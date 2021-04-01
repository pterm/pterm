package pterm

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestInterfaceImplementation(t *testing.T) {
	// If a printer doesn't fit into the slice, the printer doesn't has the right interface anymore.
	_ = []TextPrinter{&DefaultBasicText, DefaultBox, DefaultCenter, &DefaultHeader, &DefaultParagraph, &Info, &DefaultSection}
	_ = []LivePrinter{DefaultProgressbar, &DefaultSpinner}
	_ = []RenderPrinter{DefaultBarChart, DefaultBigText, DefaultBulletList, DefaultPanel, DefaultTable, DefaultTree}
}

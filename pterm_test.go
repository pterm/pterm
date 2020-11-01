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

func TestEnableOutput(t *testing.T) {
	DisableOutput()
	EnableOutput()
	assert.True(t, Output)
}

func TestDisableOutput(t *testing.T) {
	DisableOutput()
	assert.False(t, Output)
}

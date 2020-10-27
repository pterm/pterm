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

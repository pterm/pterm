package pterm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTerminalHeight(t *testing.T) {
	assert.NotEmpty(t, GetTerminalHeight())
}

func TestGetTerminalWidth(t *testing.T) {
	assert.NotEmpty(t, GetTerminalWidth())
}

func TestGetTerminalSize(t *testing.T) {
	w, h, _ := GetTerminalSize()
	assert.NotEmpty(t, w)
	assert.NotEmpty(t, h)
}

package pterm_test

import (
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestGetTerminalHeight(t *testing.T) {
	assert.NotEmpty(t, pterm.GetTerminalHeight())
}

func TestGetTerminalWidth(t *testing.T) {
	assert.NotEmpty(t, pterm.GetTerminalWidth())
}

func TestGetTerminalSize(t *testing.T) {
	w, h, _ := pterm.GetTerminalSize()
	assert.NotEmpty(t, w)
	assert.NotEmpty(t, h)
}

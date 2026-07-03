package color_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal/color"
)

func TestSequence(t *testing.T) {
	assert.Equal(t, "\x1b[31m", color.Sequence("31"))
	assert.Equal(t, "\x1b[31;45;1m", color.Sequence("31;45;1"))
}

func TestForegroundRGB(t *testing.T) {
	assert.Equal(t, "\x1b[38;2;1;2;3m", color.ForegroundRGB(1, 2, 3))
	assert.Equal(t, "\x1b[38;2;255;255;255m", color.ForegroundRGB(255, 255, 255))
}

func TestBackgroundRGB(t *testing.T) {
	assert.Equal(t, "\x1b[48;2;1;2;3m", color.BackgroundRGB(1, 2, 3))
	assert.Equal(t, "\x1b[48;2;0;0;0m", color.BackgroundRGB(0, 0, 0))
}

func TestWrap(t *testing.T) {
	assert.Equal(t, "\x1b[31mhi\x1b[0m", color.Wrap("31", "hi"))
}

func TestWrapEmptyInputStaysEmpty(t *testing.T) {
	assert.Equal(t, "", color.Wrap("31", ""))
	assert.Equal(t, "hi", color.Wrap("", "hi"))
}

func TestStrip(t *testing.T) {
	assert.Equal(t, "hi", color.Strip("\x1b[31mhi\x1b[0m"))
	assert.Equal(t, "plain", color.Strip("plain"))
	assert.Equal(t, "ab", color.Strip("\x1b[38;2;1;2;3ma\x1b[0m\x1b[1mb\x1b[0m"))
}

func TestStripKeepsNonSGRSequences(t *testing.T) {
	// Cursor movement is not an SGR sequence and must survive.
	assert.Equal(t, "\x1b[2Ahi", color.Strip("\x1b[2A\x1b[31mhi\x1b[0m"))
}

func TestWrapStripRoundtrip(t *testing.T) {
	assert.Equal(t, "hello", color.Strip(color.Wrap("32;1", "hello")))
}

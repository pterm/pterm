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

func TestForeground256(t *testing.T) {
	assert.Equal(t, "\x1b[38;5;196m", color.Foreground256(196))
	assert.Equal(t, "\x1b[38;5;0m", color.Foreground256(0))
}

func TestBackground256(t *testing.T) {
	assert.Equal(t, "\x1b[48;5;196m", color.Background256(196))
	assert.Equal(t, "\x1b[48;5;255m", color.Background256(255))
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
	assert.Equal(t, "", color.Strip(""))
	assert.Equal(t, "x", color.Strip(color.Foreground256(196)+"x"+color.Reset))
	assert.Equal(t, "x", color.Strip(color.Background256(196)+"x"+color.Reset))
	assert.Equal(t, "x", color.Strip(color.BackgroundRGB(1, 2, 3)+"x"+color.Reset))
	assert.Equal(t, "reset only", color.Strip("\x1b[mreset only")) // bare reset sequence without a code
}

func TestStripEveryLevelSequence(t *testing.T) {
	// Whatever level a color was downsampled to, Strip must remove it again.
	for _, level := range []color.Level{color.LevelNone, color.LevelBasic, color.Level256, color.LevelTrueColor} {
		assert.Equal(t, "x", color.Strip(level.ForegroundRGB(12, 34, 56)+"x"), "level %s", level)
		assert.Equal(t, "x", color.Strip(level.BackgroundRGB(12, 34, 56)+"x"), "level %s", level)
	}
}

func TestStripKeepsNonSGRSequences(t *testing.T) {
	// Cursor movement is not an SGR sequence and must survive.
	assert.Equal(t, "\x1b[2Ahi", color.Strip("\x1b[2A\x1b[31mhi\x1b[0m"))
}

func TestWrapStripRoundtrip(t *testing.T) {
	assert.Equal(t, "hello", color.Strip(color.Wrap("32;1", "hello")))
}

// The sequence builders guarantee color output: they are pure functions that
// always emit escape sequences, no matter what the environment says. Whether
// colors should be printed at all is the caller's decision (pterm gates on
// pterm.PrintColor), so an environment opting out of color must not silently
// disable them here.
func TestSequenceBuildersIgnoreEnvironment(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	t.Setenv("TERM", "dumb")
	t.Setenv("FORCE_COLOR", "0")
	t.Setenv("COLORTERM", "")

	assert.Equal(t, "\x1b[31m", color.Sequence("31"))
	assert.Equal(t, "\x1b[31mhi\x1b[0m", color.Wrap("31", "hi"))
	assert.Equal(t, "\x1b[38;2;1;2;3m", color.ForegroundRGB(1, 2, 3))
	assert.Equal(t, "\x1b[48;2;1;2;3m", color.BackgroundRGB(1, 2, 3))
	assert.Equal(t, "\x1b[38;5;196m", color.Foreground256(196))
	assert.Equal(t, "\x1b[48;5;196m", color.Background256(196))
	assert.Equal(t, "\x1b[38;2;1;2;3m", color.LevelTrueColor.ForegroundRGB(1, 2, 3))
}

func TestDetectLevelHonorsOptOut(t *testing.T) {
	t.Setenv("FORCE_COLOR", "")
	t.Setenv("CLICOLOR_FORCE", "")
	t.Setenv("NO_COLOR", "1")

	assert.Equal(t, color.LevelNone, color.DetectLevel())
	assert.False(t, color.SupportsANSI())
	assert.False(t, color.SupportsTrueColor())
}

func TestDetectLevelHonorsTrueColorEnvironment(t *testing.T) {
	t.Setenv("FORCE_COLOR", "")
	t.Setenv("CLICOLOR_FORCE", "")
	t.Setenv("NO_COLOR", "")
	t.Setenv("CLICOLOR", "")
	t.Setenv("COLORTERM", "truecolor")

	assert.Equal(t, color.LevelTrueColor, color.DetectLevel())
	assert.True(t, color.SupportsANSI())
	assert.True(t, color.SupportsTrueColor())
}

package pterm

import (
	"os"

	"go.uber.org/atomic"
	"golang.org/x/term"
)

// FallbackTerminalWidth is the value used for GetTerminalWidth, if the actual width can not be detected
// You can override that value if necessary.
var FallbackTerminalWidth = 80

// FallbackTerminalHeight is the value used for GetTerminalHeight, if the actual height can not be detected
// You can override that value if necessary.
var FallbackTerminalHeight = 10

// forcedTerminalWidth, when set along with forcedTerminalHeight, forces the terminal width value.
var forcedTerminalWidth *atomic.Int64 = atomic.NewInt64(0)

// forcedTerminalHeight, when set along with forcedTerminalWidth, forces the terminal height value.
var forcedTerminalHeight *atomic.Int64 = atomic.NewInt64(0)

// GetTerminalWidth returns the terminal width of the active terminal.
func GetTerminalWidth() int {
	if forcedTerminalWidth.Load() > 0 {
		return int(forcedTerminalWidth.Load())
	}
	width, _, _ := GetTerminalSize()
	return width
}

// GetTerminalHeight returns the terminal height of the active terminal.
func GetTerminalHeight() int {
	if forcedTerminalHeight.Load() > 0 {
		return int(forcedTerminalHeight.Load())
	}
	_, height, _ := GetTerminalSize()
	return height
}

// GetTerminalSize returns the width and the height of the active terminal.
func GetTerminalSize() (width, height int, err error) {
	if forcedTerminalWidth.Load() > 0 && forcedTerminalHeight.Load() > 0 {
		return int(forcedTerminalWidth.Load()), int(forcedTerminalHeight.Load()), nil
	}
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if w <= 0 {
		w = FallbackTerminalWidth
	}
	if h <= 0 {
		h = FallbackTerminalHeight
	}
	if err != nil {
		err = ErrTerminalSizeNotDetectable
	}
	return w, h, err
}

// setForcedTerminalSize turns off terminal size autodetection. Usuful for unified tests.
func SetForcedTerminalSize(width int, height int) {
	forcedTerminalWidth.Store(int64(width))
	forcedTerminalHeight.Store(int64(height))
	RecalculateTerminalSize()
}

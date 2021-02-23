package pterm

import (
	"os"

	"golang.org/x/term"
)

// FallbackTerminalWidth is the value used for GetTerminalWidth, if the actual width can not be detected
// You can override that value if necessary.
var FallbackTerminalWidth = 80

// FallbackTerminalHeight is the value used for GetTerminalHeight, if the actual height can not be detected
// You can override that value if necessary.
var FallbackTerminalHeight = 10

// GetTerminalWidth returns the terminal width of the active terminal.
func GetTerminalWidth() int {
	width, _, _ := GetTerminalSize()
	return width
}

// GetTerminalHeight returns the terminal height of the active terminal.
func GetTerminalHeight() int {
	_, height, _ := GetTerminalSize()
	return height
}

// GetTerminalSize returns the width and the height of the active terminal.
func GetTerminalSize() (width, height int, err error) {
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

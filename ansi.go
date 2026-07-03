package pterm

import (
	"github.com/pterm/pterm/internal/color"
)

// resetSequence is the ANSI escape sequence that resets all styling.
const resetSequence = color.Reset

// renderCode wraps s in the ANSI escape sequence for the given SGR code
// (e.g. "31" or "31;45;1"). When colors are disabled it returns s with any
// embedded escape codes stripped instead, so disabling colors also cleans
// strings that already contain styling.
func renderCode(code string, s string) string {
	if !printColorEnabled() {
		return color.Strip(s)
	}

	return color.Wrap(code, s)
}

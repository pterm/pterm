// Package color is pterm's own ANSI coloring engine.
//
// The sequence builders (Sequence, Wrap, ForegroundRGB, ...) are pure
// functions that always emit escape sequences: whether colors should be
// printed at all is decided by the caller (pterm gates on pterm.PrintColor).
// On top of that the package detects how many colors the current terminal can
// render (DetectLevel) so callers can degrade RGB colors gracefully, and on
// Windows it switches the console into virtual terminal mode so ANSI colors
// also work in classic consoles like cmd.exe — the only state the package
// keeps is the result of that one-time console handshake.
//
// The package is internal for now — it is an experiment in giving pterm full
// control over its color rendering. If it proves itself, it may be promoted
// to a public package later.
package color

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	esc = "\x1b"

	// Reset is the ANSI escape sequence that resets all active styling.
	Reset = esc + "[0m"

	// ClearToEOL is the ANSI escape sequence that clears from the cursor to
	// the end of the line. Used after background colors so the color fills
	// the rest of the row.
	ClearToEOL = esc + "[K"
)

// sgrRegex matches ANSI SGR (color and style) escape sequences.
var sgrRegex = regexp.MustCompile(`\x1b\[[\d;]*m`)

// Sequence returns the ANSI SGR escape sequence for the given code.
// The code may contain multiple attributes separated by semicolons:
//
//	Sequence("31")   == "\x1b[31m"
//	Sequence("31;1") == "\x1b[31;1m"
func Sequence(code string) string {
	return esc + "[" + code + "m"
}

// ForegroundRGB returns the ANSI escape sequence that sets the foreground to
// the given 24-bit RGB color.
func ForegroundRGB(r, g, b uint8) string {
	return rgbSequence("38", r, g, b)
}

// BackgroundRGB returns the ANSI escape sequence that sets the background to
// the given 24-bit RGB color.
func BackgroundRGB(r, g, b uint8) string {
	return rgbSequence("48", r, g, b)
}

// Foreground256 returns the ANSI escape sequence that sets the foreground to
// the given xterm 256-color palette index.
func Foreground256(n uint8) string {
	return esc + "[38;5;" + strconv.Itoa(int(n)) + "m"
}

// Background256 returns the ANSI escape sequence that sets the background to
// the given xterm 256-color palette index.
func Background256(n uint8) string {
	return esc + "[48;5;" + strconv.Itoa(int(n)) + "m"
}

func rgbSequence(layer string, r, g, b uint8) string {
	var sb strings.Builder
	sb.WriteString(esc + "[")
	sb.WriteString(layer)
	sb.WriteString(";2;")
	sb.WriteString(strconv.Itoa(int(r)))
	sb.WriteByte(';')
	sb.WriteString(strconv.Itoa(int(g)))
	sb.WriteByte(';')
	sb.WriteString(strconv.Itoa(int(b)))
	sb.WriteByte('m')

	return sb.String()
}

// Wrap wraps s in the SGR sequence for code and a trailing reset:
//
//	Wrap("31", "hi") == "\x1b[31mhi\x1b[0m"
//
// Empty input is returned unchanged so no stray sequences are emitted.
func Wrap(code, s string) string {
	if s == "" || code == "" {
		return s
	}

	return Sequence(code) + s + Reset
}

// Strip removes all ANSI SGR (color and style) escape sequences from s.
func Strip(s string) string {
	return sgrRegex.ReplaceAllString(s, "")
}

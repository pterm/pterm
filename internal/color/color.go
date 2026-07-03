// Package color is pterm's own ANSI coloring engine.
//
// It is deliberately stateless: every function is pure, so the package is
// safe for concurrent use without any locking. Whether colors should be
// emitted at all is decided by the caller (pterm gates on pterm.PrintColor);
// this package only knows how to build and remove escape sequences.
//
// The package is internal for now — it is an experiment in giving pterm full
// control over its color rendering. If it proves itself, it may be promoted
// to a public package later.
package color

import (
	"os"
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

// SupportsTrueColor reports whether the current terminal advertises 24-bit
// (RGB) color support.
func SupportsTrueColor() bool {
	colorTerm := strings.ToLower(os.Getenv("COLORTERM"))
	if strings.Contains(colorTerm, "truecolor") || strings.Contains(colorTerm, "24bit") {
		return true
	}

	term := strings.ToLower(os.Getenv("TERM"))
	if strings.Contains(term, "24bit") || strings.Contains(term, "truecolor") {
		return true
	}

	// Windows Terminal and common GUI terminals support true color but do not
	// always set COLORTERM.
	if os.Getenv("WT_SESSION") != "" || os.Getenv("TERMINAL_EMULATOR") == "JetBrains-JediTerm" {
		return true
	}

	switch os.Getenv("TERM_PROGRAM") {
	case "iTerm.app", "Hyper", "vscode", "ghostty", "WezTerm", "Apple_Terminal":
		return true
	}

	return false
}

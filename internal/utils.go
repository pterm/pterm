package internal

import (
	"os"
	"regexp"

	"github.com/gookit/color"
)

var osc8Regex = regexp.MustCompile(`\x1b\]8;;[^\x07]*?(?:\x07|\x1b\\)(.*?)\x1b\]8;;(?:\x07|\x1b\\)`)

// RunsInCi returns true if the current build is running on a CI server.
func RunsInCi() bool {
	return os.Getenv("CI") != ""
}

// StripOSC8Hyperlinks removes OSC 8 hyperlink sequences and keeps only the visible link text.
// OSC 8 format: \033]8;;URL\033\\LINK_TEXT\033]8;;\033\\
// Only LINK_TEXT is visible in the terminal, so we keep only that part.
func StripOSC8Hyperlinks(s string) string {
	// Replace OSC 8 hyperlinks with just the link text (capture group 1)
	return osc8Regex.ReplaceAllString(s, "$1")
}

// RemoveEscapeCodes removes both OSC 8 hyperlinks and ANSI color codes from a string.
// This should be used when calculating the visible width of a string for rendering.
func RemoveEscapeCodes(s string) string {
	// Strip OSC 8 hyperlinks first, then color codes
	s = StripOSC8Hyperlinks(s)
	return color.ClearCode(s)
}

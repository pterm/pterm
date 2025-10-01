package internal

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

// GetStringMaxWidth returns the maximum width of a string with multiple lines.
func GetStringMaxWidth(s string) int {
	var maxString int
	ss := strings.Split(s, "\n")
	for _, s2 := range ss {
		// Strip OSC 8 hyperlinks and color codes
		s2WithoutEscapes := RemoveEscapeCodes(s2)
		if runewidth.StringWidth(s2WithoutEscapes) > maxString {
			maxString = runewidth.StringWidth(s2WithoutEscapes)
		}
	}
	return maxString
}

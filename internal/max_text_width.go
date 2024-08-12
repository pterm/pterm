package internal

import (
	"strings"

	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
)

// GetStringMaxWidth returns the maximum width of a string with multiple lines.
func GetStringMaxWidth(s string) int {
	var maxString int
	ss := strings.Split(s, "\n")
	for _, s2 := range ss {
		s2WithoutColor := color.ClearCode(s2)
		if runewidth.StringWidth(s2WithoutColor) > maxString {
			maxString = runewidth.StringWidth(s2WithoutColor)
		}
	}
	return maxString
}

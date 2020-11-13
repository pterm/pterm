package internal

import (
	"strings"

	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
)

// GetStringMaxWidth returns the maximum width of a string with multiple lines.
func GetStringMaxWidth(s string) int {
	var max int
	ss := strings.Split(s, "\n")
	for _, s2 := range ss {
		if runewidth.StringWidth(color.ClearCode(s2)) > max {
			max = runewidth.StringWidth(color.ClearCode(s2))
		}
	}
	return max
}

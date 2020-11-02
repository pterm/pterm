package internal

import (
	"github.com/mattn/go-runewidth"
	"strings"
)

// GetStringMaxWidth returns the maximum width of a string with multiple lines.
func GetStringMaxWidth(s string) int {
	var max int
	ss := strings.Split(s, "\n")
	for _, s2 := range ss {
		if runewidth.StringWidth(s2) > max {
			max = runewidth.StringWidth(s2)
		}
	}
	return max
}

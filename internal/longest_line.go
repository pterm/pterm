package internal

import (
	"strings"

	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
)

// ReturnLongestLine returns the longest line with a given separator
func ReturnLongestLine(text, sep string) string {
	lines := strings.Split(text, sep)
	var longest string
	for _, line := range lines {
		if runewidth.StringWidth(color.ClearCode(line)) > runewidth.StringWidth(color.ClearCode(longest)) {
			longest = line
		}
	}

	return longest
}

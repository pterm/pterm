package internal

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

// ReturnLongestLine returns the longest line with a given separator
func ReturnLongestLine(text, sep string) string {
	lines := strings.Split(text, sep)
	var longest string
	for _, line := range lines {
		if runewidth.StringWidth(RemoveEscapeCodes(line)) > runewidth.StringWidth(RemoveEscapeCodes(longest)) {
			longest = line
		}
	}

	return longest
}

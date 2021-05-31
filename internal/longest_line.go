package internal

import (
	"strings"

	"github.com/gookit/color"
)

// ReturnLongestLine returns the longest line with a given separator
func ReturnLongestLine(text, sep string) string {
	lines := strings.Split(text, sep)
	var longest string
	for _, line := range lines {
		if len(color.ClearCode(line)) > len(color.ClearCode(longest)) {
			longest = line
		}
	}

	return longest
}

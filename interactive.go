package pterm

import (
	"strings"
	"unicode/utf8"

	"github.com/pterm/pterm/internal"
)

func getMaxW(s string) int {
	return internal.GetStringMaxWidth(s)
}

// fit the text for terminal width,offset to the right
func textFitWidth(text string, offset ...int) string {
	return strings.Join(linesFitWidth(strings.Split(text, "\n"), offset...), "\n")
}

// fit the []string for terminal width,offset to the right
func linesFitWidth(ss []string, offset ...int) []string {
	w := max(GetTerminalWidth()-1, 1)
	firstLineOffset := 0
	if len(offset) > 0 {
		firstLineOffset = offset[0]
	}
	if getMaxW(strings.Join(ss, "\n")) >= w || getMaxW(ss[0])+firstLineOffset >= w {
		// find the last index that GetStringWidth(s[:index])<=width
		findIndex := func(s string, width int) int {
			l, r := 0, len(s)+1
			for l+1 < r {
				mid := (l + r) >> 1
				if getMaxW(s[:mid]) <= width {
					l = mid
				} else {
					r = mid
				}
			}
			return l
		}

		buffer := make([]string, 0)
		for k, s := range ss {
			if k == 0 {
				fw := w - firstLineOffset
				if getMaxW(s) <= fw {
					buffer = append(buffer, s)
					continue
				}
				// Split First Row
				i := indexSplitAtByte(s, findIndex(s, fw))
				buffer = append(buffer, s[:i])
				s = s[i:]
			}
			if getMaxW(s) <= w {
				buffer = append(buffer, s)
				continue
			}
			for len(s) > 0 {
				i := indexSplitAtByte(s, findIndex(s, w))
				buffer = append(buffer, s[:i])
				s = s[i:]
			}
		}
		return buffer
	}

	return ss
}

// Returns a corrected index that,
// when a string is split into two strings, puts the broken character into the next string
func indexSplitAtByte(s string, i int) int {
	if i < 0 {
		i = 0
	}
	if i > len(s) {
		i = len(s)
	}
	if i == 0 || i == len(s) || utf8.RuneStart(s[i]) {
		return i
	}
	// Go back to the start byte of the current character
	for i > 0 && !utf8.RuneStart(s[i]) {
		i--
	}
	return i
}

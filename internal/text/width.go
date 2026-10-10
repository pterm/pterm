// Package text measures and shapes terminal text that may contain escape
// sequences: visible width, wrapping, truncation and padding.
package text

import (
	"strings"
	"unicode/utf8"

	"atomicgo.dev/color"
	"github.com/mattn/go-runewidth"
)

// Width returns the number of terminal columns s takes on one line. Escape
// sequences take no columns and wide characters take two.
func Width(s string) int {
	return runewidth.StringWidth(color.Strip(s))
}

// MaxWidth returns the width of the widest line of s.
func MaxWidth(s string) int {
	widest := 0
	for line := range strings.SplitSeq(s, "\n") {
		widest = max(widest, Width(line))
	}

	return widest
}

// PadRight appends spaces until s is width columns wide.
func PadRight(s string, width int) string {
	return s + strings.Repeat(" ", max(0, width-Width(s)))
}

// PadLeft prepends spaces until s is width columns wide.
func PadLeft(s string, width int) string {
	return strings.Repeat(" ", max(0, width-Width(s))) + s
}

// Center pads s on both sides until it is width columns wide. An odd
// remainder goes to the right.
func Center(s string, width int) string {
	gap := max(0, width-Width(s))
	return strings.Repeat(" ", gap/2) + s + strings.Repeat(" ", gap-gap/2)
}

// Truncate cuts s to at most width columns and ends it with an ellipsis when
// it had to cut. Escape sequences are kept intact, and a reset follows the
// ellipsis so a cut style does not leak. A width of zero or less returns s
// unchanged.
func Truncate(s string, width int, ellipsis string) string {
	if width <= 0 || Width(s) <= width {
		return s
	}

	budget := width - Width(ellipsis)

	var (
		out     strings.Builder
		used    int
		escaped bool
	)

	for i := 0; i < len(s); {
		if seq := escapeAt(s, i); seq != "" {
			out.WriteString(seq)

			escaped = true
			i += len(seq)

			continue
		}

		r, size := utf8.DecodeRuneInString(s[i:])
		w := runewidth.RuneWidth(r)

		if used+w > budget {
			break
		}

		out.WriteString(s[i : i+size])

		used += w
		i += size
	}

	out.WriteString(ellipsis)

	if escaped {
		out.WriteString(color.Reset)
	}

	return out.String()
}

// escapeAt returns the escape sequence that starts at s[i], or "".
func escapeAt(s string, i int) string {
	if s[i] != '\x1b' || i+1 >= len(s) {
		return ""
	}

	switch s[i+1] {
	case '[':
		for j := i + 2; j < len(s); j++ {
			if s[j] >= 0x40 && s[j] <= 0x7e {
				return s[i : j+1]
			}
		}
	case ']':
		for j := i + 2; j < len(s); j++ {
			if s[j] == '\a' {
				return s[i : j+1]
			}

			if s[j] == '\x1b' && j+1 < len(s) && s[j+1] == '\\' {
				return s[i : j+2]
			}
		}
	}

	return ""
}

// Escapes returns the start and end offsets of every escape sequence in s.
func Escapes(s string) [][2]int {
	var locs [][2]int

	for i := 0; i < len(s); i++ {
		if seq := escapeAt(s, i); seq != "" {
			locs = append(locs, [2]int{i, i + len(seq)})
			i += len(seq) - 1
		}
	}

	return locs
}

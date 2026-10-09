package internal

import (
	"regexp"
	"strings"

	"github.com/mattn/go-runewidth"
)

// WrapText wraps text so that no line is visibly wider than width, preserving
// explicit newlines. Lines break at spaces; a single word wider than a whole
// line is broken at rune boundaries. Words that contain escape sequences are
// never broken apart, so that the sequences stay intact. A width of zero or
// less disables wrapping and only splits at existing newlines.
func WrapText(text string, width int) []string {
	text = strings.ReplaceAll(text, "\r\n", "\n")

	var lines []string

	for line := range strings.SplitSeq(text, "\n") {
		lines = append(lines, wrapLine(line, width)...)
	}

	return lines
}

// wrapLine greedily wraps a single line (no newlines) at spaces.
func wrapLine(line string, width int) []string {
	if width <= 0 || GetStringMaxWidth(line) <= width {
		return []string{line}
	}

	var words []string

	for word := range strings.SplitSeq(line, " ") {
		if GetStringMaxWidth(word) > width {
			words = append(words, breakWord(word, width)...)
		} else {
			words = append(words, word)
		}
	}

	var (
		lines        []string
		current      strings.Builder
		currentWidth int
		hasWord      bool
	)

	for _, word := range words {
		wordWidth := GetStringMaxWidth(word)

		if hasWord && currentWidth+1+wordWidth > width {
			lines = append(lines, strings.TrimRight(current.String(), " "))
			current.Reset()

			currentWidth = 0
			hasWord = false
		}

		if hasWord {
			current.WriteString(" ")

			currentWidth++
		}

		current.WriteString(word)

		currentWidth += wordWidth
		hasWord = true
	}

	return append(lines, strings.TrimRight(current.String(), " "))
}

// sgrEdgeRegex matches SGR escape sequences at the start or end of a string.
var sgrEdgeRegex = regexp.MustCompile(`^(?:\x1b\[[0-9;]*m)+|(?:\x1b\[[0-9;]*m)+$`)

// breakWord splits a single word into chunks no wider than width. Styling that
// wraps the whole word is re-applied to every chunk; words with escape
// sequences in the middle are returned unbroken, because splitting them could
// cut an escape sequence in half.
func breakWord(word string, width int) []string {
	var prefix, suffix string

	content := sgrEdgeRegex.ReplaceAllStringFunc(word, func(match string) string {
		if strings.HasPrefix(word, match) {
			prefix = match
		} else {
			suffix = match
		}

		return ""
	})

	if strings.ContainsRune(content, '\x1b') {
		return []string{word}
	}

	word = content

	var (
		parts     []string
		part      strings.Builder
		partWidth int
	)

	for _, r := range word {
		runeWidth := runewidth.RuneWidth(r)

		if partWidth > 0 && partWidth+runeWidth > width {
			parts = append(parts, part.String())
			part.Reset()

			partWidth = 0
		}

		part.WriteRune(r)

		partWidth += runeWidth
	}

	if part.Len() > 0 {
		parts = append(parts, part.String())
	}

	if prefix != "" || suffix != "" {
		for i := range parts {
			parts[i] = prefix + parts[i] + suffix
		}
	}

	return parts
}

package pterm_test

// Behavioral tests for ParagraphPrinter.
//
// Builder methods, Print*/Sprint* delegation and the global styling invariants
// are covered generically in contract_test.go. This file verifies the word
// wrapping semantics: lines never exceed MaxWidth, words are never split, and
// whitespace is normalized.

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestParagraphPrinterWrapsAtMaxWidth(t *testing.T) {
	p := pterm.DefaultParagraph.WithMaxWidth(10)

	// "aaa bb ccc" fills exactly 10 columns; "dddd" no longer fits and starts
	// the next line.
	assert.Equal(t, "aaa bb ccc\ndddd e", p.Sprint("aaa bb ccc dddd e"))
}

func TestParagraphPrinterWordFillingTheWidthExactly(t *testing.T) {
	p := pterm.DefaultParagraph.WithMaxWidth(7)

	assert.Equal(t, "abc def\nghi", p.Sprint("abc def ghi"))
}

func TestParagraphPrinterNeverExceedsMaxWidth(t *testing.T) {
	const width = 12

	input := "the quick brown fox jumps over the lazy dog again and again"
	out := pterm.DefaultParagraph.WithMaxWidth(width).Sprint(input)

	for i, line := range strings.Split(out, "\n") {
		assert.LessOrEqualf(t, len(line), width, "line %d %q exceeds MaxWidth", i, line)
		assert.NotEmptyf(t, line, "line %d must not be empty", i)
		assert.Equalf(t, strings.TrimSpace(line), line, "line %d must not have leading/trailing spaces", i)
	}

	// Wrapping must preserve every word in order.
	assert.Equal(t, strings.Fields(input), strings.Fields(strings.ReplaceAll(out, "\n", " ")))
}

func TestParagraphPrinterDoesNotSplitOverlongWords(t *testing.T) {
	p := pterm.DefaultParagraph.WithMaxWidth(5)

	// A single word longer than MaxWidth is kept intact on its own line.
	assert.Equal(t, "abcdefghij\nk", p.Sprint("abcdefghij k"))
}

func TestParagraphPrinterNormalizesWhitespace(t *testing.T) {
	p := pterm.DefaultParagraph.WithMaxWidth(80)

	// Runs of spaces, tabs and newlines all collapse into single spaces.
	assert.Equal(t, "a b c d", p.Sprint("a  b\nc\td"))
}

func TestParagraphPrinterEmptyInput(t *testing.T) {
	assert.Equal(t, "", pterm.DefaultParagraph.Sprint(""))
	assert.Equal(t, "", pterm.DefaultParagraph.Sprint("   \n\t "))
}

func TestParagraphPrinterSprintlnAppendsNewline(t *testing.T) {
	assert.Equal(t, "aaa bb\n", pterm.DefaultParagraph.WithMaxWidth(10).Sprintln("aaa bb"))
}

func TestParagraphPrinterRawOutputDoesNotWrap(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableStyling()

	input := "this text would normally be wrapped"

	assert.Equal(t, input, pterm.DefaultParagraph.WithMaxWidth(10).Sprint(input))
}

func TestParagraphPrinterDefaultMaxWidthIsTerminalWidth(t *testing.T) {
	// TestMain forces an 80-column terminal; RecalculateTerminalSize keeps the
	// default paragraph printer in sync with it.
	assert.Equal(t, terminalWidth, pterm.DefaultParagraph.MaxWidth)
}

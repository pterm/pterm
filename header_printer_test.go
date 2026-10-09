package pterm_test

// Behavioral tests for HeaderPrinter.
//
// Builder methods, Print*/Sprint* delegation and the global styling invariants
// are covered generically in contract_test.go; the styled default output is
// locked by snapshot tests. This file verifies the layout math: margins,
// centering, full width, multiline padding and terminal-width clamping.
//
// TestMain forces the terminal size to 80x60.

import (
	"strings"
	"testing"

	"github.com/mattn/go-runewidth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// headerLines strips ANSI codes and splits the header block into its lines
// (dropping the trailing newline).
func headerLines(t *testing.T, out string) []string {
	t.Helper()

	plain := stripANSI(out)
	require.True(t, strings.HasSuffix(plain, "\n"), "header output must end with a newline")

	return strings.Split(strings.TrimSuffix(plain, "\n"), "\n")
}

func TestHeaderPrinterCentersTextWithinMargin(t *testing.T) {
	// Margin 5 on each side of "Hello" (5 cells) -> 15-wide header with a
	// blank padded line above and below.
	expected := "               \n" +
		"     Hello     \n" +
		"               \n"

	assert.Equal(t, expected, stripANSI(pterm.DefaultHeader.Sprint("Hello")))
}

func TestHeaderPrinterWithMargin(t *testing.T) {
	expected := "         \n" +
		"  Hello  \n" +
		"         \n"

	assert.Equal(t, expected, stripANSI(pterm.DefaultHeader.WithMargin(2).Sprint("Hello")))
}

func TestHeaderPrinterZeroMarginShrinksToText(t *testing.T) {
	p := pterm.HeaderPrinter{}

	expected := "     \n" +
		"Hello\n" +
		"     \n"

	assert.Equal(t, expected, stripANSI(p.Sprint("Hello")))
}

func TestHeaderPrinterFullWidthSpansTerminal(t *testing.T) {
	lines := headerLines(t, pterm.DefaultHeader.WithFullWidth().Sprint("Hello"))
	require.Len(t, lines, 3)

	for i, line := range lines {
		assert.Lenf(t, line, 80, "line %d must span the forced 80-column terminal", i)
	}

	// The text is centered: (80-5)/2 = 37 cells on the left, the rest padded.
	assert.Equal(t, strings.Repeat(" ", 37)+"Hello"+strings.Repeat(" ", 38), lines[1])
}

func TestHeaderPrinterMultilinePadsToWidestLine(t *testing.T) {
	expected := "               \n" +
		"     Hello     \n" +
		"     Hi        \n" +
		"               \n"

	assert.Equal(t, expected, stripANSI(pterm.DefaultHeader.Sprint("Hello\nHi")))
}

func TestHeaderPrinterUnicodeKeepsLinesEqualWidth(t *testing.T) {
	// "汉字" is 4 terminal cells wide, so all lines must be 4+2*5 = 14 cells.
	lines := headerLines(t, pterm.DefaultHeader.Sprint("汉字"))
	require.Len(t, lines, 3)

	for i, line := range lines {
		assert.Equalf(t, 14, runewidth.StringWidth(line), "line %d %q", i, line)
	}
}

func TestHeaderPrinterLongInputStaysWithinTerminalWidth(t *testing.T) {
	lines := headerLines(t, pterm.DefaultHeader.Sprint(strings.Repeat("a", 200)))
	require.Greater(t, len(lines), 3, "long input must be wrapped onto multiple lines")

	for i, line := range lines {
		assert.LessOrEqualf(t, len(line), 80, "line %d must not exceed the terminal width", i)
	}
}

func TestHeaderPrinterRawOutputPassesTextThrough(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableStyling()

	assert.Equal(t, "Hello", pterm.DefaultHeader.Sprint("Hello"))
}

func TestHeaderPrinterAppliesBackgroundStyle(t *testing.T) {
	// Every line of the block, including the blank ones, is wrapped in the
	// background style (BgCyan = SGR 46).
	for i, line := range strings.SplitAfter(strings.TrimSuffix(pterm.DefaultHeader.Sprint("Hi"), "\n"), "\n") {
		assert.Truef(t, strings.HasPrefix(line, "\x1b[46m"), "line %d %q must start with the background style", i, line)
	}
}

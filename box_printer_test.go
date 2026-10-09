package pterm_test

// Behavioral tests for BoxPrinter.
//
// Builder methods, Print*/Sprint* delegation and the global styling invariants
// are covered generically in contract_test.go. This file verifies the box
// geometry: border characters, padding, title placement and visible-width
// alignment for multiline, unicode and pre-colored content.

import (
	"strings"
	"testing"

	"github.com/mattn/go-runewidth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// boxLines strips ANSI codes from a rendered box and returns its lines.
func boxLines(s string) []string {
	return strings.Split(stripANSI(s), "\n")
}

// assertUniformVisibleWidth asserts that every rendered line occupies the same
// number of terminal cells, i.e. the right border is perfectly aligned.
func assertUniformVisibleWidth(t *testing.T, out string) {
	t.Helper()

	lines := boxLines(out)
	require.NotEmpty(t, lines)

	want := runewidth.StringWidth(lines[0])
	for i, line := range lines {
		assert.Equalf(t, want, runewidth.StringWidth(line), "line %d %q has a different visible width", i, line)
	}
}

func TestBoxPrinterDefaultBox(t *testing.T) {
	expected := "╭───────╮\n" +
		"│ Hello │\n" +
		"╰───────╯"

	assert.Equal(t, expected, stripANSI(pterm.DefaultBox.Sprint("Hello")))
}

func TestBoxPrinterMultilinePadsToWidestLine(t *testing.T) {
	expected := "╭──────────────╮\n" +
		"│ longest line │\n" +
		"│ short        │\n" +
		"╰──────────────╯"

	out := pterm.DefaultBox.Sprint("longest line\nshort")

	assert.Equal(t, expected, stripANSI(out))
	assertUniformVisibleWidth(t, out)
}

func TestBoxPrinterUnicodeContentKeepsBordersAligned(t *testing.T) {
	t.Run("CJK", func(t *testing.T) {
		// "汉字" is 4 terminal cells wide, not 6 bytes.
		expected := "╭──────╮\n" +
			"│ 汉字 │\n" +
			"╰──────╯"

		out := pterm.DefaultBox.Sprint("汉字")

		assert.Equal(t, expected, stripANSI(out))
		assertUniformVisibleWidth(t, out)
	})

	t.Run("emoji", func(t *testing.T) {
		expected := "╭──────────╮\n" +
			"│ 🦄 emoji │\n" +
			"╰──────────╯"

		out := pterm.DefaultBox.Sprint("🦄 emoji")

		assert.Equal(t, expected, stripANSI(out))
		assertUniformVisibleWidth(t, out)
	})
}

func TestBoxPrinterColoredContentKeepsBordersAligned(t *testing.T) {
	out := pterm.DefaultBox.Sprint("a " + pterm.FgRed.Sprint("red") + " b")

	// The escape codes must not count towards the width: the stripped output
	// is exactly the box around the plain text.
	assert.Equal(t, stripANSI(pterm.DefaultBox.Sprint("a red b")), stripANSI(out))
	assertUniformVisibleWidth(t, out)
}

func TestBoxPrinterTitlePlacement(t *testing.T) {
	// Content "Hello World" (11 cells) + 1 cell padding each side = 13 border
	// cells for the title line.
	box := pterm.DefaultBox.WithTitle("Title")

	tests := []struct {
		name     string
		printer  *pterm.BoxPrinter
		expected string
	}{
		{
			name:    "top left (default)",
			printer: box,
			expected: "╭─ Title ─────╮\n" +
				"│ Hello World │\n" +
				"╰─────────────╯",
		},
		{
			name:    "top right",
			printer: box.WithTitleTopRight(),
			expected: "╭───── Title ─╮\n" +
				"│ Hello World │\n" +
				"╰─────────────╯",
		},
		{
			name:    "top center",
			printer: box.WithTitleTopCenter(),
			expected: "╭─── Title ───╮\n" +
				"│ Hello World │\n" +
				"╰─────────────╯",
		},
		{
			name:    "bottom left",
			printer: box.WithTitleBottomLeft(),
			expected: "╭─────────────╮\n" +
				"│ Hello World │\n" +
				"╰─ Title ─────╯",
		},
		{
			name:    "bottom right",
			printer: box.WithTitleBottomRight(),
			expected: "╭─────────────╮\n" +
				"│ Hello World │\n" +
				"╰───── Title ─╯",
		},
		{
			name:    "bottom center",
			printer: box.WithTitleBottomCenter(),
			expected: "╭─────────────╮\n" +
				"│ Hello World │\n" +
				"╰─── Title ───╯",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := tc.printer.Sprint("Hello World")

			assert.Equal(t, tc.expected, stripANSI(out))
			assertUniformVisibleWidth(t, out)
		})
	}
}

func TestBoxPrinterTitleWiderThanContentGrowsTheBox(t *testing.T) {
	expected := "╭─ A very long title ─╮\n" +
		"│ Hi                  │\n" +
		"╰─────────────────────╯"

	out := pterm.DefaultBox.WithTitle("A very long title").Sprint("Hi")

	assert.Equal(t, expected, stripANSI(out))
	assertUniformVisibleWidth(t, out)
}

func TestBoxPrinterTitleNewlinesAreReplacedBySpaces(t *testing.T) {
	out := boxLines(pterm.DefaultBox.WithTitle("a\nb").Sprint("Hello World"))

	assert.Contains(t, out[0], " a b ")
	assert.Len(t, out, 3, "a newline in the title must not add lines to the box")
}

func TestBoxPrinterPadding(t *testing.T) {
	t.Run("all sides", func(t *testing.T) {
		expected := "╭──────╮\n" +
			"│      │\n" +
			"│      │\n" +
			"│  Hi  │\n" +
			"│      │\n" +
			"│      │\n" +
			"╰──────╯"

		assert.Equal(t, expected, stripANSI(pterm.DefaultBox.WithPadding(2).Sprint("Hi")))
	})

	t.Run("asymmetric horizontal", func(t *testing.T) {
		expected := "╭──────╮\n" +
			"│   Hi │\n" +
			"╰──────╯"

		p := pterm.DefaultBox.WithLeftPadding(3).WithRightPadding(1)

		assert.Equal(t, expected, stripANSI(p.Sprint("Hi")))
	})
}

// The corner-string fields are historically swapped: the string configured as
// BottomRightCornerString is rendered at the TOP LEFT, BottomLeftCornerString
// at the top right, TopRightCornerString at the bottom left and
// TopLeftCornerString at the bottom right (DefaultBox compensates by holding
// the visually matching glyphs, e.g. TopLeftCornerString = "╯"). Existing
// consumers — including _examples/demo — rely on this mapping, so it is locked
// here as-is instead of being "fixed".
func TestBoxPrinterCustomBorderStrings(t *testing.T) {
	p := pterm.DefaultBox.
		WithBottomRightCornerString("1").
		WithBottomLeftCornerString("2").
		WithTopRightCornerString("3").
		WithTopLeftCornerString("4").
		WithHorizontalString("-").
		WithVerticalString("|")

	expected := "1----2\n" +
		"| ab |\n" +
		"3----4"

	assert.Equal(t, expected, stripANSI(p.Sprint("ab")))
}

func TestBoxPrinterSprintlnAppendsNewline(t *testing.T) {
	expected := "╭────╮\n" +
		"│ Hi │\n" +
		"╰────╯\n"

	assert.Equal(t, expected, stripANSI(pterm.DefaultBox.Sprintln("Hi")))
}

func TestBoxPrinterEmptyInput(t *testing.T) {
	// An empty string still renders a (collapsed) box.
	expected := "╭──╮\n" +
		"│  │\n" +
		"╰──╯"

	assert.Equal(t, expected, stripANSI(pterm.DefaultBox.Sprint("")))
}

func TestBoxPrinterZeroValueDoesNotPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		p := pterm.BoxPrinter{}
		_ = p.Sprint("Hello, World!")
	})
}

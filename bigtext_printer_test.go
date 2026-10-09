package pterm_test

// Behavioral tests for BigTextPrinter: letters are rendered from the
// BigCharacters font map and joined horizontally with per-letter width
// padding; unknown characters are skipped. The builder/contract plumbing is
// covered in contract_test.go, one representative output in snapshot_test.go.

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mattn/go-runewidth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// joinedBigText builds the expected BigText output for the given characters
// from the default font map: the tallest letter defines the line count, every
// letter's lines are padded to that letter's width, and the letters are
// concatenated line by line.
func joinedBigText(t *testing.T, text string) string {
	t.Helper()

	entries := make([]string, 0, len(text))

	for _, r := range text {
		entry, ok := pterm.DefaultBigText.BigCharacters[string(r)]
		require.True(t, ok, "character %q must exist in the font map", r)

		entries = append(entries, entry)
	}

	var height int

	for _, entry := range entries {
		if h := len(strings.Split(entry, "\n")); h > height {
			height = h
		}
	}

	var sb strings.Builder

	for i := 0; i < height; i++ {
		for _, entry := range entries {
			lines := strings.Split(entry, "\n")

			var width int

			for _, line := range lines {
				if w := runewidth.StringWidth(line); w > width {
					width = w
				}
			}

			var line string
			if i < len(lines) {
				line = lines[i]
			}

			sb.WriteString(line + strings.Repeat(" ", width-runewidth.StringWidth(line)))
		}

		sb.WriteByte('\n')
	}

	return sb.String()
}

func TestBigTextPrinter_SingleLetterMatchesFontMap(t *testing.T) {
	printer := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("A"))

	// Every line of "A" in the font map is already 8 cells wide, so the
	// rendered output is exactly the map entry plus a trailing newline.
	assert.Equal(t, pterm.DefaultBigText.BigCharacters["A"]+"\n", srenderPlain(t, printer))
}

func TestBigTextPrinter_LettersAreJoinedHorizontally(t *testing.T) {
	printer := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("AB"))

	assert.Equal(t, joinedBigText(t, "AB"), srenderPlain(t, printer))
}

func TestBigTextPrinter_ShorterLettersArePaddedToTallestLetter(t *testing.T) {
	// "Q" is six lines tall while "A" only has five; "A" must be padded with
	// a line of spaces at the bottom.
	printer := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("AQ"))

	out := srenderPlain(t, printer)
	assert.Equal(t, joinedBigText(t, "AQ"), out)
	assert.Len(t, strings.Split(strings.TrimSuffix(out, "\n"), "\n"), 6)
}

func TestBigTextPrinter_UnknownCharactersAreSkipped(t *testing.T) {
	withUnknown := srenderPlain(t, pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("A~B")))
	without := srenderPlain(t, pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("AB")))

	assert.Equal(t, without, withUnknown)
}

func TestBigTextPrinter_NoLettersRenderSingleEmptyLine(t *testing.T) {
	// Documents current behavior: without (known) letters a single newline
	// is emitted.
	assert.Equal(t, "\n", srenderPlain(t, pterm.BigTextPrinter{}))
}

func TestBigTextPrinter_RawOutputReturnsPlainText(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableStyling()

	out, err := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Hi")).Srender()
	require.NoError(t, err)

	assert.Equal(t, "Hi", out)
}

func TestBigTextPrinter_LetterStyleIsApplied(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	printer := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromStringWithStyle("I", style))

	styled, err := printer.Srender()
	require.NoError(t, err)

	assert.Contains(t, styled, style.Sprint("██ "), "letter lines must be wrapped in the letter's style")
}

func TestBigTextPrinter_LetterRGBIsApplied(t *testing.T) {
	printer := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromStringWithRGB("I", pterm.NewRGB(255, 0, 0)))

	styled, err := printer.Srender()
	require.NoError(t, err)

	assert.Contains(t, styled, "38;2;255;0;0", "letter lines must use the letter's RGB color")
}

func TestBigTextPrinter_SrenderIsPure(t *testing.T) {
	letters := pterm.NewLettersFromString("AB")
	printer := pterm.DefaultBigText.WithLetters(letters)

	first, err := printer.Srender()
	require.NoError(t, err)

	second, err := printer.Srender()
	require.NoError(t, err)

	assert.Equal(t, first, second, "rendering twice must yield identical output")
	assert.Equal(t, pterm.NewLettersFromString("AB"), letters, "rendering must not modify the input letters")
}

func TestNewLettersFromString(t *testing.T) {
	t.Run("default style", func(t *testing.T) {
		assert.Equal(t, pterm.Letters{
			{String: "a", Style: &pterm.ThemeDefault.LetterStyle},
			{String: "b", Style: &pterm.ThemeDefault.LetterStyle},
		}, pterm.NewLettersFromString("ab"))
	})

	t.Run("with style", func(t *testing.T) {
		style := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
		assert.Equal(t, pterm.Letters{
			{String: "a", Style: style},
			{String: "b", Style: style},
		}, pterm.NewLettersFromStringWithStyle("ab", style))
	})

	t.Run("with RGB", func(t *testing.T) {
		rgb := pterm.NewRGB(1, 2, 3)
		assert.Equal(t, pterm.Letters{
			{String: "a", Style: pterm.NewStyle(), RGB: rgb},
			{String: "b", Style: pterm.NewStyle(), RGB: rgb},
		}, pterm.NewLettersFromStringWithRGB("ab", rgb))
	})
}

// Letter is a builder-style helper type that is not part of the printer list
// in contract_test.go, so its With* methods are verified here.
func TestLetter_Builders(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	rgb := pterm.NewRGB(1, 2, 3)
	letter := pterm.Letter{}.WithString("x").WithStyle(style).WithRGB(rgb)

	assert.Equal(t, pterm.Letter{String: "x", Style: style, RGB: rgb}, *letter)
}

// The letter join pads every letter to the height of the tallest letter in
// the rendered text, but printers combining BigText output (and the README
// examples) assume the default font stays at most six lines tall.
func TestDefaultLettersMaxHeight(t *testing.T) {
	maxHeight := 5

	for s, l := range pterm.DefaultBigText.BigCharacters {
		h := strings.Count(l, "\n")
		assert.LessOrEqual(t, h, maxHeight, fmt.Sprintf("%q is too high", s))
	}
}

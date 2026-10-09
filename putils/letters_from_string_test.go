package putils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestLettersFromString(t *testing.T) {
	letters := LettersFromString("ab")

	expected := pterm.Letters{
		{String: "a", Style: &pterm.ThemeDefault.LetterStyle},
		{String: "b", Style: &pterm.ThemeDefault.LetterStyle},
	}
	assert.Equal(t, expected, letters)

	// Every letter must reference the theme style, not a copy of it.
	for _, letter := range letters {
		assert.Same(t, &pterm.ThemeDefault.LetterStyle, letter.Style)
	}
}

func TestLettersFromStringEmpty(t *testing.T) {
	assert.Empty(t, LettersFromString(""))
}

func TestLettersFromStringSplitsRunes(t *testing.T) {
	letters := LettersFromString("Gö日")

	assert.Len(t, letters, 3)
	assert.Equal(t, "G", letters[0].String)
	assert.Equal(t, "ö", letters[1].String)
	assert.Equal(t, "日", letters[2].String)
}

func TestLettersFromStringWithStyle(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed, pterm.Bold)
	letters := LettersFromStringWithStyle("hi", style)

	expected := pterm.Letters{
		{String: "h", Style: style},
		{String: "i", Style: style},
	}
	assert.Equal(t, expected, letters)

	for _, letter := range letters {
		assert.Same(t, style, letter.Style)
	}
}

func TestLettersFromStringWithRGB(t *testing.T) {
	rgb := pterm.NewRGB(1, 2, 3)
	letters := LettersFromStringWithRGB("hi", rgb)

	expected := pterm.Letters{
		{String: "h", Style: &pterm.Style{}, RGB: rgb},
		{String: "i", Style: &pterm.Style{}, RGB: rgb},
	}
	assert.Equal(t, expected, letters)
}

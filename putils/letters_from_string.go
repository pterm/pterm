package putils

import (
	"strings"

	"github.com/pterm/pterm"
)

// LettersFromString creates a Letters object from a string, which is prefilled with the LetterStyle from ThemeDefault.
// You can override the ThemeDefault LetterStyle if you want to.
func LettersFromString(text string) pterm.Letters {
	return LettersFromStringWithStyle(text, &pterm.ThemeDefault.LetterStyle)
}

// LettersFromStringWithStyle creates a Letters object from a string and applies a Style to it.
func LettersFromStringWithStyle(text string, style *pterm.Style) pterm.Letters {
	s := strings.Split(text, "")
	l := pterm.Letters{}

	for _, s2 := range s {
		l = append(l, pterm.Letter{
			String: s2,
			Style:  style,
		})
	}

	return l
}

// LettersFromStringWithRGB creates a Letters object from a string and applies an RGB color to it (overwrites style).
func LettersFromStringWithRGB(text string, rgb pterm.RGB) pterm.Letters {
	s := strings.Split(text, "")
	l := pterm.Letters{}

	for _, s2 := range s {
		l = append(l, pterm.Letter{
			String: s2,
			Style:  &pterm.Style{},
			RGB:    rgb,
		})
	}

	return l
}

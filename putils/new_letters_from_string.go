package putils

import (
	"strings"

	"github.com/pterm/pterm"
)

// NewLettersFromString creates a Letters object from a string, which is prefilled with the LetterStyle from ThemeDefault.
// You can override the ThemeDefault LetterStyle if you want to.
func NewLettersFromString(text string) pterm.Letters {
	return NewLettersFromStringWithStyle(text, &pterm.ThemeDefault.LetterStyle)
}

// NewLettersFromStringWithStyle creates a Letters object from a string and applies a Style to it.
func NewLettersFromStringWithStyle(text string, style *pterm.Style) pterm.Letters {
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

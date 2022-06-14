package pterm

import "strings"

// NewLettersFromString creates a Letters object from a string, which is prefilled with the LetterStyle from ThemeDefault.
// You can override the ThemeDefault LetterStyle if you want to.
//
// Deprecated: use putils.NewLettersFromString instead.
func NewLettersFromString(text string) Letters {
	return NewLettersFromStringWithStyle(text, &ThemeDefault.LetterStyle)
}

// NewLettersFromStringWithStyle creates a Letters object from a string and applies a Style to it.
//
// Deprecated: use putils.NewLettersFromStringWithStyle instead.
func NewLettersFromStringWithStyle(text string, style *Style) Letters {
	s := strings.Split(text, "")
	l := Letters{}

	for _, s2 := range s {
		l = append(l, Letter{
			String: s2,
			Style:  style,
		})
	}

	return l
}

// NewLettersFromStringWithRGB creates a Letters object from a string and applies an RGB color to it (overwrites style).
//
// Deprecated: use putils.NewLettersFromStringWithRGB instead.
func NewLettersFromStringWithRGB(text string, rgb RGB) Letters {
	s := strings.Split(text, "")
	l := Letters{}

	for _, s2 := range s {
		l = append(l, Letter{
			String: s2,
			Style:  &Style{},
			RGB:    rgb,
		})
	}

	return l
}

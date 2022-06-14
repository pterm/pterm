package pterm

import (
	"strings"

	"github.com/pterm/pterm/internal"
)

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

// NewBulletListFromStrings returns a BulletListPrinter with Text using the NewTreeListItemFromString method.
//
// Deprecated: use putils.NewBulletListFromStrings instead.
func NewBulletListFromStrings(s []string, padding string) BulletListPrinter {
	var lis []BulletListItem
	for _, line := range s {
		lis = append(lis, NewBulletListItemFromString(line, padding))
	}
	return *DefaultBulletList.WithItems(lis)
}

// NewBulletListItemFromString returns a BulletListItem with a Text. The padding is counted in the Text to define the Level of the ListItem.
//
// Deprecated: use putils.NewBulletListItemFromString instead.
func NewBulletListItemFromString(text string, padding string) BulletListItem {
	s, l := internal.RemoveAndCountPrefix(text, padding)
	return BulletListItem{
		Level: l,
		Text:  s,
	}
}

// NewBulletListFromString returns a BulletListPrinter with Text using the NewTreeListItemFromString method, splitting after return (\n).
//
// Deprecated: use putils.NewBulletListFromString instead.
func NewBulletListFromString(s string, padding string) BulletListPrinter {
	return NewBulletListFromStrings(strings.Split(s, "\n"), padding)
}
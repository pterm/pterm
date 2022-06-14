package putils

import (
	"strings"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal"
)

// NewBulletListFromStrings returns a BulletListPrinter with Text using the NewTreeListItemFromString method.
func NewBulletListFromStrings(s []string, padding string) pterm.BulletListPrinter {
	var lis []pterm.BulletListItem
	for _, line := range s {
		lis = append(lis, NewBulletListItemFromString(line, padding))
	}
	return *pterm.DefaultBulletList.WithItems(lis)
}

// NewBulletListItemFromString returns a BulletListItem with a Text. The padding is counted in the Text to define the Level of the ListItem.
func NewBulletListItemFromString(text string, padding string) pterm.BulletListItem {
	s, l := internal.RemoveAndCountPrefix(text, padding)
	return pterm.BulletListItem{
		Level: l,
		Text:  s,
	}
}

// NewBulletListFromString returns a BulletListPrinter with Text using the NewTreeListItemFromString method, splitting after return (\n).
func NewBulletListFromString(s string, padding string) pterm.BulletListPrinter {
	return NewBulletListFromStrings(strings.Split(s, "\n"), padding)
}

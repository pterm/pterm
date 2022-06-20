package putils

import (
	"strings"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal"
)

// BulletListFromStrings returns a BulletListPrinter with Text using the NewTreeListItemFromString method.
func BulletListFromStrings(s []string, padding string) pterm.BulletListPrinter {
	var lis []pterm.BulletListItem
	for _, line := range s {
		lis = append(lis, BulletListItemFromString(line, padding))
	}
	return *pterm.DefaultBulletList.WithItems(lis)
}

// BulletListItemFromString returns a BulletListItem with a Text. The padding is counted in the Text to define the Level of the ListItem.
func BulletListItemFromString(text string, padding string) pterm.BulletListItem {
	s, l := internal.RemoveAndCountPrefix(text, padding)
	return pterm.BulletListItem{
		Level: l,
		Text:  s,
	}
}

// BulletListFromString returns a BulletListPrinter with Text using the NewTreeListItemFromString method, splitting after return (\n).
func BulletListFromString(s string, padding string) pterm.BulletListPrinter {
	return BulletListFromStrings(strings.Split(s, "\n"), padding)
}

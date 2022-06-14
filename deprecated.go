package pterm

import (
	"strconv"
	"strings"

	"github.com/pterm/pterm/internal"
)

// NewLettersFromString creates a Letters object from a string, which is prefilled with the LetterStyle from ThemeDefault.
// You can override the ThemeDefault LetterStyle if you want to.
//
// Deprecated: use putils.LettersFromString instead.
func NewLettersFromString(text string) Letters {
	return NewLettersFromStringWithStyle(text, &ThemeDefault.LetterStyle)
}

// NewLettersFromStringWithStyle creates a Letters object from a string and applies a Style to it.
//
// Deprecated: use putils.LettersFromStringWithStyle instead.
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
// Deprecated: use putils.LettersFromStringWithRGB instead.
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
// Deprecated: use putils.BulletListFromStrings instead.
func NewBulletListFromStrings(s []string, padding string) BulletListPrinter {
	var lis []BulletListItem
	for _, line := range s {
		lis = append(lis, NewBulletListItemFromString(line, padding))
	}
	return *DefaultBulletList.WithItems(lis)
}

// NewBulletListItemFromString returns a BulletListItem with a Text. The padding is counted in the Text to define the Level of the ListItem.
//
// Deprecated: use putils.BulletListItemFromString instead.
func NewBulletListItemFromString(text string, padding string) BulletListItem {
	s, l := internal.RemoveAndCountPrefix(text, padding)
	return BulletListItem{
		Level: l,
		Text:  s,
	}
}

// NewBulletListFromString returns a BulletListPrinter with Text using the NewTreeListItemFromString method, splitting after return (\n).
//
// Deprecated: use putils.BulletListFromString instead.
func NewBulletListFromString(s string, padding string) BulletListPrinter {
	return NewBulletListFromStrings(strings.Split(s, "\n"), padding)
}

// NewTreeFromLeveledList converts a TreeItems list to a TreeNode and returns it.
//
// Deprecated: use putils.TreeFromLeveledList instead.
func NewTreeFromLeveledList(leveledListItems LeveledList) TreeNode {
	if len(leveledListItems) == 0 {
		return TreeNode{}
	}

	root := &TreeNode{
		Children: []TreeNode{},
		Text:     leveledListItems[0].Text,
	}

	for i, record := range leveledListItems {
		last := root

		if record.Level < 0 {
			record.Level = 0
			leveledListItems[i].Level = 0
		}

		if len(leveledListItems)-1 != i {
			if leveledListItems[i+1].Level-1 > record.Level {
				leveledListItems[i+1].Level = record.Level + 1
			}
		}

		for i := 0; i < record.Level; i++ {
			lastIndex := len(last.Children) - 1
			last = &last.Children[lastIndex]
		}
		last.Children = append(last.Children, TreeNode{
			Children: []TreeNode{},
			Text:     record.Text,
		})
	}

	return *root
}

// NewRGBFromHEX converts a HEX and returns a new RGB.
//
// Deprecated: use putils.RGBFromHEX instead.
func NewRGBFromHEX(hex string) (RGB, error) {
	hex = strings.ToLower(hex)
	hex = strings.ReplaceAll(hex, "#", "")
	hex = strings.ReplaceAll(hex, "0x", "")

	if len(hex) == 3 {
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}
	if len(hex) != 6 {
		return RGB{}, ErrHexCodeIsInvalid
	}

	i64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return RGB{}, err
	}
	c := int(i64)

	return RGB{
		R: uint8(c >> 16),
		G: uint8((c & 0x00FF00) >> 8),
		B: uint8(c & 0x0000FF),
	}, nil
}

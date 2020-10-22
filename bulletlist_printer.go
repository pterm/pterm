package pterm

import (
	"strings"

	"github.com/pterm/pterm/internal"
)

// NewListFromStrings returns a BulletList with Letters using the NewListItemFromString method.
func NewListFromStrings(s []string, padding string) BulletList {
	var lis []BulletListItem
	for _, line := range s {
		lis = append(lis, NewListItemFromString(line, padding))
	}
	return *DefaultBulletList.WithItems(lis)
}

// NewListItemFromString returns a ListItem with a Letters. The padding is counted in the Letters to define the Level of the ListItem.
func NewListItemFromString(text string, padding string) BulletListItem {
	s, l := internal.RemoveAndCountPrefix(text, padding)
	return BulletListItem{
		Level: l,
		Text:  s,
	}
}

// DefaultListItem contains standards, which can be used to print a ListItem.
var DefaultListItem = BulletListItem{
	Bullet:      "•",
	TextStyle:   &ThemeDefault.ListTextStyle,
	BulletStyle: &ThemeDefault.ListBulletStyle,
}

// BulletListItem is able to render a ListItem.
type BulletListItem struct {
	Level       int
	Text        string
	TextStyle   *Style
	Bullet      string
	BulletStyle *Style
}

// WithLevel returns a new ListItem with a specific Level.
func (p BulletListItem) WithLevel(level int) BulletListItem {
	p.Level = level
	return p
}

// WithText returns a new ListItem with a specific Text.
func (p BulletListItem) WithText(text string) BulletListItem {
	p.Text = text
	return p
}

// WithTextStyle returns a new ListItem with a specific TextStyle.
func (p BulletListItem) WithTextStyle(style *Style) BulletListItem {
	p.TextStyle = style
	return p
}

// WithBullet returns a new ListItem with a specific Prefix.
func (p BulletListItem) WithBullet(bullet string) BulletListItem {
	p.Bullet = bullet
	return p
}

// WithBulletStyle returns a new ListItem with a specific BulletStyle.
func (p BulletListItem) WithBulletStyle(style *Style) BulletListItem {
	p.BulletStyle = style
	return p
}

// Render renders the ListItem as a string.
func (p BulletListItem) Render() {
	Println(p.Srender())
}

// Srender renders the ListItem as a string.
func (p BulletListItem) Srender() string {
	if p.TextStyle == nil {
		p.TextStyle = NewStyle()
	}
	if p.BulletStyle == nil {
		p.BulletStyle = &ThemeDefault.ListBulletStyle
	}
	return strings.Repeat(" ", p.Level) + p.BulletStyle.Sprint(p.Bullet) + " " + p.TextStyle.Sprint(p.Text)
}

// NewListFromString returns a BulletList with Letters using the NewListItemFromString method, splitting after return (\n).
func NewListFromString(s string, padding string) BulletList {
	return NewListFromStrings(strings.Split(s, "\n"), padding)
}

// DefaultBulletList contains standards, which can be used to print a BulletList.
var DefaultBulletList = BulletList{
	Bullet:      "•",
	TextStyle:   &ThemeDefault.ListTextStyle,
	BulletStyle: &ThemeDefault.ListBulletStyle,
}

// BulletList is able to render a list.
type BulletList struct {
	Items       []BulletListItem
	TextStyle   *Style
	Bullet      string
	BulletStyle *Style
}

// WithItems returns a new list with specific Items.
func (l BulletList) WithItems(items []BulletListItem) *BulletList {
	l.Items = append(l.Items, items...)
	return &l
}

// WithTextStyle returns a new list with a specific text style.
func (l BulletList) WithTextStyle(style *Style) *BulletList {
	l.TextStyle = style
	return &l
}

// WithBullet returns a new list with a specific bullet.
func (l BulletList) WithBullet(bullet string) *BulletList {
	l.Bullet = bullet
	return &l
}

// WithBulletStyle returns a new list with a specific bullet style.
func (l BulletList) WithBulletStyle(style *Style) *BulletList {
	l.BulletStyle = style
	return &l
}

// Render prints the list to the terminal.
func (l BulletList) Render() {
	Println(l.Srender())
}

// Srender renders the list as a string.
func (l BulletList) Srender() string {
	var ret string
	for _, item := range l.Items {
		if item.Bullet == "" {
			ret += item.WithBullet(l.Bullet).Srender() + "\n"
		} else {
			ret += item.WithBullet(item.Bullet).Srender() + "\n"
		}
	}
	return ret
}

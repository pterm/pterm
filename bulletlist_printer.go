package pterm

import (
	"strings"

	"github.com/pterm/pterm/internal"
)

// NewBulletListFromStrings returns a BulletList with Text using the NewTreeListItemFromString method.
func NewBulletListFromStrings(s []string, padding string) BulletList {
	var lis []BulletListItem
	for _, line := range s {
		lis = append(lis, NewBulletListItemFromString(line, padding))
	}
	return *DefaultBulletList.WithItems(lis)
}

// NewBulletListItemFromString returns a BulletListItem with a Text. The padding is counted in the Text to define the Level of the ListItem.
func NewBulletListItemFromString(text string, padding string) BulletListItem {
	s, l := internal.RemoveAndCountPrefix(text, padding)
	return BulletListItem{
		Level: l,
		Text:  s,
	}
}

// DefaultListItem contains standards, which can be used to print a ListItem.
var DefaultListItem = BulletListItem{
	Bullet:      "•",
	TextStyle:   &ThemeDefault.BulletListTextStyle,
	BulletStyle: &ThemeDefault.BulletListBulletStyle,
}

// BulletListItem is able to render a ListItem.
type BulletListItem struct {
	Level       int
	Text        string
	TextStyle   *Style
	Bullet      string
	BulletStyle *Style
}

// WithLevel returns a new BulletListItem with a specific Level.
func (p BulletListItem) WithLevel(level int) BulletListItem {
	p.Level = level
	return p
}

// WithText returns a new BulletListItem with a specific Text.
func (p BulletListItem) WithText(text string) BulletListItem {
	p.Text = text
	return p
}

// WithTextStyle returns a new BulletListItem with a specific TextStyle.
func (p BulletListItem) WithTextStyle(style *Style) BulletListItem {
	p.TextStyle = style
	return p
}

// WithBullet returns a new BulletListItem with a specific Prefix.
func (p BulletListItem) WithBullet(bullet string) BulletListItem {
	p.Bullet = bullet
	return p
}

// WithBulletStyle returns a new BulletListItem with a specific BulletStyle.
func (p BulletListItem) WithBulletStyle(style *Style) BulletListItem {
	p.BulletStyle = style
	return p
}

// Render renders the BulletListItem as a string.
func (p BulletListItem) Render() {
	Println(p.Srender())
}

// Srender renders the BulletListItem as a string.
func (p BulletListItem) Srender() string {
	if p.TextStyle == nil {
		p.TextStyle = NewStyle()
	}
	if p.BulletStyle == nil {
		p.BulletStyle = &ThemeDefault.BulletListBulletStyle
	}
	return strings.Repeat(" ", p.Level) + p.BulletStyle.Sprint(p.Bullet) + " " + p.TextStyle.Sprint(p.Text)
}

// NewBulletListFromString returns a BulletList with Text using the NewTreeListItemFromString method, splitting after return (\n).
func NewBulletListFromString(s string, padding string) BulletList {
	return NewBulletListFromStrings(strings.Split(s, "\n"), padding)
}

// DefaultBulletList contains standards, which can be used to print a BulletList.
var DefaultBulletList = BulletList{
	Bullet:      "•",
	TextStyle:   &ThemeDefault.BulletListTextStyle,
	BulletStyle: &ThemeDefault.BulletListBulletStyle,
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
func (l BulletList) Render() error {
	s, err := l.Srender()
	if err != nil {
		return err
	}
	Println(s)

	return nil
}

// Srender renders the list as a string.
func (l BulletList) Srender() (string, error) {
	var ret string
	for _, item := range l.Items {
		if item.Bullet == "" {
			ret += item.WithBullet(l.Bullet).Srender() + "\n"
		} else {
			ret += item.WithBullet(item.Bullet).Srender() + "\n"
		}
	}
	return ret, nil
}

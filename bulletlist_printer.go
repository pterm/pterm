package pterm

import (
	"io"
	"strings"
)

// BulletListItem is able to render a ListItem.
type BulletListItem struct {
	Level       int
	Text        string
	TextStyle   *Style
	Bullet      string
	BulletStyle *Style
}

// WithLevel returns a new BulletListItem with a specific Level.
func (p BulletListItem) WithLevel(level int) *BulletListItem {
	p.Level = level
	return &p
}

// WithText returns a new BulletListItem with a specific Text.
func (p BulletListItem) WithText(text string) *BulletListItem {
	p.Text = text
	return &p
}

// WithTextStyle returns a new BulletListItem with a specific TextStyle.
func (p BulletListItem) WithTextStyle(style *Style) *BulletListItem {
	p.TextStyle = style
	return &p
}

// WithBullet returns a new BulletListItem with a specific Prefix.
func (p BulletListItem) WithBullet(bullet string) *BulletListItem {
	p.Bullet = bullet
	return &p
}

// WithBulletStyle returns a new BulletListItem with a specific BulletStyle.
func (p BulletListItem) WithBulletStyle(style *Style) *BulletListItem {
	p.BulletStyle = style
	return &p
}

// DefaultBulletList contains standards, which can be used to print a BulletListPrinter.
var DefaultBulletList = BulletListPrinter{
	Bullet:      "•",
	TextStyle:   &ThemeDefault.BulletListTextStyle,
	BulletStyle: &ThemeDefault.BulletListBulletStyle,
}

// BulletListPrinter is able to render a list.
type BulletListPrinter struct {
	Items       []BulletListItem
	TextStyle   *Style
	Bullet      string
	BulletStyle *Style
	Writer      io.Writer
}

// WithItems returns a new list with specific Items.
func (l BulletListPrinter) WithItems(items []BulletListItem) *BulletListPrinter {
	l.Items = append(l.Items, items...)
	return &l
}

// WithTextStyle returns a new list with a specific text style.
func (l BulletListPrinter) WithTextStyle(style *Style) *BulletListPrinter {
	l.TextStyle = style
	return &l
}

// WithBullet returns a new list with a specific bullet.
func (l BulletListPrinter) WithBullet(bullet string) *BulletListPrinter {
	l.Bullet = bullet
	return &l
}

// WithBulletStyle returns a new list with a specific bullet style.
func (l BulletListPrinter) WithBulletStyle(style *Style) *BulletListPrinter {
	l.BulletStyle = style
	return &l
}

// WithWriter sets the custom Writer.
func (l BulletListPrinter) WithWriter(writer io.Writer) *BulletListPrinter {
	l.Writer = writer
	return &l
}

// Render prints the list to the terminal.
func (l BulletListPrinter) Render() error {
	s, _ := l.Srender()
	Fprintln(l.Writer, s)

	return nil
}

// Srender renders the list as a string.
func (l BulletListPrinter) Srender() (string, error) {
	var ret string
	for _, item := range l.Items {
		if item.TextStyle == nil {
			if l.TextStyle == nil {
				item.TextStyle = &ThemeDefault.BulletListTextStyle
			} else {
				item.TextStyle = l.TextStyle
			}
		}
		if item.BulletStyle == nil {
			if l.BulletStyle == nil {
				item.BulletStyle = &ThemeDefault.BulletListBulletStyle
			} else {
				item.BulletStyle = l.BulletStyle
			}
		}

		split := strings.Split(item.Text, "\n")
		for i, line := range split {
			if i == 0 {
				if item.Bullet == "" {
					ret += strings.Repeat(" ", item.Level) + item.BulletStyle.Sprint(l.Bullet) + " " + item.TextStyle.Sprint(line) + "\n"
				} else {
					ret += strings.Repeat(" ", item.Level) + item.BulletStyle.Sprint(item.Bullet) + " " + item.TextStyle.Sprint(line) + "\n"
				}
			} else {
				ret += strings.Repeat(" ", item.Level) + strings.Repeat(" ", len(item.Bullet)) + "  " + item.TextStyle.Sprint(line) + "\n"
			}
		}
	}
	return ret, nil
}

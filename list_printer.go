package pterm

import (
	"strings"

	"github.com/pterm/pterm/internal"
)

func NewListFromStrings(s []string, padding string) BulletList {
	var lis []ListItem
	for _, line := range s {
		lis = append(lis, NewListItemFromString(line, padding))
	}
	return *DefaultBulletList.WithItems(lis)
}

func NewListItemFromString(text string, padding string) ListItem {
	s, l := internal.RemoveAndCountPrefix(text, padding)
	return ListItem{
		Level: l,
		Text:  s,
	}
}

var DefaultListItem = ListItem{
	Prefix: "•",
}

type ListItem struct {
	Level       int
	Text        string
	TextStyle   Style
	Prefix      string
	BulletStyle Style
}

func (p ListItem) WithLevel(level int) ListItem {
	p.Level = level
	return p
}

func (p ListItem) WithText(text string) ListItem {
	p.Text = text
	return p
}

func (p ListItem) WithTextStyle(style Style) ListItem {
	p.TextStyle = style
	return p
}

func (p ListItem) WithBullet(bullet string) ListItem {
	p.Prefix = bullet
	return p
}

func (p ListItem) WithBulletStyle(style Style) ListItem {
	p.BulletStyle = style
	return p
}

func (p ListItem) Render() {
	Println(p.Srender())
}

func (p ListItem) Srender() string {
	return strings.Repeat(" ", p.Level) + p.Prefix + " " + p.Text
}

func NewListFromString(s string, padding string) BulletList {
	return NewListFromStrings(strings.Split(s, "\n"), padding)
}

var DefaultBulletList = BulletList{
	Bullet: "•",
}

type BulletList struct {
	Items       []ListItem
	TextStyle   Style
	Bullet      string
	BulletStyle Style
}

func (l BulletList) WithItems(items []ListItem) *BulletList {
	l.Items = append(l.Items, items...)
	return &l
}

func (l BulletList) WithTextStyle(style Style) *BulletList {
	l.TextStyle = style
	return &l
}

func (l BulletList) WithBullet(bullet string) *BulletList {
	l.Bullet = bullet
	return &l
}

func (l BulletList) WithBulletStyle(style Style) *BulletList {
	l.BulletStyle = style
	return &l
}

func (l BulletList) Render() {
	Println(l.Srender())
}

func (l BulletList) Srender() string {
	var ret string
	for _, item := range l.Items {
		ret += item.WithBullet(l.Bullet).Srender() + "\n"
	}
	return ret
}

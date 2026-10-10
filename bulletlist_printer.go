package pterm

import (
	"io"
	"strconv"
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// BulletListItem is one item of a bullet list. Bullet and Tone override the
// list's bullet and bullet tone for this item; leave them empty to keep the
// list's.
type BulletListItem struct {
	// Level is the nesting depth, starting at 0.
	Level int
	// Text is the item. Long text wraps under itself.
	Text string
	// Bullet replaces the bullet of the list for this item.
	Bullet string
	// Tone colors the bullet of this item.
	Tone Tone
}

// BulletListPrinter renders a list with muted bullets. Every level indents
// its items by two more columns, and long items wrap under their own text.
// An ordered list numbers its items instead, counting per level.
type BulletListPrinter struct {
	// Items are the items in order.
	Items []BulletListItem
	// Bullet is drawn before every item. Empty uses the bullet of the glyph
	// set.
	Bullet string
	// BulletTone colors the bullets and numbers.
	BulletTone Tone
	// TextStyle styles the items.
	TextStyle Style
	// Ordered numbers the items per level instead of drawing bullets.
	Ordered bool
	// Indent is the number of columns each level is indented by. Ordered
	// lists indent nested items under the text of their parent instead.
	Indent int
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// BulletList is the default bullet list.
var BulletList = BulletListPrinter{
	BulletTone: ToneMuted,
	TextStyle:  ToneText.Style(),
	Indent:     2,
}

// WithItems returns a copy of p that renders items.
func (p BulletListPrinter) WithItems(items []BulletListItem) *BulletListPrinter {
	p.Items = items
	return &p
}

// WithTexts returns a copy of p with one top-level item per text.
func (p BulletListPrinter) WithTexts(texts ...string) *BulletListPrinter {
	p.Items = make([]BulletListItem, len(texts))
	for i, t := range texts {
		p.Items[i] = BulletListItem{Text: t}
	}

	return &p
}

// WithBullet returns a copy of p that draws b before every item.
func (p BulletListPrinter) WithBullet(b string) *BulletListPrinter {
	p.Bullet = b
	return &p
}

// WithBulletTone returns a copy of p whose bullets and numbers use tone t.
func (p BulletListPrinter) WithBulletTone(t Tone) *BulletListPrinter {
	p.BulletTone = t
	return &p
}

// WithTextStyle returns a copy of p whose items use s.
func (p BulletListPrinter) WithTextStyle(s Style) *BulletListPrinter {
	p.TextStyle = s
	return &p
}

// WithOrdered returns a copy of p that numbers its items, or draws bullets
// again with false.
func (p BulletListPrinter) WithOrdered(b ...bool) *BulletListPrinter {
	p.Ordered = flag(b)
	return &p
}

// WithIndent returns a copy of p that indents each level by n columns.
func (p BulletListPrinter) WithIndent(n int) *BulletListPrinter {
	p.Indent = max(n, 0)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p BulletListPrinter) WithWriter(w io.Writer) *BulletListPrinter {
	p.Writer = w
	return &p
}

func (p BulletListPrinter) out() io.Writer { return orDefault(p.Writer) }

// Render prints the list.
func (p BulletListPrinter) Render() error {
	s, err := p.Srender()
	return renderTo(p.out(), s, err)
}

// Srender returns the list. It returns ErrEmpty when there are no items.
func (p BulletListPrinter) Srender() (string, error) {
	if len(p.Items) == 0 {
		return "", ErrEmpty
	}

	c := canvasFor(p.out())
	markers := p.markers(c)
	widths := markerWidths(p.Items, markers)

	var lines []string

	for i, item := range p.Items {
		level := max(item.Level, 0)
		indent := p.indent(level, widths)

		tone := p.BulletTone
		if item.Tone != (Tone{}) {
			tone = item.Tone
		}

		marker := markers[i]
		if p.Ordered {
			marker = text.PadLeft(marker, widths[level])
		}

		hang := indent + text.Width(marker) + 1
		body := c.wrap(item.Text, hang)
		pad := strings.Repeat(" ", indent)

		for j, line := range body {
			lead := c.paint(tone, marker) + " "
			if j > 0 {
				lead = strings.Repeat(" ", text.Width(marker)+1)
			}

			lines = append(lines, pad+lead+c.style(p.TextStyle, line))
		}
	}

	return strings.Join(lines, "\n"), nil
}

// indent returns the indent of an item on level. Nested items of an
// ordered list start under the text of their parent, past its number.
func (p BulletListPrinter) indent(level int, widths []int) int {
	if !p.Ordered {
		return level * p.Indent
	}

	indent := 0
	for l := range level {
		indent += widths[l] + 1
	}

	return indent
}

// markers returns the bullet or number of every item. Numbers count per
// level and start again under every new parent.
func (p BulletListPrinter) markers(c canvas) []string {
	markers := make([]string, len(p.Items))

	var counts []int

	for i, item := range p.Items {
		level := max(item.Level, 0)

		if !p.Ordered || item.Bullet != "" {
			markers[i] = item.Bullet
			if markers[i] == "" {
				markers[i] = p.Bullet
			}

			if markers[i] == "" {
				markers[i] = c.g.Bullet
			}
		}

		for len(counts) <= level {
			counts = append(counts, 0)
		}

		counts = counts[:level+1]
		counts[level]++

		if markers[i] == "" {
			markers[i] = strconv.Itoa(counts[level]) + "."
		}
	}

	return markers
}

// markerWidths returns the widest marker per level, so numbers of one
// level line up on their dots.
func markerWidths(items []BulletListItem, markers []string) []int {
	var widths []int

	for i, item := range items {
		level := max(item.Level, 0)
		for len(widths) <= level {
			widths = append(widths, 0)
		}

		widths[level] = max(widths[level], text.Width(markers[i]))
	}

	return widths
}

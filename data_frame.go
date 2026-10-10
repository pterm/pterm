package pterm

import (
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// dataFrame draws the rounded frame that boxed tables and panels share:
//
//	╭─ Title ───╮
//	│ content   │
//	├───────────┤
//	╰───────────╯
//
// inner is the width between the two vertical lines.
type dataFrame struct {
	c     canvas
	tone  Tone
	inner int
}

// top returns the top line, with title inset after one horizontal line.
// A title that does not fit is cut.
func (f dataFrame) top(title string, titleStyle Style) string {
	b := f.c.g.Box
	if title == "" || f.inner < 5 {
		return f.c.paint(f.tone, b.TopLeft+strings.Repeat(b.Horizontal, f.inner)+b.TopRight)
	}

	title = f.c.cut(title, f.inner-4)
	rest := f.inner - 3 - text.Width(title)

	return f.c.paint(f.tone, b.TopLeft+b.Horizontal+" ") +
		f.c.style(titleStyle, title) +
		f.c.paint(f.tone, " "+strings.Repeat(b.Horizontal, rest)+b.TopRight)
}

// rule returns a separator line that joins both sides with tees.
func (f dataFrame) rule() string {
	b := f.c.g.Box
	return f.c.paint(f.tone, b.TeeLeft+strings.Repeat(b.Horizontal, f.inner)+b.TeeRight)
}

// bottom returns the bottom line.
func (f dataFrame) bottom() string {
	b := f.c.g.Box
	return f.c.paint(f.tone, b.BottomLeft+strings.Repeat(b.Horizontal, f.inner)+b.BottomRight)
}

// line frames one line of content that is already inner columns wide.
func (f dataFrame) line(s string) string {
	v := f.c.paint(f.tone, f.c.g.Box.Vertical)
	return v + s + v
}

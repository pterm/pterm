package pterm

import (
	"fmt"
	"io"
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// BoxTitlePosition is where the title of a box sits in its frame.
type BoxTitlePosition uint8

// The title positions of a box.
const (
	BoxTitleTopLeft BoxTitlePosition = iota
	BoxTitleTopCenter
	BoxTitleTopRight
	BoxTitleBottomLeft
	BoxTitleBottomCenter
	BoxTitleBottomRight
)

// BoxPrinter draws a rounded frame around text, with an optional title set
// into the border:
//
//	╭─ Result ───────────╮
//	│ 3 files changed    │
//	╰────────────────────╯
//
// Boxes are text printers, so a box can hold the output of other printers,
// including other boxes.
type BoxPrinter struct {
	// Title is set into the border. Empty draws a plain frame.
	Title string
	// TitlePosition is where the title sits.
	TitlePosition BoxTitlePosition
	// TitleInset is the number of border glyphs between the corner and a
	// title on the left or right.
	TitleInset int
	// TitleStyle styles the title.
	TitleStyle Style
	// TextStyle styles the text inside the box.
	TextStyle Style
	// BorderTone colors the frame.
	BorderTone Tone
	// Glyphs draw the frame. Empty fields use the box glyphs of the glyph
	// set.
	Glyphs BoxGlyphs
	// TopPadding, BottomPadding, LeftPadding and RightPadding are the
	// space between the frame and the text, in lines and columns.
	TopPadding, BottomPadding, LeftPadding, RightPadding int
	// Width is the width of the whole box in columns. Zero fits the box to
	// its text, up to the terminal width. Longer lines wrap.
	Width int
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Box is the default box.
var Box = BoxPrinter{
	TitleInset:   1,
	TitleStyle:   TonePrimary.Bold(),
	TextStyle:    ToneText.Style(),
	BorderTone:   ToneBorder,
	LeftPadding:  1,
	RightPadding: 1,
}

// WithTitle returns a copy of p with title t.
func (p BoxPrinter) WithTitle(t string) *BoxPrinter {
	p.Title = t
	return &p
}

// WithTitlePosition returns a copy of p with the title at pos.
func (p BoxPrinter) WithTitlePosition(pos BoxTitlePosition) *BoxPrinter {
	p.TitlePosition = pos
	return &p
}

// WithTitleInset returns a copy of p with n border glyphs between the
// corner and the title.
func (p BoxPrinter) WithTitleInset(n int) *BoxPrinter {
	p.TitleInset = max(n, 0)
	return &p
}

// WithTitleStyle returns a copy of p that styles the title with s.
func (p BoxPrinter) WithTitleStyle(s Style) *BoxPrinter {
	p.TitleStyle = s
	return &p
}

// WithTextStyle returns a copy of p that styles the text with s.
func (p BoxPrinter) WithTextStyle(s Style) *BoxPrinter {
	p.TextStyle = s
	return &p
}

// WithBorderTone returns a copy of p whose frame is in tone t.
func (p BoxPrinter) WithBorderTone(t Tone) *BoxPrinter {
	p.BorderTone = t
	return &p
}

// WithGlyphs returns a copy of p that draws its frame with g. Empty fields
// of g keep the glyphs of the glyph set.
func (p BoxPrinter) WithGlyphs(g BoxGlyphs) *BoxPrinter {
	p.Glyphs = g
	return &p
}

// WithPadding returns a copy of p with n columns of space left and right
// of the text and n lines above and below it.
func (p BoxPrinter) WithPadding(n int) *BoxPrinter {
	n = max(n, 0)
	p.TopPadding, p.BottomPadding, p.LeftPadding, p.RightPadding = n, n, n, n

	return &p
}

// WithHorizontalPadding returns a copy of p with n columns of space left
// and right of the text.
func (p BoxPrinter) WithHorizontalPadding(n int) *BoxPrinter {
	p.LeftPadding, p.RightPadding = max(n, 0), max(n, 0)
	return &p
}

// WithVerticalPadding returns a copy of p with n blank lines above and
// below the text.
func (p BoxPrinter) WithVerticalPadding(n int) *BoxPrinter {
	p.TopPadding, p.BottomPadding = max(n, 0), max(n, 0)
	return &p
}

// WithTopPadding returns a copy of p with n blank lines above the text.
func (p BoxPrinter) WithTopPadding(n int) *BoxPrinter {
	p.TopPadding = max(n, 0)
	return &p
}

// WithBottomPadding returns a copy of p with n blank lines below the text.
func (p BoxPrinter) WithBottomPadding(n int) *BoxPrinter {
	p.BottomPadding = max(n, 0)
	return &p
}

// WithLeftPadding returns a copy of p with n columns of space left of the
// text.
func (p BoxPrinter) WithLeftPadding(n int) *BoxPrinter {
	p.LeftPadding = max(n, 0)
	return &p
}

// WithRightPadding returns a copy of p with n columns of space right of
// the text.
func (p BoxPrinter) WithRightPadding(n int) *BoxPrinter {
	p.RightPadding = max(n, 0)
	return &p
}

// WithWidth returns a copy of p that is w columns wide. Zero fits the box
// to its text.
func (p BoxPrinter) WithWidth(w int) *BoxPrinter {
	p.Width = max(w, 0)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p BoxPrinter) WithWriter(w io.Writer) *BoxPrinter {
	p.Writer = w
	return &p
}

func (p BoxPrinter) out() io.Writer { return orDefault(p.Writer) }

// glyphs returns the frame glyphs, with empty fields taken from g.
func (p BoxPrinter) glyphs(g BoxGlyphs) BoxGlyphs {
	pick := func(own, fallback string) string {
		if own != "" {
			return own
		}

		return fallback
	}

	return BoxGlyphs{
		TopLeft:     pick(p.Glyphs.TopLeft, g.TopLeft),
		TopRight:    pick(p.Glyphs.TopRight, g.TopRight),
		BottomLeft:  pick(p.Glyphs.BottomLeft, g.BottomLeft),
		BottomRight: pick(p.Glyphs.BottomRight, g.BottomRight),
		Horizontal:  pick(p.Glyphs.Horizontal, g.Horizontal),
		Vertical:    pick(p.Glyphs.Vertical, g.Vertical),
	}
}

func (p BoxPrinter) render(c canvas, s string) string {
	g := p.glyphs(c.g.Box)
	title := strings.Join(strings.Fields(p.Title), " ")
	hw := max(text.Width(g.Horizontal), 1)
	vw := text.Width(g.Vertical)
	pad := p.LeftPadding + p.RightPadding
	frame := text.Width(g.TopLeft) + text.Width(g.TopRight)

	// limit is the widest the box may get, 0 when there is no limit.
	limit := c.width
	if p.Width > 0 {
		limit = p.Width
		if c.width > 0 {
			limit = min(limit, c.width)
		}
	}

	textLimit := 0
	if limit > 0 {
		textLimit = max(limit-2*vw-pad, 1)
	}

	lines := text.Wrap(strings.TrimRight(s, "\n"), textLimit)

	// inner is the width between the two vertical lines.
	inner := text.MaxWidth(strings.Join(lines, "\n")) + pad
	if title != "" {
		inner = max(inner, text.Width(title)+2+2*p.TitleInset*hw)
	}

	if p.Width > 0 {
		inner = max(inner, p.Width-2*vw)
	}

	if limit > 0 {
		inner = min(inner, limit-2*vw)
	}

	inner = max(inner, 0)
	border := func(n int) string {
		return c.paint(p.BorderTone, strings.Repeat(g.Horizontal, max(n, 0)/hw))
	}
	edge := func(left, right string, withTitle bool) string {
		if !withTitle || title == "" {
			return c.paint(p.BorderTone, left) + border(inner+2*vw-frame) + c.paint(p.BorderTone, right)
		}

		return c.paint(p.BorderTone, left) + p.titled(c, title, inner+2*vw-frame, hw, border) + c.paint(p.BorderTone, right)
	}

	vertical := c.paint(p.BorderTone, g.Vertical)
	blank := vertical + strings.Repeat(" ", inner) + vertical
	contentW := max(inner-pad, 0)

	out := make([]string, 0, len(lines)+p.TopPadding+p.BottomPadding+2)
	out = append(out, edge(g.TopLeft, g.TopRight, p.TitlePosition <= BoxTitleTopRight))

	for range p.TopPadding {
		out = append(out, blank)
	}

	for _, line := range lines {
		line = text.PadRight(c.style(p.TextStyle, c.cut(line, contentW)), contentW)
		out = append(out, vertical+strings.Repeat(" ", p.LeftPadding)+line+strings.Repeat(" ", max(inner-p.LeftPadding-contentW, 0))+vertical)
	}

	for range p.BottomPadding {
		out = append(out, blank)
	}

	out = append(out, edge(g.BottomLeft, g.BottomRight, p.TitlePosition >= BoxTitleBottomLeft))

	return strings.Join(out, "\n")
}

// titled draws a border of width columns with the title set into it.
func (p BoxPrinter) titled(c canvas, title string, width, hw int, border func(int) string) string {
	inset := p.TitleInset * hw
	title = c.cut(title, max(width-2-2*inset, 1))
	rest := max(width-text.Width(title)-2, 0)
	label := " " + c.style(p.TitleStyle, title) + " "

	var left int

	switch p.TitlePosition {
	case BoxTitleTopCenter, BoxTitleBottomCenter:
		left = rest / 2 / hw * hw
	case BoxTitleTopRight, BoxTitleBottomRight:
		left = rest - min(inset, rest)
	default:
		left = min(inset, rest)
	}

	return border(left) + label + border(rest-left)
}

// Sprint returns the text in a box.
func (p BoxPrinter) Sprint(a ...any) string { return sprintText(p, fmt.Sprint(a...)) }

// Sprintln is like Sprint and adds a newline.
func (p BoxPrinter) Sprintln(a ...any) string { return sprintText(p, sprintln(a...)) + "\n" }

// Sprintf formats the text and returns it in a box.
func (p BoxPrinter) Sprintf(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...))
}

// Sprintfln is like Sprintf and adds a newline.
func (p BoxPrinter) Sprintfln(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...)) + "\n"
}

// Print prints the text in a box. A box is a whole block, so Print and
// Println print the same.
func (p BoxPrinter) Print(a ...any) { printLine(p, fmt.Sprint(a...)) }

// Println prints the text in a box.
func (p BoxPrinter) Println(a ...any) { printLine(p, sprintln(a...)) }

// Printf formats the text and prints it in a box.
func (p BoxPrinter) Printf(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// Printfln formats the text and prints it in a box.
func (p BoxPrinter) Printfln(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// PrintOnError prints a box for every non-nil error in a.
func (p BoxPrinter) PrintOnError(a ...any) { printOnError(p, a...) }

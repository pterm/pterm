package pterm

import (
	"io"
	"strings"

	"atomicgo.dev/color"

	"github.com/pterm/pterm/internal/text"
)

// PanelItem is one block of a panel layout. Data can hold several lines
// and the output of other printers.
type PanelItem struct {
	// Title is inset into the frame of a boxed item, or printed above the
	// data of an item without a frame.
	Title string
	// Data is the content.
	Data string
}

// Panels is a grid of items: rows of items that sit side by side.
type Panels [][]PanelItem

// PanelPrinter lays out blocks of text side by side, in rows. Boxed items
// get a rounded frame with their title inset:
//
//	╭─ Uptime ─╮  ╭─ Errors ─╮
//	│ 99.99%   │  │ 0.02%    │
//	╰──────────╯  ╰──────────╯
//
// A row that is wider than the terminal continues on the next line.
type PanelPrinter struct {
	// Panels holds the rows of items.
	Panels Panels
	// Padding is the number of columns between two items of a row.
	Padding int
	// BottomPadding is the number of blank lines between two rows.
	BottomPadding int
	// SameColumnWidth makes all items of a column as wide as the widest.
	SameColumnWidth bool
	// Boxed draws a rounded frame around every item.
	Boxed bool
	// BoxPadding is the space between a frame and its content.
	BoxPadding int
	// BorderTone colors the frames.
	BorderTone Tone
	// TitleStyle styles the titles.
	TitleStyle Style
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Panel is the default panel layout.
var Panel = PanelPrinter{
	Padding:    2,
	BoxPadding: 1,
	BorderTone: ToneBorder,
	TitleStyle: TonePrimary.Bold(),
}

// WithPanels returns a copy of p that lays out panels.
func (p PanelPrinter) WithPanels(panels Panels) *PanelPrinter {
	p.Panels = panels
	return &p
}

// WithRow returns a copy of p with one more row of items.
func (p PanelPrinter) WithRow(items ...PanelItem) *PanelPrinter {
	p.Panels = append(append(Panels(nil), p.Panels...), items)
	return &p
}

// WithPadding returns a copy of p with n columns between two items.
func (p PanelPrinter) WithPadding(n int) *PanelPrinter {
	p.Padding = max(n, 0)
	return &p
}

// WithBottomPadding returns a copy of p with n blank lines between rows.
func (p PanelPrinter) WithBottomPadding(n int) *PanelPrinter {
	p.BottomPadding = max(n, 0)
	return &p
}

// WithSameColumnWidth returns a copy of p that makes all items of a column
// as wide as the widest, or not with false.
func (p PanelPrinter) WithSameColumnWidth(b ...bool) *PanelPrinter {
	p.SameColumnWidth = flag(b)
	return &p
}

// WithBoxed returns a copy of p that frames every item, or none with false.
func (p PanelPrinter) WithBoxed(b ...bool) *PanelPrinter {
	p.Boxed = flag(b)
	return &p
}

// WithBoxPadding returns a copy of p with n spaces between a frame and its
// content.
func (p PanelPrinter) WithBoxPadding(n int) *PanelPrinter {
	p.BoxPadding = max(n, 0)
	return &p
}

// WithBorderTone returns a copy of p that draws its frames in tone t.
func (p PanelPrinter) WithBorderTone(t Tone) *PanelPrinter {
	p.BorderTone = t
	return &p
}

// WithTitleStyle returns a copy of p whose titles use s.
func (p PanelPrinter) WithTitleStyle(s Style) *PanelPrinter {
	p.TitleStyle = s
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p PanelPrinter) WithWriter(w io.Writer) *PanelPrinter {
	p.Writer = w
	return &p
}

func (p PanelPrinter) out() io.Writer { return orDefault(p.Writer) }

// Render prints the panels.
func (p PanelPrinter) Render() error {
	s, err := p.Srender()
	return renderTo(p.out(), s, err)
}

// panelBlock is an item split into lines, before framing.
type panelBlock struct {
	title string
	lines []string
	width int // content width
}

// Srender returns the panels. It returns ErrEmpty when there are no items.
func (p PanelPrinter) Srender() (string, error) {
	c := canvasFor(p.out())
	rows := p.layout(c)

	if len(rows) == 0 {
		return "", ErrEmpty
	}

	widths := p.columnWidths(rows)

	var out []string

	for r, row := range rows {
		if r > 0 {
			for range p.BottomPadding {
				out = append(out, "")
			}
		}

		out = append(out, p.renderRow(c, row, widths)...)
	}

	return strings.Join(out, "\n"), nil
}

// layout splits the items into blocks and breaks rows that are wider than
// the canvas.
func (p PanelPrinter) layout(c canvas) [][]panelBlock {
	frame := 0
	if p.Boxed {
		frame = 2 + 2*p.BoxPadding
	}

	var rows [][]panelBlock

	for _, items := range p.Panels {
		var (
			row  []panelBlock
			used int
		)

		for _, item := range items {
			b := p.block(c, item, frame)

			w := b.width + frame
			if len(row) > 0 && c.width > 0 && used+p.Padding+w > c.width {
				rows = append(rows, row)
				row, used = nil, 0
			}

			if len(row) > 0 {
				used += p.Padding
			}

			row = append(row, b)
			used += w
		}

		if len(row) > 0 {
			rows = append(rows, row)
		}
	}

	return rows
}

// block returns item as lines, cut to fit the canvas with its frame.
func (p PanelPrinter) block(c canvas, item PanelItem, frame int) panelBlock {
	data := strings.TrimRight(strings.ReplaceAll(item.Data, "\r\n", "\n"), "\n")
	b := panelBlock{title: item.Title, lines: strings.Split(data, "\n")}

	if !p.Boxed && item.Title != "" {
		b.lines = append([]string{c.style(p.TitleStyle, item.Title)}, b.lines...)
	}

	for i, line := range b.lines {
		if c.width > 0 {
			line = c.cut(line, max(c.width-frame, 1))
		}

		// End every styled line with a reset, so its colors cannot leak
		// into the item to its right.
		if strings.Contains(line, "\x1b") {
			line += color.Reset
		}

		b.lines[i] = line
		b.width = max(b.width, text.Width(line))
	}

	if p.Boxed && item.Title != "" {
		// The title needs "─ " before and " ─" after it inside the frame.
		b.width = max(b.width, text.Width(item.Title)+4-2*p.BoxPadding)
		if c.width > 0 {
			b.width = min(b.width, max(c.width-frame, 1))
		}
	}

	return b
}

// columnWidths returns the content width of every column with
// SameColumnWidth, or nil.
func (p PanelPrinter) columnWidths(rows [][]panelBlock) []int {
	if !p.SameColumnWidth {
		return nil
	}

	var widths []int

	for _, row := range rows {
		for i, b := range row {
			if i >= len(widths) {
				widths = append(widths, 0)
			}

			widths[i] = max(widths[i], b.width)
		}
	}

	return widths
}

// renderRow returns the lines of one row of blocks.
func (p PanelPrinter) renderRow(c canvas, row []panelBlock, widths []int) []string {
	// Frames of one row are as tall as the tallest, so they end on the
	// same line.
	inner := 0
	for _, b := range row {
		inner = max(inner, len(b.lines))
	}

	cells := make([][]string, len(row))
	cellWidths := make([]int, len(row))
	height := 0

	for i, b := range row {
		w := b.width
		if i < len(widths) {
			w = widths[i]
		}

		cells[i] = p.frameBlock(c, b, w, inner)
		cellWidths[i] = text.MaxWidth(strings.Join(cells[i], "\n"))
		height = max(height, len(cells[i]))
	}

	gap := strings.Repeat(" ", p.Padding)
	lines := make([]string, height)

	for l := range height {
		var b strings.Builder

		for i := range row {
			if i > 0 {
				b.WriteString(gap)
			}

			b.WriteString(text.PadRight(at(cells[i], l), cellWidths[i]))
		}

		lines[l] = strings.TrimRight(b.String(), " ")
	}

	return lines
}

// frameBlock returns the lines of b, framed when p is boxed, with content
// width w and, when boxed, height lines of content.
func (p PanelPrinter) frameBlock(c canvas, b panelBlock, w, height int) []string {
	if !p.Boxed {
		return b.lines
	}

	f := dataFrame{c: c, tone: p.BorderTone, inner: w + 2*p.BoxPadding}
	pad := strings.Repeat(" ", p.BoxPadding)

	lines := []string{f.top(b.title, p.TitleStyle)}
	for i := range max(height, len(b.lines)) {
		lines = append(lines, f.line(pad+text.PadRight(at(b.lines, i), w)+pad))
	}

	return append(lines, f.bottom())
}

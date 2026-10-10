package pterm

import (
	"io"
	"math"
	"strings"

	"atomicgo.dev/color"

	"github.com/pterm/pterm/internal/text"
)

// HeatmapData holds the values of a heatmap, one slice per row.
type HeatmapData [][]float64

// HeatmapAxis labels the columns (XAxis) and rows (YAxis) of a heatmap.
type HeatmapAxis struct {
	XAxis []string
	YAxis []string
}

// HeatmapPrinter renders a grid of values as colored cells. Colors run
// along a gradient of theme tones from the smallest to the largest value,
// and a legend under the grid shows the scale.
//
//	Mon  ██ ██ ██ ██
//	Tue  ██ ██ ██ ██
//	     9h 12h 15h 18h
//
//	     2 ████████████████ 31
type HeatmapPrinter struct {
	// Data holds the values, one slice per row.
	Data HeatmapData
	// Axis labels the rows and columns. Leave it empty for no labels.
	Axis HeatmapAxis
	// Gradient lists the tones the colors run through, from the smallest
	// value to the largest.
	Gradient []Tone
	// Min and Max fix the scale. When they are equal, the scale runs from
	// the smallest to the largest value.
	Min, Max float64
	// ShowValues prints the values in their color instead of drawing
	// blocks.
	ShowValues bool
	// Format turns a value into text. Nil prints up to two decimals.
	Format func(v float64) string
	// CellWidth is the minimum width of a cell. Cells grow to fit axis
	// labels and values.
	CellWidth int
	// Grid draws lines between the cells.
	Grid bool
	// Boxed draws a frame around the grid. It implies Grid.
	Boxed bool
	// Legend adds the scale under the grid.
	Legend bool
	// LegendLabel names the legend, such as "commits".
	LegendLabel string
	// LegendWidth is the width of the gradient in the legend.
	LegendWidth int
	// AxisTone colors the axis labels and the legend text.
	AxisTone Tone
	// BorderTone colors the grid and the frame.
	BorderTone Tone
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Heatmap is the default heatmap printer. Its gradient runs from Border
// through Primary to Accent.
var Heatmap = HeatmapPrinter{
	Gradient:    []Tone{ToneBorder, TonePrimary, ToneAccent},
	CellWidth:   2,
	Legend:      true,
	LegendWidth: 20,
	AxisTone:    ToneMuted,
	BorderTone:  ToneBorder,
}

// WithData returns a copy of p that draws data.
func (p HeatmapPrinter) WithData(data HeatmapData) *HeatmapPrinter {
	p.Data = data
	return &p
}

// WithAxisData returns a copy of p with axis labels.
func (p HeatmapPrinter) WithAxisData(axis HeatmapAxis) *HeatmapPrinter {
	p.Axis = axis
	return &p
}

// WithGradient returns a copy of p whose colors run through tones.
func (p HeatmapPrinter) WithGradient(tones ...Tone) *HeatmapPrinter {
	p.Gradient = tones
	return &p
}

// WithRange returns a copy of p with a fixed scale from lo to hi.
func (p HeatmapPrinter) WithRange(lo, hi float64) *HeatmapPrinter {
	p.Min, p.Max = lo, hi
	return &p
}

// WithShowValues returns a copy of p that prints the values instead of
// blocks.
func (p HeatmapPrinter) WithShowValues(b ...bool) *HeatmapPrinter {
	p.ShowValues = flag(b)
	return &p
}

// WithFormat returns a copy of p that turns values into text with f.
func (p HeatmapPrinter) WithFormat(f func(v float64) string) *HeatmapPrinter {
	p.Format = f
	return &p
}

// WithCellWidth returns a copy of p whose cells are at least w columns
// wide.
func (p HeatmapPrinter) WithCellWidth(w int) *HeatmapPrinter {
	p.CellWidth = w
	return &p
}

// WithGrid returns a copy of p with lines between the cells. Turning the
// grid off also removes the frame.
func (p HeatmapPrinter) WithGrid(b ...bool) *HeatmapPrinter {
	p.Grid = flag(b)
	if !p.Grid {
		p.Boxed = false
	}

	return &p
}

// WithBoxed returns a copy of p with a frame around the grid.
func (p HeatmapPrinter) WithBoxed(b ...bool) *HeatmapPrinter {
	p.Boxed = flag(b)
	if p.Boxed {
		p.Grid = true
	}

	return &p
}

// WithLegend returns a copy of p with the scale under the grid.
func (p HeatmapPrinter) WithLegend(b ...bool) *HeatmapPrinter {
	p.Legend = flag(b)
	return &p
}

// WithLegendLabel returns a copy of p whose legend is named label. It
// turns the legend on.
func (p HeatmapPrinter) WithLegendLabel(label string) *HeatmapPrinter {
	p.LegendLabel = label
	p.Legend = true

	return &p
}

// WithLegendWidth returns a copy of p whose legend gradient is w columns
// wide.
func (p HeatmapPrinter) WithLegendWidth(w int) *HeatmapPrinter {
	p.LegendWidth = w
	return &p
}

// WithAxisTone returns a copy of p whose axis labels use tone t.
func (p HeatmapPrinter) WithAxisTone(t Tone) *HeatmapPrinter {
	p.AxisTone = t
	return &p
}

// WithBorderTone returns a copy of p whose grid and frame use tone t.
func (p HeatmapPrinter) WithBorderTone(t Tone) *HeatmapPrinter {
	p.BorderTone = t
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p HeatmapPrinter) WithWriter(w io.Writer) *HeatmapPrinter {
	p.Writer = w
	return &p
}

func (p HeatmapPrinter) out() io.Writer { return orDefault(p.Writer) }

// Srender returns the heatmap. It returns ErrEmpty without data.
func (p HeatmapPrinter) Srender() (string, error) {
	cols := 0
	for _, row := range p.Data {
		cols = max(cols, len(row))
	}

	if cols == 0 {
		return "", ErrEmpty
	}

	return p.render(canvasFor(p.out()), cols), nil
}

// Render prints the heatmap.
func (p HeatmapPrinter) Render() error {
	s, err := p.Srender()
	return renderTo(p.out(), s, err)
}

func (p HeatmapPrinter) format(v float64) string {
	if p.Format != nil {
		return p.Format(v)
	}

	return formatNumber(v)
}

func (p HeatmapPrinter) render(c canvas, cols int) string {
	var values []float64
	for _, row := range p.Data {
		values = append(values, row...)
	}

	sc := fitScale(p.Min, p.Max, values, false)
	grid := p.Grid || p.Boxed

	cellW := max(p.CellWidth, 1)
	for _, label := range p.Axis.XAxis {
		cellW = max(cellW, text.Width(label))
	}

	if p.ShowValues {
		for _, v := range values {
			cellW = max(cellW, text.Width(p.format(v)))
		}
	}

	yW := 0
	for _, label := range p.Axis.YAxis {
		yW = max(yW, text.Width(label))
	}

	// Without a grid, cells are apart by a space. With one, they are
	// apart by a line and get a space of room on both sides.
	sep, edge, pad := " ", "", ""
	if grid {
		sep, pad = c.paint(p.BorderTone, c.g.Box.Vertical), " "
	}

	if p.Boxed {
		edge = sep
	}

	// rule draws a horizontal grid line with the given joints.
	rule := func(left, cross, right string) string {
		parts := make([]string, 0, cols+1)
		if yW > 0 {
			parts = append(parts, strings.Repeat(c.g.Box.Horizontal, yW+2*len(pad)))
		}

		for range cols {
			parts = append(parts, strings.Repeat(c.g.Box.Horizontal, cellW+2*len(pad)))
		}

		if !p.Boxed {
			left, right = "", ""
		}

		return c.paint(p.BorderTone, left+strings.Join(parts, cross)+right)
	}

	row := func(label string, cells []string) string {
		parts := make([]string, 0, len(cells)+1)
		if yW > 0 {
			parts = append(parts, pad+text.PadRight(c.paint(p.AxisTone, label), yW)+pad)
		}

		for _, cell := range cells {
			parts = append(parts, pad+cell+pad)
		}

		return edge + strings.Join(parts, sep) + edge
	}

	b := c.g.Box

	var lines []string

	if p.Boxed {
		lines = append(lines, rule(b.TopLeft, b.TeeTop, b.TopRight))
	}

	for y, data := range p.Data {
		if y > 0 && grid {
			lines = append(lines, rule(b.TeeLeft, b.Cross, b.TeeRight))
		}

		cells := make([]string, cols)
		for x := range cols {
			cells[x] = strings.Repeat(" ", cellW)
			if x < len(data) && !math.IsNaN(data[x]) {
				cells[x] = p.cell(c, sc, data[x], cellW)
			}
		}

		label := ""
		if y < len(p.Axis.YAxis) {
			label = p.Axis.YAxis[y]
		}

		lines = append(lines, row(label, cells))
	}

	if len(p.Axis.XAxis) > 0 {
		if grid {
			lines = append(lines, rule(b.TeeLeft, b.Cross, b.TeeRight))
		}

		cells := make([]string, cols)
		for x := range cols {
			label := ""
			if x < len(p.Axis.XAxis) {
				label = p.Axis.XAxis[x]
			}

			// Labels line up with what is above them: centered under
			// blocks, right-aligned under numbers.
			if p.ShowValues {
				cells[x] = text.PadLeft(c.paint(p.AxisTone, label), cellW)
			} else {
				cells[x] = text.Center(c.paint(p.AxisTone, label), cellW)
			}
		}

		lines = append(lines, row("", cells))
	}

	if p.Boxed {
		lines = append(lines, rule(b.BottomLeft, b.TeeBottom, b.BottomRight))
	}

	if p.Legend && StylingEnabled() {
		indent := 0
		if yW > 0 {
			indent = yW + 2*len(pad) + 1
		}

		if p.Boxed {
			indent += 1 + len(pad)
		} else {
			indent += len(pad)
		}

		lines = append(lines, "", strings.Repeat(" ", indent)+p.legend(c, sc))
	}

	for i, line := range lines {
		lines[i] = c.cut(strings.TrimRight(line, " "), c.width)
	}

	return strings.Join(lines, "\n")
}

// cell draws value v: a block in its color, or the value itself.
func (p HeatmapPrinter) cell(c canvas, sc scale, v float64, width int) string {
	r := sc.ratio(v)
	if p.ShowValues {
		return text.PadLeft(c.paint(Fixed(gradientAt(c.theme, p.Gradient, r)), p.format(v)), width)
	}

	return p.block(c, r, width)
}

// block draws n columns of blocks at ratio r of the scale. Without
// colors, the height of the blocks shows the value instead.
func (p HeatmapPrinter) block(c canvas, r float64, n int) string {
	if c.level == color.LevelNone && len(c.g.Columns) > 0 {
		return strings.Repeat(c.g.Columns[int(math.Round(r*float64(len(c.g.Columns)-1)))], n)
	}

	return c.paint(Fixed(gradientAt(c.theme, p.Gradient, r)), strings.Repeat(c.g.MeterFill, n))
}

// legend draws the scale: the smallest value, the gradient and the
// largest value.
func (p HeatmapPrinter) legend(c canvas, sc scale) string {
	width := max(p.LegendWidth, 2)

	var b strings.Builder

	if p.LegendLabel != "" {
		b.WriteString(c.paint(p.AxisTone, p.LegendLabel) + "  ")
	}

	b.WriteString(c.paint(p.AxisTone, p.format(sc.lo)) + " ")

	for i := range width {
		b.WriteString(p.block(c, float64(i)/float64(width-1), 1))
	}

	b.WriteString(" " + c.paint(p.AxisTone, p.format(sc.hi)))

	return b.String()
}

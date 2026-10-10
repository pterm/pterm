package pterm

import (
	"io"
	"math"
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// Bar is one bar of a simple bar chart.
type Bar struct {
	// Label names the bar.
	Label string
	// Value is the length of the bar. Negative values grow the other way.
	Value float64
	// Tone colors the bar. The zero value takes the chart's tone, or the
	// next tone of the series cycle.
	Tone Tone
}

// Bars is a list of bars.
type Bars []Bar

// BarGroup is one label of a grouped or stacked chart, with one value per
// series.
type BarGroup struct {
	// Label names the group.
	Label string
	// Values holds one value per series, in the order of the chart's
	// series.
	Values []float64
}

// BarGroups is a list of bar groups.
type BarGroups []BarGroup

// BarChartPrinter renders bar charts. The default chart is horizontal:
// one bar per line, exact to an eighth of a column, on a dotted track,
// with the value after it.
//
//	npm    ████████████████▍░░░░░░░░░░░░   12
//	pnpm   ██████████████████████████▊░░   19
//	yarn   █████████▋░░░░░░░░░░░░░░░░░░░    7
//
// Groups with several series render grouped (one bar per series) or
// stacked (one bar with a segment per series), with a legend. Vertical
// charts draw columns over a baseline instead. Negative values grow from
// a zero axis in the other direction.
type BarChartPrinter struct {
	// Bars are the bars of a simple chart.
	Bars Bars
	// Groups are the bars of a grouped or stacked chart. When set, Bars
	// are ignored.
	Groups BarGroups
	// Series names the values of every group, for the legend.
	Series []string
	// Stacked draws groups as one bar with a segment per series, instead
	// of one bar per series. Negative values count as zero when stacked.
	Stacked bool
	// Vertical draws columns over a baseline instead of horizontal bars.
	Vertical bool
	// Width is the width of a horizontal chart, labels and values
	// included. Zero draws bars of up to 48 columns that shrink to fit
	// the terminal, or 40 columns when its width is unknown.
	Width int
	// Height is the number of lines the columns of a vertical chart take.
	Height int
	// BarWidth is the width of one column of a vertical chart.
	BarWidth int
	// Min and Max fix the scale. When they are equal, the scale runs from
	// zero to the largest value, or to the smallest negative value.
	Min, Max float64
	// Tone colors every bar of a simple chart. The zero value colors each
	// bar with the next tone of Tones.
	Tone Tone
	// Tones is the cycle bars and series take their tones from. Nil uses
	// SeriesTones.
	Tones []Tone
	// ShowValues prints the value next to every bar.
	ShowValues bool
	// ShowTrack draws the dotted track behind horizontal bars.
	ShowTrack bool
	// ShowLegend adds a legend to grouped and stacked charts that have
	// series names.
	ShowLegend bool
	// Unit follows every value, such as "%" or " MB".
	Unit string
	// Format turns a value into text. Nil prints up to two decimals.
	Format func(v float64) string
	// LabelTone colors the labels.
	LabelTone Tone
	// ValueTone colors the values.
	ValueTone Tone
	// TrackTone colors the track.
	TrackTone Tone
	// AxisTone colors the baseline of vertical charts and the zero axis.
	AxisTone Tone
	// LegendTone colors the series names in the legend.
	LegendTone Tone
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// The default lengths of horizontal tracks: at most maxTrack columns on a
// terminal, unknownTrack when the width is unknown.
const (
	maxTrack     = 48
	unknownTrack = 40
)

// BarChart is the default bar chart printer.
var BarChart = BarChartPrinter{
	Height:     8,
	BarWidth:   2,
	ShowValues: true,
	ShowTrack:  true,
	ShowLegend: true,
	LabelTone:  ToneText,
	ValueTone:  ToneMuted,
	TrackTone:  ToneBorder,
	AxisTone:   ToneBorder,
	LegendTone: ToneText,
}

// WithBars returns a copy of p that draws bars.
func (p BarChartPrinter) WithBars(bars Bars) *BarChartPrinter {
	p.Bars = bars
	return &p
}

// WithGroups returns a copy of p that draws groups, one value per series.
func (p BarChartPrinter) WithGroups(groups BarGroups) *BarChartPrinter {
	p.Groups = groups
	return &p
}

// WithSeries returns a copy of p whose series have the given names. They
// make up the legend of grouped and stacked charts.
func (p BarChartPrinter) WithSeries(names ...string) *BarChartPrinter {
	p.Series = names
	return &p
}

// WithStacked returns a copy of p that stacks the values of every group
// into one bar.
func (p BarChartPrinter) WithStacked(b ...bool) *BarChartPrinter {
	p.Stacked = flag(b)
	return &p
}

// WithVertical returns a copy of p that draws columns instead of bars.
func (p BarChartPrinter) WithVertical(b ...bool) *BarChartPrinter {
	p.Vertical = flag(b)
	return &p
}

// WithWidth returns a copy of p that is w columns wide.
func (p BarChartPrinter) WithWidth(w int) *BarChartPrinter {
	p.Width = w
	return &p
}

// WithHeight returns a copy of p whose columns are h lines tall.
func (p BarChartPrinter) WithHeight(h int) *BarChartPrinter {
	p.Height = h
	return &p
}

// WithBarWidth returns a copy of p whose vertical columns are w columns
// wide.
func (p BarChartPrinter) WithBarWidth(w int) *BarChartPrinter {
	p.BarWidth = w
	return &p
}

// WithRange returns a copy of p with a fixed scale from lo to hi. Bars
// that grow over time, or several charts that should compare, need one.
func (p BarChartPrinter) WithRange(lo, hi float64) *BarChartPrinter {
	p.Min, p.Max = lo, hi
	return &p
}

// WithTone returns a copy of p that colors every bar in tone t.
func (p BarChartPrinter) WithTone(t Tone) *BarChartPrinter {
	p.Tone = t
	return &p
}

// WithTones returns a copy of p that takes the tones of bars and series
// from tones, in turn.
func (p BarChartPrinter) WithTones(tones ...Tone) *BarChartPrinter {
	p.Tones = tones
	return &p
}

// WithShowValues returns a copy of p that prints values next to the bars.
func (p BarChartPrinter) WithShowValues(b ...bool) *BarChartPrinter {
	p.ShowValues = flag(b)
	return &p
}

// WithShowTrack returns a copy of p that draws the track behind bars.
func (p BarChartPrinter) WithShowTrack(b ...bool) *BarChartPrinter {
	p.ShowTrack = flag(b)
	return &p
}

// WithShowLegend returns a copy of p that adds a legend to grouped and
// stacked charts.
func (p BarChartPrinter) WithShowLegend(b ...bool) *BarChartPrinter {
	p.ShowLegend = flag(b)
	return &p
}

// WithUnit returns a copy of p that prints unit after every value.
func (p BarChartPrinter) WithUnit(unit string) *BarChartPrinter {
	p.Unit = unit
	return &p
}

// WithFormat returns a copy of p that turns values into text with f.
func (p BarChartPrinter) WithFormat(f func(v float64) string) *BarChartPrinter {
	p.Format = f
	return &p
}

// WithLabelTone returns a copy of p whose labels use tone t.
func (p BarChartPrinter) WithLabelTone(t Tone) *BarChartPrinter {
	p.LabelTone = t
	return &p
}

// WithValueTone returns a copy of p whose values use tone t.
func (p BarChartPrinter) WithValueTone(t Tone) *BarChartPrinter {
	p.ValueTone = t
	return &p
}

// WithTrackTone returns a copy of p whose track uses tone t.
func (p BarChartPrinter) WithTrackTone(t Tone) *BarChartPrinter {
	p.TrackTone = t
	return &p
}

// WithAxisTone returns a copy of p whose baseline and zero axis use tone t.
func (p BarChartPrinter) WithAxisTone(t Tone) *BarChartPrinter {
	p.AxisTone = t
	return &p
}

// WithLegendTone returns a copy of p whose legend uses tone t.
func (p BarChartPrinter) WithLegendTone(t Tone) *BarChartPrinter {
	p.LegendTone = t
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p BarChartPrinter) WithWriter(w io.Writer) *BarChartPrinter {
	p.Writer = w
	return &p
}

func (p BarChartPrinter) out() io.Writer { return orDefault(p.Writer) }

// Srender returns the chart. It returns ErrEmpty without bars.
func (p BarChartPrinter) Srender() (string, error) {
	groups := p.groups()
	if len(groups) == 0 {
		return "", ErrEmpty
	}

	c := canvasFor(p.out())

	if !StylingEnabled() {
		return p.renderRaw(groups), nil
	}

	var lines []string
	if p.Vertical {
		lines = p.renderVertical(c, groups)
	} else {
		lines = p.renderHorizontal(c, groups)
	}

	if legend := p.legend(c); legend != "" {
		indent := 0
		if !p.Vertical {
			indent = p.labelWidth(c, groups) + 2
		}

		lines = append(lines, "", strings.Repeat(" ", indent)+legend)
	}

	for i, line := range lines {
		lines[i] = c.cut(strings.TrimRight(line, " "), c.width)
	}

	return strings.Join(lines, "\n"), nil
}

// Render prints the chart.
func (p BarChartPrinter) Render() error {
	s, err := p.Srender()
	return renderTo(p.out(), s, err)
}

// chartGroup is a label with its entries, the shape every mode renders.
type chartGroup struct {
	label   string
	entries []segment
}

// total is the length of a stacked bar: the sum of its positive entries.
func (g chartGroup) total() float64 {
	sum := 0.0
	for _, e := range g.entries {
		sum += max(e.length, 0)
	}

	return sum
}

// groups turns bars or groups into chart groups with their tones.
func (p BarChartPrinter) groups() []chartGroup {
	if len(p.Groups) > 0 {
		groups := make([]chartGroup, 0, len(p.Groups))

		for _, g := range p.Groups {
			cg := chartGroup{label: flat(g.Label)}

			for i, v := range g.Values {
				if p.Stacked {
					v = max(v, 0)
				}

				cg.entries = append(cg.entries, segment{v, cycleTone(p.Tones, i)})
			}

			groups = append(groups, cg)
		}

		return groups
	}

	groups := make([]chartGroup, 0, len(p.Bars))

	for i, b := range p.Bars {
		tone := b.Tone
		if tone == (Tone{}) {
			tone = p.Tone
		}

		if tone == (Tone{}) {
			tone = cycleTone(p.Tones, i)
		}

		groups = append(groups, chartGroup{label: flat(b.Label), entries: []segment{{b.Value, tone}}})
	}

	return groups
}

// stacked reports whether groups render as one bar each.
func (p BarChartPrinter) stacked() bool {
	return p.Stacked && len(p.Groups) > 0
}

// rows returns the bars to draw for g: one per entry, or one in total
// when stacked.
func (p BarChartPrinter) rows(g chartGroup) [][]segment {
	if p.stacked() {
		return [][]segment{g.entries}
	}

	rows := make([][]segment, len(g.entries))
	for i, e := range g.entries {
		rows[i] = []segment{e}
	}

	return rows
}

// value is the text of a bar: its value, or the total of a stack.
func (p BarChartPrinter) value(segs []segment) string {
	v := 0.0
	for _, s := range segs {
		v += s.length
	}

	if p.Format != nil {
		return p.Format(v) + p.Unit
	}

	return formatNumber(v) + p.Unit
}

// scale fits the chart's scale to what it draws.
func (p BarChartPrinter) scale(groups []chartGroup) scale {
	var values []float64

	for _, g := range groups {
		if p.stacked() {
			values = append(values, g.total())
			continue
		}

		for _, e := range g.entries {
			values = append(values, e.length)
		}
	}

	return fitScale(p.Min, p.Max, values, true)
}

// labelWidth is the width of the label column of a horizontal chart. On
// a known width it takes at most a third of it.
func (p BarChartPrinter) labelWidth(c canvas, groups []chartGroup) int {
	w := 0
	for _, g := range groups {
		w = max(w, text.Width(g.label))
	}

	if c.width > 0 {
		w = min(w, max(c.width/3, 4))
	}

	return w
}

func (p BarChartPrinter) renderHorizontal(c canvas, groups []chartGroup) []string {
	labelW := p.labelWidth(c, groups)

	valueW := 0

	if p.ShowValues {
		for _, g := range groups {
			for _, row := range p.rows(g) {
				valueW = max(valueW, text.Width(p.value(row)))
			}
		}
	}

	fixed := labelW + 2
	if valueW > 0 {
		fixed += valueW + 2
	}

	width := p.Width - fixed

	switch {
	case p.Width <= 0 && c.width > 0:
		width = min(c.width-fixed, maxTrack)
	case p.Width <= 0:
		width = unknownTrack
	case c.width > 0:
		width = min(width, c.width-fixed)
	}

	width = max(width, 4)

	sc := p.scale(groups)
	axis := sc.lo < 0

	cells := width
	if axis {
		cells--
	}

	perUnit := 0.0
	if sc.hi > sc.lo {
		perUnit = float64(cells) / (sc.hi - sc.lo)
	}

	neg := int(math.Round(-sc.lo * perUnit))
	pos := cells - neg

	track := ""
	if p.ShowTrack {
		track = c.g.MeterTrack
	}

	lines := make([]string, 0, len(groups))
	spaced := len(p.Groups) > 0 && !p.stacked() && len(groups) > 1

	for gi, g := range groups {
		if spaced && gi > 0 {
			lines = append(lines, "")
		}

		for ri, row := range p.rows(g) {
			label := ""
			if ri == 0 {
				label = c.cut(g.label, labelW)
			}

			line := text.PadRight(c.paint(p.LabelTone, label), labelW) + "  "

			v := 0.0
			for _, s := range row {
				v += s.length
			}

			scaled := make([]segment, len(row))
			for i, s := range row {
				scaled[i] = segment{math.Abs(s.length) * perUnit, s.tone}
			}

			if v < 0 {
				line += drawStackLeft(c, scaled, neg, track, p.TrackTone)
			} else {
				line += c.track(track, p.TrackTone, neg)
			}

			if axis {
				line += c.paint(p.AxisTone, c.g.Box.Vertical)
			}

			if v >= 0 {
				line += drawStack(c, scaled, pos, track, p.TrackTone)
			} else {
				line += c.track(track, p.TrackTone, pos)
			}

			if valueW > 0 {
				line += "  " + text.PadLeft(c.paint(p.ValueTone, p.value(row)), valueW)
			}

			lines = append(lines, line)
		}
	}

	return lines
}

// cell is one cell of a vertical column: a glyph in a tone.
type cell struct {
	glyph string
	tone  Tone
}

// column is a vertical bar: its cells from the baseline outwards, and
// whether it grows up or down.
type column struct {
	cells []cell
	up    bool
}

// slot is the space of one group in a vertical chart: its columns, label
// and value, centered in a common width.
type slot struct {
	width int
	label string
	value string
	cols  []column
}

func (p BarChartPrinter) renderVertical(c canvas, groups []chartGroup) []string {
	height := max(p.Height, 1)
	barW := max(p.BarWidth, 1)
	sc := p.scale(groups)

	perUnit := 0.0
	if sc.hi > sc.lo {
		perUnit = float64(height) / (sc.hi - sc.lo)
	}

	neg := int(math.Round(-sc.lo * perUnit))
	pos := height - neg

	slots := make([]slot, len(groups))

	for gi, g := range groups {
		rows := p.rows(g)
		s := slot{label: g.label}

		for _, row := range rows {
			v := 0.0
			scaled := make([]segment, len(row))

			for i, seg := range row {
				v += seg.length
				scaled[i] = segment{math.Abs(seg.length) * perUnit, seg.tone}
			}

			if v >= 0 {
				s.cols = append(s.cols, column{columnUp(c, scaled, pos), true})
			} else {
				s.cols = append(s.cols, column{columnDown(c, scaled, neg), false})
			}
		}

		// A value names one column. Grouped columns are too narrow for
		// theirs, so the bars speak for themselves there.
		if p.ShowValues && len(rows) == 1 {
			s.value = p.value(rows[0])
		}

		s.width = max(len(rows)*(barW+1)-1, text.Width(s.label), text.Width(s.value))
		slots[gi] = s
	}

	// Values sit right above a column, or right below a negative one, so
	// each side gets a spare line for them.
	above, below := pos, neg

	for _, s := range slots {
		if s.value == "" {
			continue
		}

		if s.cols[0].up {
			above = pos + 1
		} else {
			below = neg + 1
		}
	}

	const gap = "   "

	line := func(part func(s slot) string) string {
		parts := make([]string, len(slots))
		for i, s := range slots {
			parts[i] = part(s)
		}

		return strings.Join(parts, gap)
	}

	lines := make([]string, 0, above+below+2)

	for y := above - 1; y >= 0; y-- {
		lines = append(lines, line(func(s slot) string { return p.slotLine(c, s, barW, y, true) }))
	}

	total := len(gap) * (len(slots) - 1)
	for _, s := range slots {
		total += s.width
	}

	lines = append(lines, c.paint(p.AxisTone, strings.Repeat(c.g.Rule, total)))

	for y := range below {
		lines = append(lines, line(func(s slot) string { return p.slotLine(c, s, barW, y, false) }))
	}

	lines = append(lines, line(func(s slot) string {
		return text.Center(c.paint(p.LabelTone, s.label), s.width)
	}))

	return lines
}

// slotLine draws line y of a slot, counted from the baseline up (or down
// when up is false): the cells of its columns, or its value right past
// the end of its column.
func (p BarChartPrinter) slotLine(c canvas, s slot, barW, y int, up bool) string {
	if s.value != "" && len(s.cols) == 1 && s.cols[0].up == up && y == len(s.cols[0].cells) {
		return text.Center(c.paint(p.ValueTone, s.value), s.width)
	}

	var b strings.Builder

	pad := s.width - (len(s.cols)*(barW+1) - 1)
	b.WriteString(strings.Repeat(" ", pad/2))

	for i, col := range s.cols {
		if i > 0 {
			b.WriteString(" ")
		}

		if col.up == up && y < len(col.cells) {
			b.WriteString(c.paint(col.cells[y].tone, strings.Repeat(col.cells[y].glyph, barW)))
		} else {
			b.WriteString(strings.Repeat(" ", barW))
		}
	}

	b.WriteString(strings.Repeat(" ", pad-pad/2))

	return b.String()
}

// columnUp builds a column that grows up from the baseline, at most
// limit lines tall. Segment boundaries fall on whole lines; the top is
// exact to an eighth of a line.
func columnUp(c canvas, segs []segment, limit int) []cell {
	levels := len(c.g.Columns)
	if levels == 0 || limit <= 0 {
		return nil
	}

	bounds, end := segmentBounds(segs, limit)

	steps := int(math.Round(end * float64(levels)))
	full, part := steps/levels, steps%levels

	cells := make([]cell, 0, full+1)
	for y := range full {
		cells = append(cells, cell{c.g.Columns[levels-1], segs[segmentAt(bounds, y)].tone})
	}

	if part > 0 && full < limit {
		cells = append(cells, cell{c.g.Columns[part-1], segs[len(segs)-1].tone})
	}

	return cells
}

// columnDown builds a column that grows down from the baseline. Without
// eighth blocks for this direction, its end is exact to half a line.
func columnDown(c canvas, segs []segment, limit int) []cell {
	if limit <= 0 {
		return nil
	}

	bounds, end := segmentBounds(segs, limit)

	halves := int(math.Round(end * 2))
	full, half := halves/2, halves%2 == 1

	if half && c.g.UpperHalf == "" {
		full, half = min(full+1, limit), false
	}

	cells := make([]cell, 0, full+1)
	for y := range full {
		cells = append(cells, cell{c.g.MeterFill, segs[segmentAt(bounds, y)].tone})
	}

	if half && full < limit {
		cells = append(cells, cell{c.g.UpperHalf, segs[len(segs)-1].tone})
	}

	return cells
}

// segmentBounds returns where each segment ends, rounded to whole cells,
// and where the last one ends exactly, all capped at limit.
func segmentBounds(segs []segment, limit int) ([]int, float64) {
	bounds := make([]int, len(segs))
	end := 0.0

	for i, s := range segs {
		end += max(s.length, 0)
		bounds[i] = min(int(math.Round(end)), limit)
	}

	return bounds, min(end, float64(limit))
}

// segmentAt returns the segment that covers cell y. Cells past the last
// rounded bound belong to the last segment.
func segmentAt(bounds []int, y int) int {
	for i, b := range bounds {
		if y < b {
			return i
		}
	}

	return len(bounds) - 1
}

// legend returns the series names with their tones, for grouped and
// stacked charts.
func (p BarChartPrinter) legend(c canvas) string {
	if !p.ShowLegend || len(p.Groups) == 0 || len(p.Series) == 0 {
		return ""
	}

	parts := make([]string, len(p.Series))
	for i, name := range p.Series {
		parts[i] = c.paint(cycleTone(p.Tones, i), c.g.Swatch) + " " + c.paint(p.LegendTone, name)
	}

	return strings.Join(parts, "   ")
}

// renderRaw prints the chart as plain "label: value" lines, for output
// without styling.
func (p BarChartPrinter) renderRaw(groups []chartGroup) string {
	lines := make([]string, len(groups))

	for i, g := range groups {
		values := make([]string, len(g.entries))
		for j, e := range g.entries {
			values[j] = p.value([]segment{e})
			if j < len(p.Series) {
				values[j] = p.Series[j] + "=" + values[j]
			}
		}

		lines[i] = g.label + ": " + strings.Join(values, ", ")
	}

	return strings.Join(lines, "\n")
}

// flat puts a label on one line.
func flat(s string) string {
	return strings.ReplaceAll(s, "\n", " ")
}

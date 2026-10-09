package pterm

import (
	"io"
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm/internal"
)

// BarChartPrinter is used to print bar charts.
type BarChartPrinter struct {
	Writer     io.Writer
	Bars       Bars
	Horizontal bool
	ShowValue  bool
	// Height sets the maximum height of a vertical bar chart.
	// The default is calculated to fit into the terminal.
	// Ignored if Horizontal is set to true.
	Height int
	// Width sets the maximum width of a horizontal bar chart.
	// The default is calculated to fit into the terminal.
	// Ignored if Horizontal is set to false.
	Width                  int
	VerticalBarCharacter   string
	HorizontalBarCharacter string
}

var (
	// DefaultBarChart is the default BarChartPrinter.
	DefaultBarChart = BarChartPrinter{
		Horizontal:             false,
		VerticalBarCharacter:   "██",
		HorizontalBarCharacter: "█",
		// keep in sync with RecalculateTerminalSize()
		Height: GetTerminalHeight() * 2 / 3,
		Width:  GetTerminalWidth() * 2 / 3,
	}
)

// WithBars returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithBars(bars Bars) *BarChartPrinter {
	p.Bars = bars
	return &p
}

// WithVerticalBarCharacter returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithVerticalBarCharacter(char string) *BarChartPrinter {
	p.VerticalBarCharacter = char
	return &p
}

// WithHorizontalBarCharacter returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithHorizontalBarCharacter(char string) *BarChartPrinter {
	p.HorizontalBarCharacter = char
	return &p
}

// WithHorizontal returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithHorizontal(b ...bool) *BarChartPrinter {
	b2 := internal.WithBoolean(b)
	p.Horizontal = b2

	return &p
}

// WithHeight returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithHeight(value int) *BarChartPrinter {
	p.Height = value
	return &p
}

// WithWidth returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithWidth(value int) *BarChartPrinter {
	p.Width = value
	return &p
}

// WithShowValue returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithShowValue(b ...bool) *BarChartPrinter {
	p.ShowValue = internal.WithBoolean(b)
	return &p
}

// WithWriter sets the custom Writer.
func (p BarChartPrinter) WithWriter(writer io.Writer) *BarChartPrinter {
	p.Writer = writer
	return &p
}

func (p BarChartPrinter) getRawOutput() string {
	var ret strings.Builder

	for _, bar := range p.Bars {
		ret.WriteString(Sprintfln("%s: %d", bar.Label, bar.Value))
	}

	return ret.String()
}

// barRenderParams carries the geometry of a single bar while it is rendered.
type barRenderParams struct {
	// repeatCount is how many bar characters the bar's value maps to; it is
	// negative for negative values.
	repeatCount int
	bar         Bar
	// positive/negative chart part sizes: how many cells the chart reserves
	// for values above respectively below zero.
	positiveChartPartHeight int
	negativeChartPartHeight int
	positiveChartPartWidth  int
	negativeChartPartWidth  int
	// indent horizontally centers a vertical bar under its label.
	indent    string
	showValue bool
	// moveUp/moveRight shift a positive bar past the chart part reserved for
	// negative values, so both parts share one axis.
	moveUp    bool
	moveRight bool
}

// renderPositiveVerticalBar renders one bar of a vertical chart growing up
// from the zero line, one line per height cell, topmost line first.
func (p BarChartPrinter) renderPositiveVerticalBar(out *strings.Builder, rp barRenderParams) {
	if rp.showValue {
		out.WriteString(Sprint(rp.indent + strconv.Itoa(rp.bar.Value) + rp.indent + "\n"))
	}

	for i := rp.positiveChartPartHeight; i > 0; i-- {
		if i > rp.repeatCount {
			out.WriteString(rp.indent + "  " + rp.indent + " \n")
		} else {
			out.WriteString(rp.indent + rp.bar.Style.Sprint(p.VerticalBarCharacter) + rp.indent + " \n")
		}
	}

	// Used when we draw a chart with both POSITIVE and NEGATIVE values.
	// In that case we separately draw the top and bottom half of the chart,
	// and the positive part has to be MOVED UP past the chart's bottom part
	// by adding empty pillars of the bottom part's height.
	if rp.moveUp {
		for i := 0; i <= rp.negativeChartPartHeight; i++ {
			out.WriteString(rp.indent + "  " + rp.indent + " \n")
		}
	}
}

// renderNegativeVerticalBar renders one bar of a vertical chart growing down
// from the zero line.
func (p BarChartPrinter) renderNegativeVerticalBar(out *strings.Builder, rp barRenderParams) {
	for i := 0; i > -rp.negativeChartPartHeight; i-- {
		if i > rp.repeatCount {
			out.WriteString(rp.indent + rp.bar.Style.Sprint(p.VerticalBarCharacter) + rp.indent + " \n")
		} else {
			out.WriteString(rp.indent + "  " + rp.indent + " \n")
		}
	}

	if rp.showValue {
		out.WriteString(Sprint(rp.indent + strconv.Itoa(rp.bar.Value) + rp.indent + "\n"))
	}
}

// renderPositiveHorizontalBar renders one bar of a horizontal chart growing
// right from the zero line.
func (p BarChartPrinter) renderPositiveHorizontalBar(out *strings.Builder, rp barRenderParams) {
	if rp.moveRight {
		out.WriteString(strings.Repeat(" ", rp.negativeChartPartWidth))
	}

	for i := 0; i < rp.positiveChartPartWidth; i++ {
		if i < rp.repeatCount {
			out.WriteString(rp.bar.Style.Sprint(p.HorizontalBarCharacter))
		} else {
			out.WriteString(" ")
		}
	}

	if rp.showValue {
		// For positive horizontal bars we add one more space before adding the
		// value, so they align with negative values, which carry a "-" sign.
		out.WriteString(" ")
		out.WriteString(" " + strconv.Itoa(rp.bar.Value))
	}
}

// renderNegativeHorizontalBar renders one bar of a horizontal chart growing
// left from the zero line.
func (p BarChartPrinter) renderNegativeHorizontalBar(out *strings.Builder, rp barRenderParams) {
	for i := -rp.negativeChartPartWidth; i < 0; i++ {
		if i < rp.repeatCount {
			out.WriteString(" ")
		} else {
			out.WriteString(rp.bar.Style.Sprint(p.HorizontalBarCharacter))
		}
	}

	// To print the values well-aligned when the chart has both a positive and
	// a negative part, insert an indent as wide as the positive chart part.
	if rp.positiveChartPartWidth > 0 {
		out.WriteString(strings.Repeat(" ", rp.positiveChartPartWidth))
	}

	if rp.showValue {
		// A zero value is rendered by this negative renderer without a "-"
		// sign, so pad it with the space the sign would occupy to keep the
		// values column aligned.
		if rp.repeatCount == 0 {
			out.WriteString(" ")
		}

		out.WriteString(" " + strconv.Itoa(rp.bar.Value))
	}
}

// Srender renders the BarChart as a string.
func (p BarChartPrinter) Srender() (string, error) {
	if rawOutput() {
		return p.getRawOutput(), nil
	}

	p.Bars = p.themedBars()

	minBarValue, maxBarValue, maxLabelHeight := p.barStats()
	maxAbsBarValue := max(intAbs(minBarValue), intAbs(maxBarValue))

	if p.Horizontal {
		return p.srenderHorizontal(minBarValue, maxBarValue, maxAbsBarValue)
	}

	return p.srenderVertical(minBarValue, maxBarValue, maxAbsBarValue, maxLabelHeight), nil
}

// themedBars returns a copy of the configured bars with nil styles replaced
// by the theme defaults and labels pre-rendered with their label style.
// Rendering must not mutate the caller's Bars slice.
func (p BarChartPrinter) themedBars() Bars {
	bars := make(Bars, len(p.Bars))
	copy(bars, p.Bars)

	for i, bar := range bars {
		if bar.Style == nil {
			bars[i].Style = &ThemeDefault.BarStyle
		}

		if bar.LabelStyle == nil {
			bars[i].LabelStyle = &ThemeDefault.BarLabelStyle
		}

		bars[i].Label = bars[i].LabelStyle.Sprint(bar.Label)
	}

	return bars
}

// barStats returns the smallest and largest bar value (both clamped to
// include zero, the chart's baseline) and the line count of the tallest
// label.
func (p BarChartPrinter) barStats() (minValue, maxValue, maxLabelHeight int) {
	for _, bar := range p.Bars {
		minValue = min(minValue, bar.Value)
		maxValue = max(maxValue, bar.Value)
		maxLabelHeight = max(maxLabelHeight, len(strings.Split(bar.Label, "\n")))
	}

	return minValue, maxValue, maxLabelHeight
}

// srenderHorizontal renders the bars as a two-column panel: labels on the
// left, one bar per line on the right.
func (p BarChartPrinter) srenderHorizontal(minBarValue, maxBarValue, maxAbsBarValue int) (string, error) {
	var labels, bars strings.Builder

	rp := barRenderParams{
		showValue:              p.ShowValue,
		positiveChartPartWidth: p.Width,
		negativeChartPartWidth: p.Width,
	}

	// If the chart consists of a positive and a negative part, both parts
	// share the total width proportionally to the value range.
	if minBarValue < 0 && maxBarValue > 0 {
		rp.positiveChartPartWidth = intAbs(internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Width)/2, float32(p.Width)/2, float32(maxBarValue)))
		rp.negativeChartPartWidth = intAbs(internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Width)/2, float32(p.Width)/2, float32(minBarValue)))
	}

	for _, bar := range p.Bars {
		rp.bar = bar

		labels.WriteString("\n")
		labels.WriteString(bar.Label)
		bars.WriteString("\n")

		switch {
		case minBarValue >= 0:
			// Only positive values: draw only the right part of the chart.
			rp.repeatCount = internal.MapRangeToRange(0, float32(maxAbsBarValue), 0, float32(p.Width), float32(bar.Value))
			rp.moveRight = false

			p.renderPositiveHorizontalBar(&bars, rp)

		case maxBarValue <= 0:
			// Only negative values: draw only the left part of the chart.
			rp.repeatCount = internal.MapRangeToRange(-float32(maxAbsBarValue), 0, -float32(p.Width), 0, float32(bar.Value))
			rp.positiveChartPartWidth = 0

			p.renderNegativeHorizontalBar(&bars, rp)

		default:
			// Both positive and negative values: draw both parts.
			rp.repeatCount = internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Width)/2, float32(p.Width)/2, float32(bar.Value))

			if bar.Value >= 0 {
				rp.moveRight = true

				p.renderPositiveHorizontalBar(&bars, rp)
			} else {
				p.renderNegativeHorizontalBar(&bars, rp)
			}
		}
	}

	panels := Panels{[]Panel{{Data: labels.String()}, {Data: bars.String()}}}

	return DefaultPanel.WithPanels(panels).Srender()
}

// srenderVertical renders each bar as a column of lines, then assembles the
// columns side by side, bottom-aligned, with the labels underneath.
func (p BarChartPrinter) srenderVertical(minBarValue, maxBarValue, maxAbsBarValue, maxLabelHeight int) string {
	rp := barRenderParams{
		showValue:               p.ShowValue,
		positiveChartPartHeight: p.Height,
		negativeChartPartHeight: p.Height,
	}

	// If the chart consists of a positive and a negative part, both parts
	// share the total height proportionally to the value range.
	if minBarValue < 0 && maxBarValue > 0 {
		rp.positiveChartPartHeight = intAbs(internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Height)/2, float32(p.Height)/2, float32(maxBarValue)))
		rp.negativeChartPartHeight = intAbs(internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Height)/2, float32(p.Height)/2, float32(minBarValue)))
	}

	renderedBars := make([]string, len(p.Bars))

	for i, bar := range p.Bars {
		var renderedBar strings.Builder

		rp.bar = bar
		rp.indent = strings.Repeat(" ", internal.GetStringMaxWidth(RemoveColorFromString(bar.Label))/2)

		switch {
		case minBarValue >= 0:
			// Only positive values: draw only the top part of the chart.
			rp.repeatCount = internal.MapRangeToRange(0, float32(maxAbsBarValue), 0, float32(p.Height), float32(bar.Value))
			rp.moveUp = false

			p.renderPositiveVerticalBar(&renderedBar, rp)

		case maxBarValue <= 0:
			// Only negative values: draw only the bottom part of the chart.
			rp.repeatCount = internal.MapRangeToRange(-float32(maxAbsBarValue), 0, -float32(p.Height), 0, float32(bar.Value))

			p.renderNegativeVerticalBar(&renderedBar, rp)

		default:
			// Both positive and negative values: draw both parts.
			rp.repeatCount = internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Height)/2, float32(p.Height)/2, float32(bar.Value))

			if bar.Value >= 0 {
				rp.moveUp = true

				p.renderPositiveVerticalBar(&renderedBar, rp)
			} else {
				p.renderNegativeVerticalBar(&renderedBar, rp)
			}
		}

		labelHeight := len(strings.Split(bar.Label, "\n"))
		renderedBars[i] = renderedBar.String() + bar.Label + strings.Repeat("\n", maxLabelHeight-labelHeight) + " "
	}

	return joinBarColumns(renderedBars)
}

// joinBarColumns places the rendered bar columns side by side: each column is
// bottom-aligned by prepending blank lines up to the tallest column, and each
// line is padded to its column's width.
func joinBarColumns(renderedBars []string) string {
	var maxBarHeight int

	for _, bar := range renderedBars {
		maxBarHeight = max(maxBarHeight, len(strings.Split(bar, "\n")))
	}

	for i, bar := range renderedBars {
		totalBarHeight := len(strings.Split(bar, "\n"))
		if totalBarHeight < maxBarHeight {
			renderedBars[i] = strings.Repeat("\n", maxBarHeight-totalBarHeight) + renderedBars[i]
		}
	}

	barLines := make([][]string, len(renderedBars))
	barWidths := make([]int, len(renderedBars))

	for i, bar := range renderedBars {
		barLines[i] = strings.Split(bar, "\n")
		barWidths[i] = internal.GetStringMaxWidth(RemoveColorFromString(bar))
	}

	var ret strings.Builder

	for line := 0; line <= maxBarHeight; line++ {
		for i := range renderedBars {
			var barLine string
			if line < len(barLines[i]) {
				barLine = barLines[i][line]
			}

			lineWidth := runewidth.StringWidth(RemoveColorFromString(barLine))
			if lineWidth < barWidths[i] {
				barLine += strings.Repeat(" ", barWidths[i]-lineWidth)
			}

			ret.WriteString(barLine)
		}

		ret.WriteByte('\n')
	}

	return ret.String()
}

// intAbs returns the absolute value of an int.
func intAbs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

// Render prints the Template to the terminal.
func (p BarChartPrinter) Render() error {
	s, err := p.Srender()
	if err != nil {
		return err
	}

	Fprintln(p.Writer, s)

	return nil
}
